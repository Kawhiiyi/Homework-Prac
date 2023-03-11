package db

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"redenvelop-Prac/model"
	"strconv"
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

func QueryReceiveRecordByCond(c *gin.Context, req model.QueryReceiveRecordReq) ([]*model.RpReceiveRecord, error) {
	records := make([]*model.RpReceiveRecord, 0)
	// select * from receive_record where user_id = xxx and group_chat_id = xxx and create_time < cursor order by create_time desc limit size
	tx := rdb.Table(tableNameReceive).WithContext(c).Where("user_id = ?", req.UserId).Where("group_chat_id = ?", req.GroupId).Where("create_time < ?", req.Cursor)

	err := tx.Order("create_time desc").Limit(int(req.Size)).Find(&records).Error
	if err != nil {
		logrus.Errorf("dal.QueryReceiveRecordByCond query error %v", err)
		return nil, err
	}
	return records, nil
}

func QueryReceiveRecordByCondPage(c *gin.Context, req model.QueryReceiveRecordReqByPage) ([]*model.RpReceiveRecord, error) {
	records := make([]*model.RpReceiveRecord, 0)

	// select * from receive_record where user_id = xxx and (group_chat_id = ?) and create_time < ? order by create_time desc limit size
	tx := rdb.Table(tableNameReceive).WithContext(c).Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx = tx.Where("group_chat_id = ?", req.GroupId)
	}

	cursor, err := strconv.ParseInt(req.Cursor, 10, 64)
	if err != nil {
		logrus.Errorf("dal.QueryReceiveRecordByCondPage parse cursor error %v", err)
		return nil, err
	}
	tx = tx.Where("create_time < ?", cursor)

	err = tx.Order("create_time desc").Limit(int(req.Size)).Find(&records).Error
	if err != nil {
		logrus.Errorf("dal.QueryReceiveRecordByCondPage query error %v", err)
		return nil, err
	}
	return records, nil
}

func ExportReceiveRecords(c *gin.Context, req model.ExportReceiveRecordReq) ([]*model.RpReceiveRecord, error) {
	records := make([]*model.RpReceiveRecord, 0)
	// select * from receive_record where user_id = xxx and (group_chat_id = ?) order by receive_time desc limit size
	tx := rdb.Table(tableNameReceive).Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx.Where("group_chat_id = ?", req.GroupId)
	}

	err := tx.Order("receive_time desc").Find(&records).Error
	if err != nil {
		logrus.Errorf("db.ExportReceiveRecords query error %v", err)
		return nil, err
	}

	return records, nil
}

func CountReceiveRecordByCondPage(c *gin.Context, req model.QueryReceiveRecordReqByPage) (int64, error) {
	count := int64(0)

	// select count(*) from receive_record where user_id = xxx and (group_chat_id = ?) and create_time < ?
	tx := rdb.Table(tableNameReceive).WithContext(c).Select("count(*)").Where("user_id = ?", req.UserId)
	if req.GroupId != "" {
		tx = tx.Where("group_chat_id = ?", req.GroupId)
	}

	cursor, err := strconv.ParseInt(req.Cursor, 10, 64)
	if err != nil {
		logrus.Errorf("dal.CountReceiveRecordByCondPage parse cursor error %v", err)
		return 0, err
	}
	tx = tx.Where("create_time < ?", cursor)

	err = tx.Count(&count).Error
	if err != nil {
		logrus.Errorf("dal.CountReceiveRecordByCondPage query error %v", err)
		return 0, err
	}

	return count, nil
}

func InsertReceiveRecord(c *gin.Context, record *model.RpReceiveRecord) (int64, error) {
	err := rdb.Table(tableNameReceive).WithContext(c).Create(record).Error
	// err有两种情况 1. 数据库有问题   2. 数据插入重复
	if err != nil {
		logrus.Errorf("dal.InsertReceiveRecord error %v", err)
		return 0, err
	}
	return record.Id, err
}

func QueryReceiveRecordByBizOutNoAndUserId(c *gin.Context, bizOutNo, userId string) (*model.RpReceiveRecord, error) {
	var record model.RpReceiveRecord
	// find 和first的区别：record not find报错--first；find不报错
	err := rdb.Table(tableNameReceive).WithContext(c).Where("user_id = ?", userId).Where("biz_out_no = ?", bizOutNo).First(&record).Error
	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logrus.Errorf("dal.QuerySendRecordByBizOutNoAndUserId query error %v", err)
		return nil, err
	}
	return &record, nil
}

func InsertRecord(record *model.RpReceiveRecord) (int64, error) {
	err := rdb.Table(tableNameReceive).Create(&record).Error
	if err != nil {
		log.Printf("insert data err: %v\n", err)
		return 0, err
	}
	return record.Id, nil

}
