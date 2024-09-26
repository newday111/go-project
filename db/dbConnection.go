package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func GetMysqlConnection() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("E:\\goPro\\goPro4\\config\\") // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: %s \n", err)
	}

	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	dbName := viper.GetString("mysql.dbname")

	mysqlMaxConns := viper.GetInt("mysql.max_open_connects")
	mysqlIdeConns := viper.GetInt("mysql.max_idle_connects")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)

	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   false, //	设置为false表示每次数据库操作时, 默认使用事务处理
		NamingStrategy:                           nil,   // 设置为nil表示使用默认的命名策略,Gorm的默认命名策略是驼峰命名，例如：UserName -> user_name
		FullSaveAssociations:                     false, // 设置为false表示在保存关联关系时不会递归保存所有关联对象
		Logger:                                   nil,   // 设置为nil表示使用默认的日志记录器
		NowFunc:                                  nil,   // 设置为nil表示使用系统当前时间作为NOW()函数的实现
		DryRun:                                   false, // 设置为 false 表示不执行模拟模式，在此模式下，SQL 语句会被打印出来但不会被执行。
		PrepareStmt:                              false, // 设置为 false 表示不使用预编译语句
		DisableAutomaticPing:                     false, // 设置为 false 表示启用自动心跳检测，定期检查数据库连接的有效性。
		DisableForeignKeyConstraintWhenMigrating: false, // 设置为 false 表示在迁移数据库结构时不禁用外键约束
		IgnoreRelationshipsWhenMigrating:         false, // 设置为 false 表示在迁移数据库结构时不忽略模型之间的关系。
		DisableNestedTransaction:                 false, // 设置为 false 表示允许嵌套事务
		AllowGlobalUpdate:                        false, // 设置为 false 表示不允许全局更新操作
		QueryFields:                              false, // 设置为 false 表示在查询时使用默认的行为，不显示所有字段
		CreateBatchSize:                          0,     // 设置为 0 表示不启用批量创建功能
		TranslateError:                           false, // 设置为 false 表示不翻译错误信息
		PropagateUnscoped:                        false, // 设置为 false 表示不传播未作用域的操作
		ClauseBuilders:                           nil,   // 设置为 nil 表示使用默认的子句构建器
		ConnPool:                                 nil,   //	设置为 nil 表示使用默认的连接池配置
		Dialector:                                nil,   // 设置为 nil 表示使用默认的数据库方言
		Plugins:                                  nil,   // Plugins: 设置为 nil 表示不使用任何插件。
	})
	mysqlDb, _ := DB.DB()
	mysqlDb.SetMaxOpenConns(mysqlMaxConns)
	mysqlDb.SetMaxIdleConns(mysqlIdeConns)
}
