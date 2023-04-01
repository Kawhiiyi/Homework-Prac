package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"redenvelop-Prac/consts"
	"redenvelop-Prac/model"
	"redenvelop-Prac/utils"
	"testing"
)

// 定义 Mock 对象
type mockDB struct {
	mock.Mock
}

func (m *mockDB) QueryReceiveRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo string, userId string) (*model.RpReceiveRecord, error) {
	args := m.Called(c, bizOutNo, userId)
	return args.Get(0).(*model.RpReceiveRecord), args.Error(1)
}

func (m *mockDB) QuerySendRecordByRpId(c *gin.Context, rpId string) (*model.RpSendRecord, error) {
	args := m.Called(c, rpId)
	return args.Get(0).(*model.RpSendRecord), args.Error(1)
}

func (m *mockDB) UpdateSendAndCreateReceiveRecordTx(c *gin.Context, sendRecord *model.RpSendRecord, receiveRecord *model.RpReceiveRecord) error {
	args := m.Called(c, sendRecord, receiveRecord)
	return args.Error(0)
}

// 定义测试用例
func TestReceiveRedPacket(t *testing.T) {
	// 构造测试数据
	rReq := model.ReceiveRpReq{
		BizOutNo: "123456",
		RpId:     001,
		UserId:   "test_user_id",
	}

	// 构造 Mock 对象
	mockDb := new(mockDB)
	mockDb.On("QueryReceiveRecordByBizOutNoAndUserId", mock.Anything, rReq.BizOutNo, rReq.UserId).Return(nil, nil)
	mockDb.On("QuerySendRecordByRpId", mock.Anything, rReq.RpId).Return(&model.RpSendRecord{Status: consts.RpStatusSend}, nil)
	mockDb.On("UpdateSendAndCreateReceiveRecordTx", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// 构造 gin.Context 对象
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/receive", utils.StringToReader(utils.Struct2Json(rReq)))
	c.Request.Header.Add("Content-Type", "application/json")

	// 执行函数
	ReceiveRedPacket(c)

	// 断言结果
	assert.Equal(t, http.StatusOK, w.Code)

	var result model.RpReceiveRecord
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Id)
	assert.NotEmpty(t, result.BizOutNo)
	assert.NotEmpty(t, result.UserId)
	assert.NotEmpty(t, result.Amount)
	assert.NotEmpty(t, result.CreateTime)
}
