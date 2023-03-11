package model

type PackageInfo struct {
	UserId        string `json:"user_id"`
	ReceiveAmount int64  `json:"receive_amount"`
}

type SendRpReq struct {
	UserId   string `json:"user_id"`
	GroupId  string `json:"group_id"`
	Amount   int64  `json:"amount"`
	Number   int64  `json:"number"`
	BizOutNo string `json:"biz_out_no"`
}

type ReceiveRpReq struct {
	UserId   string `json:"user_id"`
	GroupId  string `json:"group_id"`
	BizOutNo string `json:"biz_out_no"`
	Amount   int64  `json:"amount"`
	Number   int64  `json:"number"`
	RpId     int64  `json:"rp_id"`
}

type QuerySendRecordReq struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`

	Size   int64  `json:"size"`   // 默认10条
	Cursor string `json:"cursor"` // 默认从"0"
}

type RpReceivePacketReq struct {
	UserId      string `json:"user_id"`
	RpId        string `json:"rp_id"`
	GroupChatId string `json:"group_chat_id"`
	Amount      int64  `json:"amount"`
	Cursor      string `json:"cursor" binding:"required"`
}

type QuerySendRecordReqByPage struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
	Size    int64  `json:"size"` // 默认20条
	Page    int64  `json:"page"` // 从第1页开始，默认是1
	Total   int64  `json:"total"`
}

type ExportSendRecordReq struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}

type ExportReceiveRecordReq struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
}

type QueryReceiveRecordReq struct {
	UserId  string `json:"user_id" binding:"required"`
	GroupId string `json:"group_id" binding:"required"`
	Cursor  string `json:"cursor" binding:"required"`
	Size    int    `json:"size"`
}

type QueryReceiveRecordResp struct {
	RpReceiveRecordList []*RpReceiveRecord `json:"rp_receive_record_list"` // 红包领取记录列表
	HasMore             bool               `json:"has_more"`               // 是否还有更多数据
	Cursor              string             `json:"cursor"`                 // 游标，用于下一页查询
}

type QueryReceiveRecordReqByPage struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
	Page    int64  `json:"page"`
	Size    int64  `json:"size"`
	Cursor  string `json:"cursor" binding:"required"`
	Total   int64  `json:"total"`
}

type QuerySendRecordRespByPage struct {
	RpSendRecordList []*RpSendRecord `json:"rp_send_record_list"`
	Total            int             `json:"total"`
}

type QueryReceiveRecordRespByPage struct {
	RpReceiveRecordList []*RpReceiveRecord `json:"rp_receive_record_list"`
	Total               int64              `json:"total"`
}

type QuerySendRecordResp struct {
	RpSendRecordList []*RpSendRecord `json:"rp_send_record_list"`
	HasMore          bool            `json:"has_more"`
	Cursor           string          `json:"cursor"`
}
