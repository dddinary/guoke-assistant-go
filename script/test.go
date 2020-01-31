package main

import "fmt"

type Book struct {
	Price	int32
	Name	string
}

func main() {
	var book Book
	fmt.Println(book)
}
