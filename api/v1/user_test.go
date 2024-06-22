package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andyliao/task-homework/dto"
	"github.com/andyliao/task-homework/module"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users", CreateUserV1Handler)

	userName := "test_user"
	userPassword := "password123"
	reqBody, _ := json.Marshal(dto.CreateUserRequest{Name: userName, Password: userPassword})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Result)
}

func TestLoginV1Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/login", LoginV1Hadnler)

	// Creating a user first
	userName := "test_user"
	userPassword := "password123"
	module.UserModule.CreateUser(context.Background(), userName, userPassword)

	reqBody, _ := json.Marshal(dto.CreateUserRequest{Name: userName, Password: userPassword})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response.Result)
	authKey := response.Result.(string)
	assert.NotEmpty(t, authKey)
}
