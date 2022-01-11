package common

import (
	"fmt"
	"sync"
	"testing"
)

func display(a interface{}) {
	fmt.Println(a)
}

func TestRange(t *testing.T) {
	lang := []string{
		"Go", "Java", "C++", "Python", "PHP", "Javascript",
	}

	var wg sync.WaitGroup
	for _, l := range lang {
		wg.Add(1)
		go func() {
			// incorrect!!!
			// in a Go for loop, the loop variable is reused for each iteration
			display(l)
			wg.Done()
		}()
	}
	wg.Wait()
}
