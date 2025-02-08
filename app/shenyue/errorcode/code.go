package errorcode

import "github.com/gin-gonic/gin"

// BuildDataResponse 构建包含错误码和错误信息的响应
func BuildDataResponse(ctx *gin.Context, data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	requestId := ctx.GetString("request_id")
	response["request_id"] = requestId
	response["code"] = 0
	response["message"] = "ok"
	response["data"] = data
	return response
}

// BuildErrorResponse 构建包含错误码和错误信息的响应
func BuildErrorResponse(ctx *gin.Context, errCode int64) map[string]interface{} {
	response := make(map[string]interface{})
	requestId := ctx.GetString("request_id")
	response["request_id"] = requestId
	response["code"] = errCode
	message := GetErrMsg(errCode)
	response["message"] = message
	return response
}

func GetErrMsg(code int64) string {
	return codeMsg[code]
}

const (
	SUCCSE = 200
	ERROR  = 500

	// common code= 1000... 1000 - 1999 通用模块的错误
	ErrParam = 1000

	// ERROR_USERNAME_USED code= 2000... 2000 - 2999 用户模块的错误
	ErrUsernameUsed  = 2001
	ErrPasswordWrong = 2002
	ErrUserNameWrong = 2003

	// Article

)

var codeMsg = map[int64]string{
	SUCCSE:           "OK",
	ERROR:            "FAIL",
	ErrParam:         "参数错误",
	ErrUsernameUsed:  "用户名已存在！",
	ErrPasswordWrong: "密码错误",
	ErrUserNameWrong: "用户名错误",
}
