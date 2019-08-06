package main

import "fmt"

func main() {
	var i interface{}
	i = 10
	var j interface{}
	j = 11
	fmt.Println(i < j)
}
