package string_mul

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(add("120", "24"))
}

func TestMul(t *testing.T) {
	fmt.Println(multiply("12", "12"))
	fmt.Println(multiply("9133", "0"))
	fmt.Println(multiply("0", "9133"))
}
