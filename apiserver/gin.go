package apiserver

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/andyliao/task-homework/api"
	"github.com/andyliao/task-homework/middleware"
)

func InitGinRouter(ctx context.Context) (*gin.Engine, error) {
	router := gin.Default()

	rootGroup := router.Group("")
	rootGroup.Use(middleware.SetGeneralContextMiddleware())
	api.AddV1HttpEndpoint(rootGroup)

	return router, nil
}
