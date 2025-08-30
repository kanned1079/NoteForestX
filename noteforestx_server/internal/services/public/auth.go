package public

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services"
	"noteforestx_server/internal/services/public/dto"
	"regexp"
)

func (this *PublicService) bindAndValidateLoginDto(ctx *gin.Context) (*dto.UserLoginRequestDto, bool) {
	var dto dto.UserLoginRequestDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return nil, false
	}

	if !this.isValidEmail(dto.Email) || len(dto.Password) < config.ExistingAppConfig.Runtime.MinPasswordLen {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email or password not valid",
		})
		return nil, false
	}

	return &dto, true
}

func (this *PublicService) isValidEmail(email string) bool {
	// 这个正则不是最严格的，但能覆盖绝大多数常见邮箱
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (this *PublicService) UserLogin(ctx *gin.Context) {
	var userLoginRequestDto dto.UserLoginRequestDto

	if err := ctx.ShouldBindJSON(&userLoginRequestDto); err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	if !this.isValidEmail(userLoginRequestDto.Email) || !(len(userLoginRequestDto.Password) >= config.ExistingAppConfig.Runtime.MinPasswordLen) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login data is not valid",
		})
		return
	}

	var existingUser models.User
	if result := dao.ExistingDbDaoInst.DbDao.Model(&models.User{}).Where("email = ?", userLoginRequestDto.Email).First(&existingUser); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "cannot find a user with email: " + userLoginRequestDto.Email,
		})
		return
	} else if result.Error != nil {
		services.SendServerInternalError(ctx, result.Error.Error())
		return
	}
	token, err := this.utils.GenerateAccessToken(existingUser)
	if err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	var userLoginResponse = &dto.UserLoginResponseDto{
		User:  existingUser,
		Token: token,
	}

	ctx.JSON(http.StatusOK, userLoginResponse)
}

func (this *PublicService) UserRegister(ctx *gin.Context) {
	if !config.ExistingAppConfig.Runtime.EnableRegister {
		ctx.JSON(http.StatusLocked, gin.H{
			"message": "register not allowed now",
		})
		return
	}

	var userLoginRequestDto dto.UserLoginRequestDto
	if err := ctx.ShouldBindJSON(&userLoginRequestDto); err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	// 基础校验
	if !this.isValidEmail(userLoginRequestDto.Email) || len(userLoginRequestDto.Password) < config.ExistingAppConfig.Runtime.MinPasswordLen {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "register data is not valid",
		})
		return
	}

	// 检查是否已存在
	var existingUser models.User
	if result := dao.ExistingDbDaoInst.DbDao.Model(&models.User{}).Where("email = ?", userLoginRequestDto.Email).First(&existingUser); result.Error == nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "this user has already existed.",
		})
		return
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 非 ErrRecordNotFound 的错误
		services.SendServerInternalError(ctx, result.Error.Error())
		return
	}

	// 密码加密存储（推荐用 bcrypt）
	hashedPwd, err := this.utils.HashPassword(userLoginRequestDto.Password)
	if err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	// 创建用户
	newUser := models.User{
		Email:    userLoginRequestDto.Email,
		Password: hashedPwd,
	}

	if result := dao.ExistingDbDaoInst.DbDao.Create(&newUser); result.Error != nil {
		services.SendServerInternalError(ctx, result.Error.Error())
		return
	}

	// 生成 token
	token, err := this.utils.GenerateAccessToken(newUser)
	if err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	var userLoginResponse = &dto.UserLoginResponseDto{
		User:  newUser,
		Token: token,
	}

	ctx.JSON(http.StatusOK, userLoginResponse)
}
