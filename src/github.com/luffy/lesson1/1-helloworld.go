package main

import "fmt"
import "time"

// 入口函数，编译成二进制的时候执行的入口函数就是这个main函数
func main()  {
	// 使用go build lesson1就可以进行编译了，go build后面接的是你那个源码package的位置比如lesson1就是你那个包名的地址
	// 当采用第二种目录组织方式的时候，就是嵌套层级的情况下应该带上嵌套的层级目录go build github.com/luffy/lesson1 以src为顶
	// 指定编译后文件生成的位置go build -o bin/hello github.com/luffy/lesson1
	fmt.Println("Hello World")
	// 调用Sleep方法进行sleep的操作，5s，在使用前需要引入对应的包
	time.Sleep(time.Second*5)
}