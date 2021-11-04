package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(400)
		w.Write([]byte("No username provided"))
		return
	}

	user, err := getUser(username)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "text/html")
	err = userTemplate.Execute(w, user)
	if err != nil {
		log.Panic(err)
	}
}
