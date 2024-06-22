package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListTasksV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/tasks", ListTasksV1Handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Result)
}

func TestCreateTaskV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/tasks", CreateTaskV1Handler)

	taskName := "Test Task"
	reqBody, _ := json.Marshal(dto.CreateTaskRequest{Name: taskName})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Result)
	task := response.Result.(map[string]interface{})
	assert.Equal(t, taskName, task["name"])
}

func TestPutTaskV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create task first
	router.POST("/tasks", CreateTaskV1Handler)

	taskName := "Test Task"
	reqBody, _ := json.Marshal(dto.CreateTaskRequest{Name: taskName})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	router.PUT("/tasks/:"+constant.PathID, PutTaskV1Handler)

	// Update task
	tID := response.Result.(map[string]interface{})["id"].(float64)
	taskID := int(tID)
	taskName = "Updated Task"
	reqBody, _ = json.Marshal(dto.PutTaskRequest{ID: taskID, Name: taskName, Status: 0})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Result)
	task := response.Result.(map[string]interface{})
	assert.Equal(t, float64(taskID), task["id"])
	assert.Equal(t, taskName, task["name"])
	assert.Equal(t, float64(0), task["status"])

	// Update task with invalid ID
	taskID = 3
	taskName = "Updated Task"
	reqBody, _ = json.Marshal(dto.PutTaskRequest{ID: taskID, Name: taskName, Status: 0})
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteTaskV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/tasks/:"+constant.PathID, DeleteTaskV1Handler)

	taskID := 1
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "", response["result"])
}
