package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"redenvelop-Prac/consts"
	"redenvelop-Prac/dal/db"
	"redenvelop-Prac/model"
	"redenvelop-Prac/utils"
	"time"
)

func ExportSendRecords(c *gin.Context) {
	// 1. 参数绑定
	var sReq model.ExportSendRecordReq
	err := c.BindJSON(&sReq)
	if err != nil {
		logrus.Error("[QuerySendRecords] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}
	// 2. 参数检查
	//ok := checkSendRecordParams(sReq)
	//if !ok {
	//	logrus.Errorf("[ReceiveRedPacket] check params error, rReq: %v", utils.Json2String(sReq))
	//	utils.RetErrJson(c, consts.ParamError)
	//	return
	//}

	records, rErr := db.ExportSendRecords(c, sReq)
	if rErr != nil {
		logrus.Errorf("[ExportSendRecords] query error %v", rErr)
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}

	count := len(records)
	exportSendRecords(c, records, count)
	if err != nil {
		utils.RetErrJson(c, consts.ServiceBusy)
	}

	utils.RetJson(c)
	// 返回数据
	result := make([]*model.RpSendRecord, 0)
	for _, record := range records {
		r := &model.RpSendRecord{
			UserId:      record.UserId,
			GroupChatId: record.GroupChatId,
			// todo 填充前端展示的数据
		}
		result = append(result, r)
	}

	retVal := &model.QuerySendRecordRespByPage{
		RpSendRecordList: result,
		Total:            count,
	}
	utils.RetJsonWithData(c, utils.Json2String(retVal))

}

func exportSendRecords(c *gin.Context, records []*model.RpSendRecord, count int) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet.
	sheetTittle := "send_record"
	index, err := f.NewSheet(sheetTittle)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//设置表头
	f.SetCellValue(sheetTittle, "A1", "用户id")
	f.SetCellValue(sheetTittle, "B1", "群聊id")
	f.SetCellValue(sheetTittle, "C1", "发放金额")
	f.SetCellValue(sheetTittle, "D1", "发放时间")

	// 设置数据
	dataIndex := 2
	for _, record := range records {
		f.SetCellValue(sheetTittle, fmt.Sprintf("A%d", dataIndex), record.UserId)
		f.SetCellValue(sheetTittle, fmt.Sprintf("B%d", dataIndex), record.GroupChatId)
		f.SetCellValue(sheetTittle, fmt.Sprintf("C%d", dataIndex), record.Amount)
		f.SetCellValue(sheetTittle, fmt.Sprintf("D%d", dataIndex), record.CreateTime) //todo 时间要换成固定模式
		dataIndex = dataIndex + 1
	}
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fmt.Sprintf("./../sendrecord/%d.xlsx", time.Now().Unix())); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
