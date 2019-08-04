package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "()[]{}"
	var a interface{}
	for i, c := range s {
		fmt.Println(i, c, reflect.TypeOf(c))
		a = c
	}
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case int32:
		fmt.Println("int32")
	default:
		fmt.Println("default")
	}
}
