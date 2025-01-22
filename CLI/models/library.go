package models

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Library []Book

type Book struct {
	Author         string
	Title          string
	CompleteStatus bool
	CreatedAt      time.Time
	CompletedAt    *time.Time
}

func (library *Library) Add(author, title string) {

	book := Book{
		Author:         author,
		Title:          title,
		CompleteStatus: false,
		CreatedAt:      time.Now(),
		CompletedAt:    nil,
	}
	*library = append(*library, book)

}

func (library *Library) ValidateIndex(index int) error {
	if index < 0 || index >= len(*library) {
		log.Fatal("Invalid index")
		return errors.New("Invalid index")
	}
	return nil
}

func (library *Library) Remove(index int) error {
	l := *library
	if err := library.ValidateIndex(index); err != nil {
		return err
	}

	*library = append(l[:index], l[index+1:]...)

	return nil
}

func (library *Library) Edit(index int, title, author string) error {

	if err := library.ValidateIndex(index); err != nil {
		return err
	}
	l := *library

	l[index].Author = author
	l[index].Title = title

	return nil

}

func (library *Library) Complete(index int) error {
	if err := library.ValidateIndex(index); err != nil {
		return err
	}

	l := *library

	if l[index].CompleteStatus {
		l[index].CompleteStatus = false
		l[index].CompletedAt = nil
		return nil
	}

	l[index].CompleteStatus = true
	CompleteTime := time.Now()
	l[index].CompletedAt = &CompleteTime
	return nil
}

func (library *Library) Print() {
	
	t := table.New(os.Stdout)
	t.SetAlignment(table.AlignLeft)
	t.SetHeaders("ID", "Title", "Author", "Completed", "Created At", "Completed At")

	for i, book := range *library {
		completed := "❌"
		completedAt := ""

		if book.CompleteStatus {
			completed = "✅"
			if book.CompletedAt != nil {
				completedAt = book.CompletedAt.Format(time.RFC1123)
			}
		}

		t.AddRow(strconv.Itoa(i), book.Title, book.Author, completed, book.CreatedAt.Format(time.RFC1123), completedAt)
	}

	t.Render()

}
