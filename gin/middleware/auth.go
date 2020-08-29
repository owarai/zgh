package ginmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/owarai/zgh"
	"github.com/owarai/zgh/gin/api"
	"github.com/owarai/zgh/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiG := api.Gin{C: c}

		token := c.Request.Header.Get("x-auth-token")
		if token == "" {
			zgh.ZLog().Error("method", "zgh.ginmiddleware.auth", "error", "token is null")
			apiG.Response(http.StatusOK, 400000001, nil)
			return
		}

		userId, err := jwt.ParseToken(token)
		if err != nil {
			zgh.ZLog().Error("method", "zgh.ginmiddleware.auth", "error", err.Error())
			apiG.Response(http.StatusOK, 400000001, nil)
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
