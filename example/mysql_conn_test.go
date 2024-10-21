package example

import (
	"goPro4/db"
	"testing"
)

func TestGetMysqlConn(t *testing.T) {
	//db.GetMysqlConnection()
	//	测试的时候这样启动,正式环境需要再服务器的配置文件中配置用户和密码
	// redis-server --protected-mode no

	db.GetRedisConnection()
}
