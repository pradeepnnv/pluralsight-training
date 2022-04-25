package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "Harry Potter and the Sorcerere's Stone",
		Author:        "J.K.Rowling",
		YearPublished: 2003,
	},
	{
		ID:            3,
		Title:         "I,Robot",
		Author:        "Isaac Asimov",
		YearPublished: 1950,
	},
	{
		ID:            4,
		Title:         "The Shining",
		Author:        "Stephen King",
		YearPublished: 1977,
	},
	{
		ID:            5,
		Title:         "The Age of Miracles",
		Author:        "Karen Thompson Walker",
		YearPublished: 2012,
	},
	{
		ID:            6,
		Title:         "Life of Pi",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            7,
		Title:         "It",
		Author:        "Stephen King",
		YearPublished: 1986,
	},
	{
		ID:            8,
		Title:         "Surely You're Joking, Mr. Feynman! Adventures of a Curious Character",
		Author:        "Richard P. Feynman",
		YearPublished: 1985,
	},
	{
		ID:            9,
		Title:         "The Lord of the Rings",
		Author:        "J.R.R. Tolkein",
		YearPublished: 1955,
	},
	{
		ID:            10,
		Title:         "And Then There Were None",
		Author:        "Agatha Christe",
		YearPublished: 1939,
	},
}
