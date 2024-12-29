package handle

import "net/http"

func Handlers() {
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	
}