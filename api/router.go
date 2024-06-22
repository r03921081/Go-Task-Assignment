package api

import (
	"fmt"

	v1 "github.com/andyliao/task-homework/api/v1"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/middleware"
	"github.com/gin-gonic/gin"
)

func AddV1HttpEndpoint(router *gin.RouterGroup) {
	apiGroup := router.Group("/api")

	v1Group := apiGroup.Group("/v1")
	{
		v1Group.POST("/user", v1.CreateUserV1Handler)
		v1Group.POST("/user/login", v1.LoginV1Hadnler)

		v1Group.Use(middleware.AccessMiddleware())
		v1Group.GET("/tasks", v1.ListTasksV1Handler)
		v1Group.POST("/task", v1.CreateTaskV1Handler)
		v1Group.PUT(fmt.Sprintf("/task/:%s", constant.PathID), v1.PutTaskV1Handler)
		v1Group.DELETE(fmt.Sprintf("/task/:%s", constant.PathID), v1.DeleteTaskV1Handler)
	}
}
