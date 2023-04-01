package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"redenvelop-Prac/model"
	"strings"
	"testing"
)

func TestQuerySendRecords(t *testing.T) {
	// 构造参数
	sReq := &model.QuerySendRecordReq{
		UserId: "",
		Cursor: "",
		Size:   0,
	}
	jsonStr, _ := json.Marshal(sReq)
	reqBody := strings.NewReader(string(jsonStr))

	// 构造请求
	req, err := http.NewRequest(http.MethodPost, "/querySendRecords", reqBody)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// 构造响应
	w := httptest.NewRecorder()

	// 构造上下文
	r := gin.Default()
	r.POST("/querySendRecords", QuerySendRecords)
	r.ServeHTTP(w, req)

	// 验证结果
	respBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "unexpected status code")
	var resp model.QuerySendRecordResp
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, false, resp.HasMore, "unexpected HasMore value")
	assert.Equal(t, "9223372036854775807", resp.Cursor, "unexpected Cursor value")
	assert.Equal(t, 0, len(resp.RpSendRecordList), "unexpected RpSendRecordList length")
}

func TestQuerySendRecordsByPage(t *testing.T) {
	// 构造参数
	sReq := &model.QuerySendRecordReqByPage{
		UserId: "",
		Page:   0,
		Size:   0,
	}
	jsonStr, _ := json.Marshal(sReq)
	reqBody := strings.NewReader(string(jsonStr))

	// 构造请求
	req, err := http.NewRequest(http.MethodPost, "/querySendRecordsByPage", reqBody)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// 构造响应
	w := httptest.NewRecorder()

	// 构造上下文
	r := gin.Default()
	r.POST("/querySendRecordsByPage", QuerySendRecordsByPage)
	r.ServeHTTP(w, req)

	// 验证结果
	respBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code, "unexpected status code")
	var resp model.QuerySendRecordRespByPage
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, 0, len(resp.RpSendRecordList), "unexpected RpSendRecordList length")
	assert.Equal(t, int(math.MaxInt64), resp.Total, "unexpected Total value")
}
