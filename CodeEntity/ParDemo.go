package main

/**
理解     不同代码块中的重名变量
        变量重声明中的变量




1、变量重声明的变量一定是在同一个代码块中。变量重声明可以理解为同一个人可以多次修改自己的名字，但是本质上是同一个人
2、变量重名就是不同的人但是名字是一样的
*/
import (
	"fmt"
)

var block = "package"

func main() {

	block := "fuction"
	{

		block := "inner"

		fmt.Printf("the block is %s.\n", block)
	}
	fmt.Printf("the block is %s.\n", block)

}
