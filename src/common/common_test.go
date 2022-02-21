package common

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func display(a interface{}) {
	fmt.Println(a)
}

func TestSlice(t *testing.T) {
	s := "Hello World"
	as := s[2:] + s[:2]
	fmt.Printf("%T\n", as)
	fmt.Println("type of as is:", reflect.TypeOf(as))
}

func TestConvert(t *testing.T) {
	// replace := "%20"
	s := "Hello World"
	fmt.Println([]byte(s))
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			ans = append(ans, s[i])
		} else {
			ans = append(ans, '%')
			ans = append(ans, '2')
			ans = append(ans, '0')
		}
	}
	fmt.Println(string(ans))
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
