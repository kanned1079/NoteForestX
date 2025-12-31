package public

import (
	"errors"
	"net/http"
	"noteforestx_server/internal/config"
	"noteforestx_server/internal/dao"
	"noteforestx_server/internal/models"
	"noteforestx_server/internal/services"
	"noteforestx_server/internal/services/public/dto"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (this *PublicService) UserLoginV1(ctx *gin.Context) {
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

	if !this.isValidEmail(userLoginRequestDto.Email) ||
		len(userLoginRequestDto.Password) < config.ExistingAppConfig.Runtime.MinPasswordLen {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "register data is not valid",
		})
		return
	}

	db := dao.ExistingDbDaoInst.DbDao

	// ===== 使用事务，防止并发问题 =====
	err := db.Transaction(func(tx *gorm.DB) error {
		// 检查是否已存在
		var existingUser models.User
		if err := tx.Where("email = ?", userLoginRequestDto.Email).
			First(&existingUser).Error; err == nil {
			//return services.ErrConflict("this user has already existed")
			return errors.New("this user has already existed")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 判断是否是第一个用户
		var userCount int64
		if err := tx.Model(&models.User{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Count(&userCount).Error; err != nil {
			return err
		}

		role := models.UserRoleUser
		if userCount == 0 {
			role = models.UserRoleAdmin
		}

		// 密码加密
		hashedPwd, err := this.utils.HashPassword(userLoginRequestDto.Password)
		if err != nil {
			return err
		}

		newUser := models.User{
			Email:    userLoginRequestDto.Email,
			Password: hashedPwd,
			Role:     role,
		}

		if err := tx.Create(&newUser).Error; err != nil {
			return err
		}

		// 生成 token
		token, err := this.utils.GenerateAccessToken(newUser)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, &dto.UserLoginResponseDto{
			User:  newUser,
			Token: token,
		})

		return nil
	})

	if err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}
}

func (this *PublicService) UserLogin(ctx *gin.Context) {
	var req dto.UserLoginRequestDto

	// 1️⃣ 解析请求
	if err := ctx.ShouldBindJSON(&req); err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	// 2️⃣ 基础校验
	if !this.isValidEmail(req.Email) || len(req.Password) < config.ExistingAppConfig.Runtime.MinPasswordLen {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login data is not valid",
		})
		return
	}

	db := dao.ExistingDbDaoInst.DbDao
	var user models.User

	// 3️⃣ 尝试查询用户
	err := db.Where("email = ?", req.Email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ===== 用户不存在 → 注册 =====
		err = db.Transaction(func(tx *gorm.DB) error {
			// 判断是否是第一个用户
			var count int64
			if err := tx.Model(&models.User{}).Count(&count).Error; err != nil {
				return err
			}

			role := models.UserRoleUser
			if count == 0 {
				role = models.UserRoleAdmin
			}

			// 密码加密
			hashedPwd, err := this.utils.HashPassword(req.Password)
			if err != nil {
				return err
			}

			user = models.User{
				Id:       uuid.New(),
				Email:    req.Email,
				Password: hashedPwd,
				Role:     role,
			}

			// 创建用户
			return tx.Create(&user).Error
		})
		if err != nil {
			services.SendServerInternalError(ctx, err.Error())
			return
		}

	} else if err != nil {
		// 查询错误
		services.SendServerInternalError(ctx, err.Error())
		return
	} else {
		// ===== 用户存在 → 校验密码 =====
		if !this.utils.CheckPasswordHash(req.Password, user.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid email or password",
			})
			return
		}
	}

	// 4️⃣ 生成 token（脱离事务）
	token, err := this.utils.GenerateAccessToken(user)
	if err != nil {
		services.SendServerInternalError(ctx, err.Error())
		return
	}

	// 5️⃣ 返回结果
	ctx.JSON(http.StatusOK, &dto.UserLoginResponseDto{
		User:  user,
		Token: token,
	})
}
