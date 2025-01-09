package controllers

import (
	"net/http"
	"todo/session"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	session.DeleteSession(w, r)

}
