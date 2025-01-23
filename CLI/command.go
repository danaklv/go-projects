package main

import (
	"cli/models"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add      string
	Remove   int
	Edit     string
	Complete int
	List     bool
}

func NewCmdFlags() *CmdFlags {
	f := CmdFlags{}

	flag.StringVar(&f.Add, "add", "",  "Add a new book. Use author:title")
	flag.IntVar(&f.Remove, "remove", -1, "Delete a book by index. Use -delete index")
	flag.StringVar(&f.Edit, "edit", "", "Edit book. Use -edit index:author:title")
	flag.IntVar(&f.Complete, "complete", -1, "Change complete status. Use -complete index")
	flag.BoolVar(&f.List, "list", false, "Show all books")

	flag.Parse()

	return &f

}

func (f *CmdFlags) Execute(library *models.Library) {
	switch {
	case f.List:
		library.Print()
	case f.Add != "":
	parts := strings.SplitN(f.Add, ":", 2)
	if len(parts) != 2 {
		fmt.Println("Invalid format. Use author:title")
		os.Exit(1)
	}
	library.Add(parts[0],parts[1])
	case f.Edit != "":
		parts := strings.SplitN(f.Add, ":", 3)
		if len(parts) != 3 {
			fmt.Println("Invalid format. Use id:author:title")
			
		}
		indx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}
		library.Edit(indx,parts[1], parts[2])
	case f.Complete != -1:
		library.Complete(f.Complete)
	case f.Remove != -1:
		library.Remove(f.Remove)
	default:
		fmt.Println("Invalid command. Use -help for command list")
		os.Exit(1)



	}
}
