package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func() string {
			return uuid.New().String()
		}),
		requestid.WithCustomHeaderStrKey("Trace-Id"),
	)
}
