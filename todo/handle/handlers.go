package handle

import "net/http"

func Handlers() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("front/styles"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("front/img"))))

	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	
	
}