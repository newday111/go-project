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
	UserCreateFailed         = 406
	UserRegisterFailed       = 407
)

var UserRespMsg = map[int]string{
	200: "登录成功",
	300: "注册成功",
	401: "该用户没有注册,请注册之后在登录",
	402: "用户创建方式失败,请重新登录",
	403: "解析前端参数错误",
	405: "请重新登录",
	406: "注册用户失败, 用户名或邮箱已经注册",
	407: "用户注册失败",
	//401: "没有登录,无权限访问该信息",
	//402: "登录token生成失败",
	//500: "服务器错误",
}

const (
	VerificationCodeSendSuccess = 600
	VerificationCodeExist       = 601
)

var VerificationCodeRespMsg = map[int]string{
	600: "验证码发送成功",
	601: "验证码已经存在",
}
