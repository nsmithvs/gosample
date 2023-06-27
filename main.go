package main

import "fmt"

func main() {
	fmt.Println(HelloWorld())
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	var j = 2/0
	
	return "hello world"
}
