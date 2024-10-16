package utils

import "net/http"

const (
	RecordLogSuccess     = http.StatusOK
	RecordLogUpdateFail  = http.StatusRequestURITooLong
	RecordLogCreateFail  = http.StatusUnsupportedMediaType
	RecordLogDeleteFail  = http.StatusRequestedRangeNotSatisfiable
	RecordLogGetFail     = http.StatusRequestHeaderFieldsTooLarge
	RecordLogTokenExpire = http.StatusNetworkAuthenticationRequired
)

const (
	UserLoginSuccess         = 200
	UserRegisterSuccess      = 300
	UserQueryFailed          = 401
	UserTokenCreateFailed    = 402
	UserParamsAnalysisFailed = 403
	UserTokenExpire          = 405
	//RespGetDataSuccess  = 201
	//RespParamsError     = 400
	//RespAuthFail        = 401
	//RespCreateTokenFail = 402
	//RespFail            = 500
)

var UserRespMsg = map[int]string{
	200: "登录成功",
	300: "注册成功",
	401: "该用户没有注册,请注册之后在登录",
	402: "用户创建方式失败,请重新登录",
	403: "解析前端参数错误",
	405: "请重新登录",
	//401: "没有登录,无权限访问该信息",
	//402: "登录token生成失败",
	//500: "服务器错误",
}
