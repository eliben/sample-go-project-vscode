package main

import "fmt"

func printer(a string) {
	b := a + " and b"
	fmt.Println(b)
}

func main() {
	p := "hello"
	printer(p)
}
