package responseCodeEnum

//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode             = 200  //成功
	UndefErrorCode          = 400  //为定义的错误
	InternalErrorCode       = 500  //内部错误
	InvalidRequestErrorCode = 401  //无效请求
	CustomizeCode           = 1000 //业务自定义错误码
	GROUPALL_SAVE_FLOWERROR = 2001
)
