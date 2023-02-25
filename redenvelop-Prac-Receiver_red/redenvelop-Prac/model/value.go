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
}
