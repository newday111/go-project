package example

import (
	"fmt"
	"goPro4/utils"
	"testing"
)

func TestSendEmail(t *testing.T) {
	//	直接执行这个的话  下面这个方法会报错配置文件读取不到, 需要自己手动添加或者想起他办法
	err := utils.SendEmail("1750930322@qq.com", "test", "hello")
	fmt.Println(err)
}
