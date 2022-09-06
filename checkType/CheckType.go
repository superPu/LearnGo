package main

import (
	"fmt"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "0", 1: "1", 2: "2"}
	value, ok := interface{}(container).([]string)
	fmt.Printf("value=%s,ok=%t", value, ok)
	fmt.Printf("the element is %q.\n", container[1])

}
