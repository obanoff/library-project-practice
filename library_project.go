package main

import (
	"fmt"
	"time"
)

const (
	Available  = false
	CheckedOut = true
)

type Availability bool

type Book struct {
	name        string
	status      Availability
	lended_at   time.Time
	returned_at time.Time
}

type Member struct {
	name  string
	books []*Book
}

type Library struct {
	books   []*Book
	members []*Member
}

func (lib *Library) checkIn(title string) {
	var book *Book
	for _, b := range lib.books {
		if b.name == title {
			book = b
			break
		}
	}

	if book == nil {
		fmt.Println("not allowed to return the book doesn't belong to the library")
		return
	}

	if !book.status {
		fmt.Println("not possible to return the book that's already located in the library")
		return
	}

	var member *Member
	var index int
	for _, m := range lib.members {
		for i, b := range m.books {
			if b.name == title {
				member = m
				index = i
				break
			}
		}
		if member != nil {
			break
		}
	}

	book.status = Available
	book.returned_at = time.Now()

	member.books[index] = member.books[len(member.books)-1]
	member.books = member.books[:len(member.books)-1]
}

func (lib *Library) checkOut(title string, member_name string) {
	var book *Book
	for _, b := range lib.books {
		if b.name == title {
			book = b
			break
		}
	}

	if book == nil {
		fmt.Println("there is no such the book in the library")
		return
	}

	if book.status {
		fmt.Println("book already checked out at", book.lended_at.Format(time.RFC822))
		return
	}

	var member *Member
	for _, m := range lib.members {
		if m.name == member_name {
			member = m
			break
		}
	}

	if member == nil {
		fmt.Println("member doesn't exist")
		return
	}

	book.status = CheckedOut
	book.lended_at = time.Now()

	member.books = append(member.books, book)
}

func (lib *Library) info() {

	printMembers := func() {
		fmt.Println("Members:")
		for _, member := range lib.members {
			books := []string{}
			for _, book := range member.books {
				books = append(books, book.name)
			}
			fmt.Println("Name:", member.name, "checked out:", books)
		}
		fmt.Println()
	}

	printBooks := func(title string, _map map[string]time.Time) {
		fmt.Println(title, "books:")
		for key, value := range _map {
			f := fmt.Sprintf("Title: %v, %v at: %v", key, title, value.Format(time.RFC822))
			fmt.Println(f)
		}
		fmt.Println()
	}

	checkedOutBooks := make(map[string]time.Time)
	availableBooks := make(map[string]time.Time)

	for _, book := range lib.books {
		if book.status {
			checkedOutBooks[book.name] = book.lended_at
		} else {
			availableBooks[book.name] = book.returned_at
		}
	}

	printMembers()
	printBooks("Checked out", checkedOutBooks)
	printBooks("Returned", availableBooks)
}

func main() {
	// books

	book1 := Book{
		name:        "Book1",
		status:      Available,
		lended_at:   time.Time{},
		returned_at: time.Time{},
	}

	book2 := Book{
		name:        "Book2",
		status:      Available,
		lended_at:   time.Time{},
		returned_at: time.Date(2021, 5, 25, 12, 31, 17, 150, time.Local),
	}

	book3 := Book{
		name:        "Book3",
		status:      CheckedOut,
		lended_at:   time.Date(2022, 8, 7, 14, 11, 23, 100, time.Local),
		returned_at: time.Time{},
	}

	book4 := Book{
		name:        "Book4",
		status:      CheckedOut,
		lended_at:   time.Date(2022, 12, 1, 10, 10, 47, 00, time.Local),
		returned_at: time.Time{},
	}

	// members

	member1 := Member{
		name:  "member1",
		books: []*Book{&book3},
	}

	member2 := Member{
		name: "member2",
	}

	member3 := Member{
		name:  "member3",
		books: []*Book{&book4},
	}

	// library

	library := Library{
		books:   []*Book{&book1, &book2, &book3, &book4},
		members: []*Member{&member1, &member2, &member3},
	}

	library.info()

	library.checkOut("Book1", "member1")

	library.info()

	library.checkIn("Book4")

	library.info()

}
