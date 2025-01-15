package handle

import (
	"net/http"
	"todo/internal/controllers"
)

func Handlers() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("front/static/"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))

	http.HandleFunc("/", controllers.HomeController)

	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/register", controllers.RegisterController)
	http.HandleFunc("/logout", controllers.LogoutController)


	http.HandleFunc("/create", controllers.CreateController)
	http.HandleFunc("/playlist/", controllers.PlaylistController)

	http.HandleFunc("/sendResetCode", controllers.SendResetCodeController)
	
	http.HandleFunc("/resetPassword", controllers.ResetPasswordController)


}
