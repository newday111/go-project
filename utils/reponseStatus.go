package utils

const (
	RespSuccess         = 200
	RespParamsError     = 400
	RespAuthFail        = 401
	RespCreateTokenFail = 402
	RespFail            = 500
)

var RespMsg = map[int]string{
	200: "登录成功",
	400: "参数错误",
	401: "没有登录,无权限访问该信息",
	402: "登录token生成失败",
	500: "服务器错误",
}
