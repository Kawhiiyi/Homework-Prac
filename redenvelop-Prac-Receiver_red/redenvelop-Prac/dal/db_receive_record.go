package dal

import (
	"github.com/sirupsen/logrus"
	"log"
	"redenvelop-Prac/model"
)

var tableNameReceive = "rp_receive_record"

func QueryByUserId(userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	err := rdb.Table(tableNameReceive).Where("user_id= ? ", userId).First(&record).Error
	if err != nil {
		log.Printf("Can't find userId amount. userId : %v err : %v", userId, err)
		return nil, err
	}
	return &record, nil
}

func InsertReceiveRecord(record model.RpReceiveRecord) (int64, error) {
	err := rdb.Table(tableNameReceive).Create(&record).Error
	if err != nil {
		logrus.Errorf("dal.InsertReceiveRecord error: %v", err)
		return 0, err
	}
	return record.Id, err
}
