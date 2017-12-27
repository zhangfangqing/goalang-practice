package main

import (
	"fmt"
)

func main() {
	fmt.Println(noReturn())
}
func noReturn() (result int) {
	defer func() {
		p := recover()
		result = p.(int)
	}()
	panic(1)
}
