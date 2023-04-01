package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"redenvelop-Prac/model"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestQueryReceiveRecords(t *testing.T) {
	// 创建一个 gin context
	router := gin.Default()
	router.POST("/query", QueryReceiveRecords)

	// 构造测试请求
	reqBody := `{"userId": "1", "groupId": "2", "cursor": "12345", "size": 10}`
	req, _ := http.NewRequest(http.MethodPost, "/query", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// 测试函数
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)

	// 解析响应结果
	var resp model.QueryReceiveRecordResp
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.RpReceiveRecordList, 0)
	assert.False(t, resp.HasMore)
	assert.Empty(t, resp.Cursor)
}

func TestQueryReceiveRecordsByPage(t *testing.T) {
	// 创建一个 gin context
	router := gin.Default()
	router.POST("/query/page", QueryReceiveRecordsByPage)

	// 构造测试请求
	reqBody := `{"userId": "1", "groupId": "2", "page": 1, "size": 10}`
	req, _ := http.NewRequest(http.MethodPost, "/query/page", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// 测试函数
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)

	// 解析响应结果
	var resp model.QueryReceiveRecordRespByPage
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp.RpReceiveRecordList, 0)
	assert.Equal(t, int64(0), resp.Total)
}
