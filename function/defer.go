package main

import "fmt"

/*
	defer应用场景
		1. 关闭文件流；
		2. 解锁一个加锁的资源；
		3. 打印最终日志；
		4. 关闭数据库连接；
 */
func loopDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 执行该函数，返回的是几？
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f())
}