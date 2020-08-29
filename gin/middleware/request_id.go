package ginmiddleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/owarai/zgh/log"
)

type RequestIDOptions struct {
	AllowSetting bool
}

func RequestID(options RequestIDOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		beginTime := strconv.FormatInt(time.Now().UnixNano(), 10)

		if options.AllowSetting {
			// If Set-Request-Id header is set on request, use that for
			// Request-Id response header. Otherwise, generate a new one.
			requestID = c.Request.Header.Get("Set-Request-Id")
		}

		if requestID == "" {
			s := uuid.New()
			//if err != nil {
			//	zgh.L().Error("message","uuid create  error","error",err.Error())
			//}
			requestID = s.String()
		}

		c.Writer.Header().Set("X-Begin-Time", beginTime)
		c.Writer.Header().Set("X-Request-Id", requestID)
		log.L().Info("Message", "API Request", "header", c.Request.Header)
		c.Next()
	}
}
