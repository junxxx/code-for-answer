package main

import (
	"fmt"

	"github.com/junxxx/greetings"
)

func main() {
	message := greetings.Hello("Glad")
	fmt.Println(message)
}
