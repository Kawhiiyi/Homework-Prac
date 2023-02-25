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

// 发红包接口
func SendRedPacket(c *gin.Context) {
	//绑定请求到结构体
	var sReq model.SendRpReq
	err := c.BindJSON(&sReq)
	if err != nil {
		logrus.Error("[SendRedPacket] bind req json error")
		utils.RetErrorJson(c, consts.BindError)
	}
	// 校验参数
	ok := checkParams(sReq)
	if !ok {
		logrus.Warnf("[SendRedPacket] check params error, sReq:%v", utils.Json2String(sReq))
		utils.RetErrorJson(c, consts.ParamsError)
	}

	// 发红包的用户和群聊不能相同

	if sReq.UserId == sReq.GroupId {
		utils.RetErrorJson(c, consts.ParamsError)
	}
	// 根据业务订单号和用户ID查询发红包记录

	record, rErr := dal.QueryRecordByBizOutNoAndUserId(sReq.BizOutNo, sReq.UserId)
	if rErr != nil {
		logrus.Error("[SendRedPacket] query db error %v", err)
		utils.RetErrorJson(c, consts.ServiceBusy)

	}
	// 若发红包记录已存在，则返回之前的记录

	if record != nil {
		logrus.Infof("[SendRedPacket] biz out has one record alreadty")
		utils.RetJsonWithData(c, utils.Json2String(record))
	}
	// 生成红包发送记录

	var newRecord model.RpSendRecord
	newRecord.RpId = strings.ReplaceAll(uuid.New().String(), "-", " ")

	newRecord.SendTime = time.Now()
	newRecord.ExpireTime = time.Now().Add(consts.ExpireTime24)
	// 计算红包金额分配情况
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

	// 构建发送红包记录
	buildSendRecord(newRecord, sReq)
	// 插入发送红包记录到数据库
	id, dErr := dal.InsertSendRecord(newRecord)
	if dErr != nil {
		var mysqlErr *mysql.MySQLError
		// 若插入记录因为唯一索引约束失败，则说明之前已经插入过相同业务订单号的记录，返回之前的记录
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
	// 绑定接收到的JSON请求数据到结构体中
	var rReq model.ReceiveRpReq
	err := c.BindJSON(&rReq)
	if err != nil {
		logrus.Error("[ReceiveRedPacket] bind req json error")
		utils.RetErrorJson(c, consts.BindError)
		return
	}

	// 构造发红包的请求结构体
	sReq := model.SendRpReq{
		UserId:   rReq.UserId,
		GroupId:  rReq.GroupId,
		Amount:   rReq.Amount,
		Number:   rReq.Number,
		BizOutNo: rReq.BizOutNo,
	}

	// 校验参数是否合法
	ok := checkParams(sReq)
	if !ok {
		logrus.Warnf("[ReceiveRedPacket] check params error, rReq:%v", utils.Json2String(rReq))
		utils.RetErrorJson(c, consts.ParamsError)
		return
	}

	// 根据商户订单号和用户ID查询发送红包记录
	record, rErr := dal.QueryRecordByBizOutNoAndUserId(rReq.BizOutNo, rReq.UserId)
	if rErr != nil {
		logrus.Error("[ReceiveRedPacket] query db error %v", err)
		utils.RetErrorJson(c, consts.ServiceBusy)
		return
	}
	if record == nil {
		logrus.Infof("[ReceiveRedPacket] cannot find rp record, bizOutNo:%s, userId:%s", rReq.BizOutNo, rReq.UserId)
		utils.RetErrorJson(c, consts.RpNotFoundError)
		return
	}
	// 校验发送红包记录的状态，只有状态为已发送时才能继续领取红包
	if record.Status != consts.RpStatusSend {
		logrus.Infof("[ReceiveRedPacket] rp record has been received or expired, status:%d", record.Status)
		utils.RetErrorJson(c, consts.RpStatusError)
		return
	}
	// 校验红包是否已过期
	if time.Now().After(record.ExpireTime) {
		logrus.Infof("[ReceiveRedPacket] rp record has been expired")
		utils.RetErrorJson(c, consts.RpExpiredError)
		return
	}
	// 根据红包ID在红包接收记录Map中查询当前用户领取的金额列表

	sMap := map[string][]int64{}
	amountListInMap, ok := sMap[record.RpId]
	if !ok {
		logrus.Errorf("[ReceiveRedPacket] sMap does not contain rpId: %s", record.RpId)
		utils.RetErrorJson(c, consts.ServiceBusy)
		return
	}

	// 校验当前用户是否已领取所有红包，如果是，则返回领取红包错误
	if record.Number <= int64(len(amountListInMap)) {
		logrus.Errorf("[ReceiveRedPacket] all red packets have been received")
		utils.RetErrorJson(c, consts.RpReceivedError)
		return
	}

	// 计算当前用户领取的金额，使用策略模式计算红包金额
	receiveIndex := len(amountListInMap)
	receiveAmount := strategy.RandomAmount(record.Amount, record.ReceiveAmount, (record.Number)-int64(receiveIndex))
	amountListInMap = append(amountListInMap, receiveAmount)
	sMap[record.RpId] = amountListInMap

	receiveRecord := model.RpReceiveRecord{
		RpId:        record.RpId,
		ReceiveTime: time.Now(),
		UserId:      rReq.UserId,
		Amount:      receiveAmount,
	}
	dal.InsertReceiveRecord(receiveRecord)
	record.ReceiveAmount += receiveAmount
	record.ModifyTime = time.Now()
	if record.ReceiveAmount == record.Amount {
		record.Status = consts.RpStatusReceived
	}

	//dal.UpdateSendRecord(record)

	utils.RetJsonWithData(c, utils.Json2String(receiveRecord))
}

func checkParams(seq model.SendRpReq) bool {
	return !(seq.UserId == "" || seq.GroupId == "" || seq.Amount <= 0 || (seq.Number*seq.Amount) <= 1 || seq.BizOutNo == "")
}

func QueryReceiveRecords(c *gin.Context) {

}
