package errorcode

func GetErrMsg(code int) string {
	return codeMsg[code]
}

var codeMsg = map[int]string{
	SUCCSE:               "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误",
}
