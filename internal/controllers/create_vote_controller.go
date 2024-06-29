package controllers

import "net/http"

func CreateVoteController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.RemoteAddr))
}
