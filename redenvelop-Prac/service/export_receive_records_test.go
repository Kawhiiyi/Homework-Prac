package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExportReceiveRecords(t *testing.T) {
	// 创建一个 gin context
	router := gin.Default()
	router.POST("/export", ExportReceiveRecords)

	// 构造测试请求
	reqBody := `{"userId": 1}`
	req, _ := http.NewRequest(http.MethodPost, "/export", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// 测试函数
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)
}
