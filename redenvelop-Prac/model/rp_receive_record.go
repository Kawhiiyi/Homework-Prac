package model

import "time"

type RpReceiveRecord struct {
	Id          int64
	UserId      string
	GroupChatId string
	Amount      int64
	password    string
	RpId        string
	CreateTime  time.Time
	ModifyTime  time.Time
}
