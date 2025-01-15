package session

import "net/http"

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	session.Save(r, w)
}
