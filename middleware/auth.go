package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muharik19/shop-backend/constant"
	"github.com/muharik19/shop-backend/models"
)

// AuthMiddleware is a sample middleware for authentication and authorization using JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(constant.AUTHORIZATION)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:         http.StatusUnauthorized,
				ResponseCode: constant.FAILED_AUTHORIZED,
				ResponseDesc: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		decodes, err := JwtClaim(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:         http.StatusUnauthorized,
				ResponseCode: constant.FAILED_AUTHORIZED,
				ResponseDesc: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		c.Set(constant.GIN_KEY, decodes)

		c.Next()
	}
}
