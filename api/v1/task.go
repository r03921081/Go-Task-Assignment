package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/dto"
	"github.com/andyliao/task-homework/model"
	"github.com/gin-gonic/gin"
)

// ListTasksV1Handler godoc
// @Summary      List tasks
// @Tags         task
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response{result=[]model.Task}
// @Router       /tasks [get]
func ListTasksV1Handler(c *gin.Context) {
	tasks := ListTasks(c)

	c.JSON(http.StatusOK, dto.NewResponse(c, tasks))
}

// CreateTaskV1Handler godoc
// @Summary      Create a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param		 name  body      dto.CreateTaskRequest  true  "Create task"
// @Success      200  {object}  dto.Response{result=model.Task}
// @Success      400  {object}  dto.Response
// @Router       /task [post]
func CreateTaskV1Handler(c *gin.Context) {
	req := dto.CreateTaskRequest{}
	if _err := c.ShouldBindJSON(&req); _err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, _err.Error()))
		return
	}

	task, err := CreateTask(c, req.Name)
	if err != nil {
		common.Logger.Error(c, fmt.Sprintf("Create task failed: %s", err.ErrorMessage()))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
		return
	}

	c.JSON(http.StatusCreated, dto.NewResponse(c, task))
}

// PutTaskV1Handler godoc
// @Summary      Put a task by id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param		 id   path      int  true  "Task ID"
// @Param		 id   body      dto.PutTaskRequest  true  "Update task"
// @Param		 name  body      dto.PutTaskRequest  true  "Update task"
// @Param		 status  body      dto.PutTaskRequest  true  "Update task"
// @Success      200  {object}  dto.Response{result=model.Task}
// @Failure      400  {object}  dto.Response
// @Router       /task/{id} [put]
func PutTaskV1Handler(c *gin.Context) {
	id := c.Param(constant.PathID)

	req := dto.PutTaskRequest{}
	if _err := c.ShouldBindJSON(&req); _err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, _err.Error()))
		return
	}

	if id != strconv.Itoa(req.ID) {
		common.Logger.Error(c, fmt.Sprintf("ID not match: %s != %s", id, strconv.Itoa(req.ID)))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, constant.ErrIDNotMatch))
		return
	}

	task := &model.Task{
		ID:     req.ID,
		Name:   req.Name,
		Status: req.Status,
	}
	task, err := PutTask(c, task)
	if err != nil {
		common.Logger.Error(c, fmt.Sprintf("Put task %s failed: %s", id, err.ErrorMessage()))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, err.ErrorCode(), err.ErrorMessage()))
		return
	}

	c.JSON(http.StatusOK, dto.NewResponse(c, task))
}

// DeleteTaskV1Handler godoc
// @Summary      Delete a task by id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param		 id   path      int  true  "Task ID"
// @Success      200
// @Router       /task/{id} [delete]
func DeleteTaskV1Handler(c *gin.Context) {
	id := c.Param(constant.PathID)

	tid, _err := strconv.Atoi(id)
	if _err != nil {
		common.Logger.Error(c, fmt.Sprintf("Invalid ID: %s", id))
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(c, constant.ErrCodeBadRequest, _err.Error()))
		return
	}

	DeleteTask(c, tid)

	c.JSON(http.StatusOK, nil)
}
