package consts

type RError struct {
	Code int64
	Msg  string
}

func buildErrorCode(code int64, msg string) RError {
	return RError{
		Code: code,
		Msg:  msg,
	}
}

var (
	Success = buildErrorCode(0, "success")

	RecordNotFound = buildErrorCode(-1, "record not found")

	BindError = buildErrorCode(-2, "bind error")

	ParamsError = buildErrorCode(-5, "params error")

	InsertError = buildErrorCode(-3, "insert data error")

	ServiceBusy = buildErrorCode(-110, "service busy ")

	RpNotFoundError = buildErrorCode(-10001, "Rp record not found")

	RpStatusError = buildErrorCode(-6, "red packet status error")

	RpExpiredError = buildErrorCode(-6, "red packet expired")

	RpReceivedError = buildErrorCode(-4, "all the red packets are received")
)
