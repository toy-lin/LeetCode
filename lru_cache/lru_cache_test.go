package lru_cache

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}

func Test2(t *testing.T) {
	c := Constructor(2)
	fmt.Println(c.Get(2))

	c.Put(2, 6)
	fmt.Println(c.Get(1))

	c.Put(1, 5)
	c.Put(1, 2)

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}

func Test3(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
