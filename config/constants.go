package config

const (
	// Success 成功
	Success = 0
	// ERROR 错误
	ERROR = 1
	// Param Invalidate 参数
	InvalidParams = -2001
)

// MsgFlags 存放msg
var MsgFlags = map[int]string{
	Success:       "success",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",
}

// GetMsg根据code获取msg信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return ""
}
