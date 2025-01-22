package main

import (
	"cli/models"
)

func main() {

	var lib models.Library

	lib.Add("Lev Tolstoy", "Anna Karenina")
	lib.Complete(0)
	lib.Print()

}
