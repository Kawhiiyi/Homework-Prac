package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"redenvelop-Prac/consts"
	"redenvelop-Prac/dal"
	"redenvelop-Prac/model"
	"redenvelop-Prac/service/strategy"
	"redenvelop-Prac/utils"
	"strings"
	"time"
)

func SendRedPacket(c *gin.Context) {
	var sReq model.SendRpReq
	err := c.BindJSON(&sReq)
	if err != nil {
		logrus.Error("[SendRedPacket] bind req json error")
		utils.RetErrorJson(c, consts.BindError)
	}
	ok := checkParams(sReq)
	if !ok {
		logrus.Warnf("[SendRedPacket] check params error, sReq:%v", utils.Json2String(sReq))
		utils.RetErrorJson(c, consts.ParamsError)
	}

	if sReq.UserId == sReq.GroupId {
		utils.RetErrorJson(c, consts.ParamsError)
	}

	record, rErr := dal.QueryRecordByBizOutNoAndUserId(sReq.BizOutNo, sReq.UserId)
	if rErr != nil {
		logrus.Error("[SendRedPacket] query db error %v", err)
		utils.RetErrorJson(c, consts.ServiceBusy)

	}

	if record != nil {
		logrus.Infof("[SendRedPacket] biz out has one record alreadty")
		utils.RetJsonWithData(c, utils.Json2String(record))
	}

	var newRecord model.RpSendRecord
	newRecord.RpId = strings.ReplaceAll(uuid.New().String(), "-", " ")

	newRecord.SendTime = time.Now()
	newRecord.ExpireTime = time.Now().Add(consts.ExpireTime24)

	sMap := map[string][]int64{}
	var receiveAmountList []int64

	remain := sReq.Amount
	sum := int64(0)
	for i := int64(0); i < sReq.Number; i++ {
		x := strategy.DoubleAverage(sReq.Number-1, remain)
		receiveAmountList = append(receiveAmountList, x)
		remain -= x
		sum += x
	}
	sMap[newRecord.RpId] = receiveAmountList

	buildSendRecord(newRecord, sReq)
	id, dErr := dal.InsertSendRecord(newRecord)
	if dErr != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dErr, &mysqlErr) && mysqlErr.Number == 1062 {

			oldRecord, oErr := dal.QueryRecordByBizOutNoAndUserId(sReq.BizOutNo, sReq.UserId)
			if oErr != nil {
				logrus.Error("[SendRedPacket]:old record query db error %v", err)
				utils.RetErrorJson(c, consts.ServiceBusy)
			}

			if oldRecord != nil {
				logrus.Infof("[SendRedPacket]:bizoutno has recorded one already ")
				utils.RetJsonWithData(c, utils.Json2String(record))
			}
		} else {

			logrus.Warnf("[SendRedPacket]:bizOutNo has one record already")
			utils.RetErrorJson(c, consts.InsertError)
		}
	}
	logrus.Infof("[SendRedPacket]:insert rp record successfully, auto increase id is : %v", id)

	if amountListInMap, okk := sMap[newRecord.RpId]; okk {
		var total int64
		for _, val := range amountListInMap {
			total += val
		}
		if total == sReq.Amount {
			logrus.Infof("[SendRedPacket]: amountListInMap equals user amount")
		} else {
			// 回滚数据库,删除发放记录，把钱反给用户

			//报错
			utils.RetErrorJson(c, consts.ServiceBusy)

		}
	}
}

func buildSendRecord(record model.RpSendRecord, req model.SendRpReq) {
	record.UserId = req.UserId
	record.GroupChatId = req.GroupId
	record.BizOutNo = req.BizOutNo
	record.Amount = req.Amount
	record.ReceiveAmount = 0
	record.Number = req.Number
	record.Status = consts.RpStatusSend
	record.CreateTime = time.Now()
	record.ModifyTime = time.Now()
}

func QuerySendRecords(c *gin.Context) {

}

func ReceiveRedPacket(c *gin.Context) {

}

func checkParams(seq model.SendRpReq) bool {
	return !(seq.UserId == "" || seq.GroupId == "" || seq.Amount <= 0 || (seq.Number*seq.Amount) <= 1 || seq.BizOutNo == "")
}
func QueryReceiveRecords(c *gin.Context) {

}
