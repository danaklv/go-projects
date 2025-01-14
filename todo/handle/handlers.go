package handle

import (
	"net/http"
	"todo/internal/controllers"
)

func Handlers() {
	// http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("front/static/styles/"))))
	// http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("front/static/img/"))))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("front/static/"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))

	http.HandleFunc("/", controllers.HomeController)
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/register", controllers.RegisterController)
	http.HandleFunc("/forgotPassword", controllers.PasswordController)

	http.HandleFunc("/create", controllers.CreateController)
	http.HandleFunc("/logout", controllers.LogoutController)

	http.HandleFunc("/playlist/", controllers.PlaylistController)


}
