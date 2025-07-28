package main

import (
	"fmt"
	"learn-go-with-tests/hello"
	"learn-go-with-tests/iteration"
)

func main() {
	fmt.Println(hello.Hello("tony", ""))
	fmt.Println(iteration.Iteration("b", 5))
}
