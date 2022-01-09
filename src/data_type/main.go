package main

import (
	"fmt"
)

// difference bewteen nil slice with empty slice
func nilSliceDiffEmptySlice() {
	// what is a nil slice? A nil slice is a slice has a length and capacity of zero and has no underlying array
	// nil slice
	var a []string
	fmt.Println("slice a is nil:", a == nil, len(a), cap(a))

	// what is a empty slice? A empty slice is a slice has a length and capacity of zero but has underying array with length zero
	// empty slice
	b := []string{} // or b:= make([]string, 0)
	fmt.Println("slice b is nil:", b == nil, len(b), cap(b))
}

func main() {
	nilSliceDiffEmptySlice()
}
