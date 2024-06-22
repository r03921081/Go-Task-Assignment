package v1

import (
	"fmt"
	"net/http"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/dto"
	"github.com/andyliao/task-homework/model"
	"github.com/gin-gonic/gin"
)

// CreateUserV1Handler godoc
// @Summary      Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  dto.CreateUserRequest  true  "create user"
// @Success      200  {object}  dto.Response{result=model.User}
// @Failure      400  {object}  dto.Response
// @Router       /users [post]
func CreateUserV1Handler(c *gin.Context) {
	req := dto.CreateUserRequest{}
	if _err := c.ShouldBindJSON(&req); _err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, _err.Error()))
		return
	}

	user, err := CreateUser(c, req.Name, req.Password)
	if err != nil {
		common.Logger.Error(c, fmt.Sprintf("Create user: %s failed: %s", req.Name, err.ErrorMessage()))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
		return
	}

	c.JSON(http.StatusOK, dto.NewResponse(c, model.User{
		Username: user.Username,
	}))
}

// LoginV1Hadnler godoc
// @Summary      Login
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  dto.CreateUserRequest  true  "create user"
// @Success      200  {object}  dto.Response{result=string}
// @Failure      400  {object}  dto.Response
// @Router       /users/login [post]
func LoginV1Hadnler(c *gin.Context) {
	req := dto.CreateUserRequest{}
	if _err := c.ShouldBindJSON(&req); _err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, _err.Error()))
		return
	}

	user, err := ValidateUser(c, req.Name, req.Password)
	if err != nil {
		common.Logger.Error(c, fmt.Sprintf("Validate user: %s failed: %s", req.Name, err.ErrorMessage()))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
		return
	}

	authKey, err := CreateAuthKey(c, user)
	if err != nil {
		common.Logger.Error(c, fmt.Sprintf("Create auth key failed: %s", err.ErrorMessage()))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
		return
	}

	c.JSON(http.StatusOK, dto.NewResponse(c, authKey))
}
