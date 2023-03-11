package db

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"redenvelop-Prac/model"
	"testing"
)

func TestQueryByUserId(t *testing.T) {
	// 创建一个新的 gin 上下文（context）
	// 测试用例 1：存在的 user_id
	record, err := QueryByUserId("1")
	assert.NoError(t, err)
	assert.NotNil(t, record)

	// 测试用例 2：不存在的 user_id
	record, err = QueryByUserId("not_exist_user_id")
	assert.Error(t, err)
	assert.Nil(t, record)
}

func TestQueryReceiveRecordByCond(t *testing.T) {
	// 创建一个新的 gin 上下文（context）
	c := gin.Context{}

	// 创建一个测试用例请求参数
	req := model.QueryReceiveRecordReq{
		UserId:  "1",
		GroupId: "group_1",
		Cursor:  "100",
		Size:    10,
	}

	// 测试用例 1：符合条件的记录存在
	records, err := QueryReceiveRecordByCond(&c, req)
	assert.NoError(t, err)
	assert.NotNil(t, records)

	// 测试用例 2：符合条件的记录不存在
	req.GroupId = "not_exist_group_id"
	records, err = QueryReceiveRecordByCond(&c, req)
	assert.NoError(t, err)
	assert.Empty(t, records)
}
