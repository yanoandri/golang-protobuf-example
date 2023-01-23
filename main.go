package main

import (
	"fmt"
	"log"

	bk "github.com/yanoandri/golang_protobuf/class"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	library := &bk.Library{
		Id:      1,
		Name:    "Nearby Library",
		Address: "Ohio",
		Books: []*bk.Book{
			{
				Id:          1,
				Title:       "How to be great developers",
				Description: "New book",
				IsBorrow:    false,
			},
			{
				Id:          2,
				Title:       "Principal Engineering edition 1",
				Description: "Principal engineering books",
				IsBorrow:    true,
				BorrowDate:  timestamppb.Now(),
			},
			{
				Id:          3,
				Title:       "Principal Engineering edition 2nd edition",
				Description: "Principal engineering books",
				IsBorrow:    false,
			},
		},
	}

	booksAvail := len(library.GetBooks())
	fmt.Println("Welcome to library : " + library.GetName())
	fmt.Printf("We have %v books available to borrow\n", booksAvail)
	for i := 0; i < booksAvail; i++ {
		book := library.GetBooks()[i]
		fmt.Printf("%s-%s\n", book.Title, book.Description)
	}

	person := &bk.Person{
		Id:   1,
		Name: "Suranto",
		Age:  30,
	}

	fmt.Printf("%v has entered the %v\n", person.GetName(), library.GetName())
	person.AttendDate = timestamppb.Now()
	library.Persons = append(library.Persons, person)
	fmt.Printf("%v has fill attendance tracker at %v\n", person.GetName(), person.GetAttendDate())

	person.Book = library.GetBooks()[0]
	fmt.Printf("%v will borrow %v books\n", person.GetName(), person.GetBook().GetTitle())

	for i := 0; i < booksAvail; i++ {
		book := library.GetBooks()[i]
		if person.Book.GetId() == book.GetId() {
			book.IsBorrow = true
			book.BorrowDate = timestamppb.Now()
			fmt.Printf("%s marked as borrow\n", book.GetTitle())
		}
	}

	libraryData, _ := proto.Marshal(library)
	fmt.Println("\n proto format : ")
	fmt.Println(libraryData)

	newLibrary := &bk.Library{}
	err := proto.Unmarshal(libraryData, newLibrary)

	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}

	fmt.Println("\n new library data : ")
	fmt.Println(newLibrary)
}
