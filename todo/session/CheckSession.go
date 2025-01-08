package session

import (
	"net/http"
)

func CheckSession(r *http.Request) (string, bool) {
	session, _ := store.Get(r, "session-name")

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		return session.Values["email"].(string), true
	}

	return "", false

}
