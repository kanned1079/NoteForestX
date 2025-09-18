package middlewares

import (
	"github.com/gin-gonic/gin"
)

//func AllowRequestTypeCors() gin.HandlerFunc {
//	return cors.New(cors.Config{
//		AllowOrigins:     []string{"*"},
//		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
//		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
//		AllowCredentials: true, // 注意必须为 false
//		MaxAge:           12 * time.Hour,
//	})
//}

func AllowRequestTypeCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			origin = "http://localhost:3000" // 瀏覽器端前端地址
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//func AllowRequestTypeCors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		origin := c.GetHeader("Origin")
//		c.Header("Access-Control-Allow-Origin", origin) // 动态允许请求来源
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
//		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}
