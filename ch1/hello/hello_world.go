package main //包

import (
	"fmt"
	"os"
) //引入代码依赖

func main() {
	//1.不支持任何返回值
	//2.通过OS返回状态
	//3.在程序中直接通过os.Args获取命令行参数
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
	os.Exit(1)
}
