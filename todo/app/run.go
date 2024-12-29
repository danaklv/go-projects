package app

import (
	"fmt"
	"net/http"
	"todo/database"
	"todo/handle"
)

func Run() {
	database.ConnectToDb()
	handle.Handlers()
	http.ListenAndServe(":8080", nil)
	fmt.Println("http://localhost:8080/")

}
