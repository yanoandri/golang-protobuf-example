package main

import (
	"fmt"
	"log"

	bk "github.com/yanoandri/golang_protobuf/class"

	"github.com/golang/protobuf/proto"
)

func main() {

	fmt.Println("Test")

	book := &bk.Book{
		Id:          1,
		Title:       "How to be great developers",
		Description: "",
	}

	data, err := proto.Marshal(book)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	newBook := &bk.Book{}
	err = proto.Unmarshal(data, newBook)

	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Println(newBook)
}
