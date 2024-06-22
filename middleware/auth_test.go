package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/andyliao/task-homework/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetGeneralContextMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(SetGeneralContextMiddleware())
	router.GET("/test", func(c *gin.Context) {
		traceID, exists := c.Get(constant.HeaderTraceID)
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "trace ID not set"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"trace_id": traceID})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	traceID, ok := response["trace_id"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, traceID)
	assert.Equal(t, len(traceID), len(util.GenerateTraceID()))
}

func TestAccessMiddleware_Authorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockIsAuthorized := func(ctx context.Context, key string) (*model.User, common.ICodeError) {
		return &model.User{Username: "test_user"}, nil
	}
	IsAuthorized = mockIsAuthorized

	router := gin.New()
	router.Use(AccessMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(constant.HeaderAPIKey, "valid_api_key")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "success"}`, w.Body.String())
}

func TestAccessMiddleware_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	errUnauthorized := common.NewCodeError(constant.ErrCodeUnauthorized, "unauthorized")
	mockIsAuthorized := func(ctx context.Context, key string) (*model.User, common.ICodeError) {
		return nil, errUnauthorized
	}
	IsAuthorized = mockIsAuthorized

	router := gin.New()
	router.Use(AccessMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(constant.HeaderAPIKey, "invalid_api_key")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"code": 100401, "message": "unauthorized"}`, w.Body.String())
}
