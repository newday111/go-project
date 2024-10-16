package example

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"testing"
)

type User struct{}
type Address struct{}
type Product struct{}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("db1"), &gorm.Config{})

	//	使用dbresolver 中间件注册第一个配置, 这个配置指定了哪些数据库作为主库和从库
	db.Use(dbresolver.Register(dbresolver.Config{
		//	指定这两个库作为主库
		Sources: []gorm.Dialector{mysql.Open("db2"), mysql.Open("db3")},
		//	指定这两个库作为从库
		Replicas: []gorm.Dialector{mysql.Open("db4"), mysql.Open("db5")},
		//	选择随机策略来决定读操作时使用那个从库
		Policy: dbresolver.RandomPolicy{},
		//	开启日志记录, 记录读写分离的选择过程
		TraceResolverMode: true,
	}).Register(dbresolver.Config{
		//	注册第二个配置, 只针对User和Address表设置的从库
		//	指定db6作为 user和address表的从库, 默认db1作为主库
		Replicas: []gorm.Dialector{mysql.Open("db6")}}, &User{}, &Address{},
	).Register(dbresolver.Config{
		//	最后注册第三个配置，为 orders 和 Product 表指定不同的主从库
		Sources:  []gorm.Dialector{mysql.Open("db7"), mysql.Open("db8")},
		Replicas: []gorm.Dialector{mysql.Open("db9")}}, "orders", &Product{}, "secondary"))
	if err != nil {
		t.Fatal(err)
	}
}
