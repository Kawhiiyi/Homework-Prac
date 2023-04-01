package service

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExportSendRecords(t *testing.T) {
	// 创建一个 gin context
	router := gin.Default()
	router.POST("/export/send", ExportSendRecords)

	// 构造测试请求
	reqBody := `{"userId": 1}`
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/export/send", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// 测试函数
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)
}
