package main

import "cli/models"

func main() {

	var lib models.Library

	storage := NewStorage[models.Library]("library.json")
	storage.Load(&lib)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&lib)

	storage.Save(lib)

}
