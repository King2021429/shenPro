package errorcode

func GetErrMsg(code int64) string {
	return codeMsg[code]
}

const (
	SUCCSE = 200
	ERROR  = 500

	// ERROR_USERNAME_USED code= 1000... 用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
)

var codeMsg = map[int64]string{
	SUCCSE:               "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误",
}
