package app

import (
	"fmt"
	"net/http"
	"todo/handle"
	"todo/repositories"
)

func Run() {
	repositories.ConnectToDb()
	handle.Handlers()
	http.ListenAndServe(":8080", nil)
	fmt.Println("http://localhost:8080/")

}
