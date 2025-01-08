package handle

import "net/http"

func Handlers() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("front/styles"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("front/img"))))

	http.HandleFunc("/", IndexPageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	http.HandleFunc("/forgotPassword", PasswordHandler)
	http.HandleFunc("/main", PersonalPageHandler)
	http.HandleFunc("/create", CreatePlaylistHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/home", HomePageHandler)

}
