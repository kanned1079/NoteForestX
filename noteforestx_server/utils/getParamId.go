package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (this *Utils) GetAndParseParamUuid(paramField string, ctx *gin.Context) (uuid.UUID, error) {
	id := ctx.Param(paramField)
	return uuid.Parse(id)
}
