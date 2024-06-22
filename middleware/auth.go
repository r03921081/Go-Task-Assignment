package middleware

import (
	"fmt"
	"net/http"

	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/util"
	"github.com/gin-gonic/gin"

	"github.com/andyliao/task-homework/dto"
)

func SetGeneralContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constant.HeaderTraceID, util.GenerateTraceID())
		c.Next()
	}
}

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAPIKey := c.Request.Header.Get(constant.HeaderAPIKey)
		user, err := IsAuthorized(c, headerAPIKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
			return
		}
		if user != nil {
			c.Set(constant.HeaderUserID, fmt.Sprintf("%d", user.ID))
		}

		c.Next()
	}
}
