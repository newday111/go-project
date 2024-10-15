package example

import (
	"fmt"
	"testing"
)

func TestGoCatchException(t *testing.T) {
	result := twoAdd()
	fmt.Println(result)
	fmt.Println("执行完毕")
}

func add(a, b int) int {
	if b == 0 {
		panic("add by zero")
	}
	return a + b
}

func twoAdd() int {
	count := 1
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("runtime panic: %v\n", err)
			count = 0
		}
	}()
	//	可以执行的
	fmt.Println("继续执行了")
	add(1, 2)
	//	执行不到了
	fmt.Println("继续执行了")
	return count
}
