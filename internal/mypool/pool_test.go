package mypool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("first get item from pool:", p)

	p.Name = "pName"

	pool.Put(p)

	fmt.Println("first get ", pool.Get().(*Person))
	fmt.Println("second get ", pool.Get().(*Person))
}
