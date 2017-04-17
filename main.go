package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	defer func() {
		//var r interface{}
		//r = recover()
		if r := recover(); r != nil {
			fmt.Println(r, "recover")
		}
	}()

	s, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	s.Write([]byte("inhai"))

	b := make([]byte, 15)

	s.Read(b)

	fmt.Println(string(b))

}
