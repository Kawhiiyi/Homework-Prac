package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"redenvelop-Prac/consts"
	"redenvelop-Prac/model"
	"redenvelop-Prac/utils"
	"strconv"
)

func QueryReceiveRecords(c *gin.Context) {
	// 1. 参数绑定
	var rReq model.QueryReceiveRecordReq
	err := c.BindJSON(&rReq)
	if err != nil {
		logrus.Error("[QueryReceiveRecords] bind req json error")
		utils.RetErrorJson(c, consts.BindError)
		return
	}

	// 2. 参数检查
	ok := checkReceiveRecordParams(rReq)
	if !ok {
		logrus.Errorf("[QueryReceiveRecords] check params error, rReq: %v", utils.Json2String(rReq))
		utils.RetErrorJson(c, consts.ParamsError)
		return
	}

	// 3. 查数据 重点：找到cursor对应的字段
	if rReq.Cursor == "" {
		rReq.Cursor = strconv.FormatInt(math.MaxInt64, 10)
	}
	if rReq.Size == 0 {
		rReq.Size = 10
	}
	temp := rReq.Size
	rReq.Size = rReq.Size + 1

	hasMore := false
	records, rErr := db.QueryReceiveRecordByCond(c, rReq)
	if rErr != nil {
		// todo
	}
	if len(records) > int(temp) {
		hasMore = true
	}

	// 返回
	result := make([]*model.RpReceiveRecord, 0)
	curInt := 0
	for i, record := range records {
		r := &model.RpReceiveRecord{
			UserId: record.UserId,
			// todo 填充前端展示的数据
		}
		curInt = i
		result = append(result, r)
	}
	retVal := &model.QueryReceiveRecordResp{
		RpReceiveRecordList: result,
		HasMore:             hasMore,
		Cursor:              strconv.FormatInt(records[curInt].CreateTime.Unix(), 10),
	}
	utils.RetJsonWithData(c, utils.Json2String(retVal))
}

func checkReceiveRecordParams(req model.RpReceivePacketReq) bool {
	if req.UserId == "" {
		return false
	}
	if req.RpId == "" {
		return false
	}
	if req.GroupChatId == "" {
		return false
	}
	return true
}

func QueryReceiveRecordsByPage(c *gin.Context) {
	// 1. 绑定请求参数
	var req model.QueryReceiveRecordReqByPage
	if err := c.BindJSON(&req); err != nil {
		utils.RetErrorJson(c, consts.BindError)
		return
	}

	// 2. 检查参数
	if err := checkReceiveRecordParams(req); err != nil {
		utils.RetErrorJson(c, consts.ParamsError)
		return
	}

	// 3. 执行查询
	req.Page = req.Page - 1
	records, err := db.QueryReceiveRecordByCondPage(c, req)
	if err != nil {
		// todo
	}

	count, err := db.CountReceiveRecordByCondPage(c, req)
	if err != nil {
		// todo
	}

	// 4. 返回结果
	result := make([]*model.RpReceiveRecord, 0, len(records))
	for _, record := range records {
		r := &model.RpReceiveRecord{
			UserId: record.UserId,
			RpId:   record.RpId,
			Amount: record.Amount,
			Time:   record.Time,
		}
		result = append(result, r)
	}

	resp := &model.QueryReceiveRecordRespByPage{
		RpReceiveRecordList: result,
		Total:               count,
	}
	utils.RetJsonWithData(c, utils.Json2String(resp))
}

func checkReceiveRecordParams(req model.QueryReceiveRecordReqByPage) error {
	if req.UserId == "" {
		return errors.New("missing user id")
	}
	return nil
}
