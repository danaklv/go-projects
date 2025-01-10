package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("I_DONT_KNOW"))

func CreateNewSession(w http.ResponseWriter, r *http.Request, email string) {

	session, _ := Store.Get(r, "session-name")

	session.Values["email"] = email
	session.Values["authenticated"] = true

	session.Save(r, w)

}

