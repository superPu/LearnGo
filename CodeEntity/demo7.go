package main

/**
:= 短变量只是一个语法糖
这里记录一个概念：
我们通常把不改变某个程序与外界的任何交互方式和规则，而只是改变其内部实现的代码修改方式，叫做对程序的重构
*/

/**
go语言有一个类型推断，就是说可以在go这个静态语言上玩出动态语言的玩法

变量重声明
*/
import (
	"flag"
)

//
//func main() {
//	//var name string
//	//var name = flag.String("name", "everyone", "the greeting object")
//	var name = getTheFlag()
//	//name := flag.String("name", "everyone", "the greeting object")
//	fmt.Println(*name)
//}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object")
}
