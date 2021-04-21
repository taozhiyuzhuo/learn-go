package main //包,表明代码所在模块
import (
	"fmt"
	"os"
)

// 功能实现
func main() {
	fmt.Println(os.Args) // 命令行参数
	fmt.Println("hello world")
	os.Exit(-1)
}

// 1.必须是main包: package main
// 2. 必须是main方法
// 3. 文件名不一定是main.go

// 退出返回值
// go中main函数不支持返回值
// 通过os.Exit来返回状态
