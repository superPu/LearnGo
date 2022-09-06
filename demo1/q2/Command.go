package main

/**
理解go语言的命令源码文件，命令这个词
*/
import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {

	flag.StringVar(&name, "name", "everyone", "the greeting object")

}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	}
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
}
