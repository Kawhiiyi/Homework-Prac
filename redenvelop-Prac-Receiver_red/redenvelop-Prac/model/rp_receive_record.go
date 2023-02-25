package model

import "time"

type RpReceiveRecord struct {
	Id            int64
	UserId        string
	GroupChatId   string
	Amount        int64
	password      string
	RpId          string
	BizOutNo      string
	ReceiveAmount int64
	Number        int64
	Status        int
	ExpireTime    time.Time
	ReceiveTime   time.Time
	CreateTime    time.Time
	ModifyTime    time.Time
}
