package main

import (
	"fmt"
	"hash/crc32"
	"log"

	"github.com/junxxx/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Duo", "Anna", "Peter"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Amazon Redshift"), crc32q))
}
