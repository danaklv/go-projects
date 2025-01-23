package controllers

// func SendResetCodeController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		tmpl.ExecuteTemplate(w, "resetPassword", nil)
// 	} else if r.Method == http.MethodPost {
// 		email := r.FormValue("email")
// 		err := services.RequestPasswordReset(email)

// 		data := models.ResetPassword{
// 			Email: email,
// 		}
// 		if err != nil {
// 			data.ErrorCode = err
// 		} else {
// 			data.SuccessCode = "Code sent successfully "
// 		}

// 		tmpl.ExecuteTemplate(w, "resetPassword", data)

// 	}

// }

// func ResetPasswordController(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		code := r.FormValue("code")
// 		newPassword := r.FormValue("password")
// 		email := r.FormValue("email")

// 		err := services.ResetPassword(code, newPassword, email)
// 		tmpl.ExecuteTemplate(w, "resetPassword", map[string]string{"Message": err})

// 	}
// }
