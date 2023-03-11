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

func ExportReceiveRecords(c *gin.Context) {
	// 1. 参数绑定
	var rReq model.ExportReceiveRecordReq
	err := c.BindJSON(&rReq)
	if err != nil {
		logrus.Error("[ExportReceiveRecords] bind req json error")
		utils.RetErrJson(c, consts.BindError)
		return
	}

	// 2. 查询领取记录
	records, rErr := db.ExportReceiveRecords(c, rReq)
	if rErr != nil {
		logrus.Errorf("[ExportReceiveRecords] query error %v", rErr)
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}

	count := len(records)

	// 3. 导出到Excel
	err = exportReceiveRecords(c, records, count)
	if err != nil {
		logrus.Errorf("[ExportReceiveRecords] export error %v", err)
		utils.RetErrJson(c, consts.ServiceBusy)
		return
	}

	// 4. 返回数据
	result := make([]*model.RpReceiveRecord, 0)
	for _, record := range records {
		r := &model.RpReceiveRecord{
			UserId:      record.UserId,
			RpId:        record.RpId,
			Amount:      record.Amount,
			ReceiveTime: record.CreateTime,
		}
		result = append(result, r)
	}

	retVal := &model.QueryReceiveRecordRespByPage{
		RpReceiveRecordList: result,
		Total:               int64(count),
	}
	utils.RetJsonWithData(c, utils.Json2String(retVal))
}

func exportReceiveRecords(c *gin.Context, records []*model.RpReceiveRecord, count int) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建新的表格
	sheetTitle := "receive_record"
	index, err := f.NewSheet(sheetTitle)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 设置表头
	f.SetCellValue(sheetTitle, "A1", "用户id")
	f.SetCellValue(sheetTitle, "B1", "红包id")
	f.SetCellValue(sheetTitle, "C1", "领取金额")
	f.SetCellValue(sheetTitle, "D1", "领取时间")

	// 填充数据
	dataIndex := 2
	for _, record := range records {
		f.SetCellValue(sheetTitle, fmt.Sprintf("A%d", dataIndex), record.UserId)
		f.SetCellValue(sheetTitle, fmt.Sprintf("B%d", dataIndex), record.RpId)
		f.SetCellValue(sheetTitle, fmt.Sprintf("C%d", dataIndex), record.Amount)
		f.SetCellValue(sheetTitle, fmt.Sprintf("D%d", dataIndex), record.CreateTime.Format("2023-01-02 15:04:05"))
		dataIndex++
	}
	f.SetActiveSheet(index)

	// 保存文件
	if err := f.SaveAs(fmt.Sprintf("./../receiverecord/%d.xlsx", time.Now().Unix())); err != nil {
		fmt.Println(err)

	}
	return nil
}
