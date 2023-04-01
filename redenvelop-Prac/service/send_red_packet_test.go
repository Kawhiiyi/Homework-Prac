package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"redenvelop-Prac/model"
	"redenvelop-Prac/utils"
	"testing"
)

// 定义 Mock 对象
func TestSendRedPacket(t *testing.T) {
	// 构造测试数据
	sReq := model.SendRpReq{
		UserId: "test_user_id",
		Number: 10,
		Amount: 100,
	}

	// 构造 Mock 对象
	mockDb := new(mockDB)
	mockDb.On("UpdateAccountTx", mock.Anything, mock.Anything).Return(nil)

	// 构造 gin.Context 对象
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/send", utils.StringToReader(utils.Struct2Json(sReq)))
	c.Request.Header.Add("Content-Type", "application/json")

	// 执行函数
	SendRedPacket(c)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)

	var result model.RpSendRecord
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Id)
	assert.NotEmpty(t, result.RpId)
	assert.NotEmpty(t, result.UserId)
	assert.NotEmpty(t, result.Number)
	assert.NotEmpty(t, result.Amount)
	assert.NotEmpty(t, result.CreateTime)
}
