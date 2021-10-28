package main

import (
	"fmt"
	"net/http"
)

func handleUserRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "User list")
	case "POST":
		fmt.Fprint(w, "User created")
	case "PUT":
		fmt.Fprint(w, "User modified")
	case "PATCH":
		fmt.Fprint(w, "User updated")
	case "DELETE":
		fmt.Fprint(w, "User deleted")
	default:
		fmt.Fprint(w, "METHOD NOT ALLOWED")
	}
}

func handleAdminRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Admin list")
	case "POST":
		fmt.Fprint(w, "Admin created")
	case "PUT":
		fmt.Fprint(w, "Admin modified")
	case "PATCH":
		fmt.Fprint(w, "Admin updated")
	case "DELETE":
		fmt.Fprint(w, "Admin deleted")
	default:
		fmt.Fprint(w, "METHOD NOT ALLOWED")
	}
}

func main() {
	http.HandleFunc("/user", handleUserRoute)
	http.HandleFunc("/admin", handleAdminRoute)
	// http.ListenAndServe(":3000", nil)

	server := http.Server{
		Addr:    ":3000",
		Handler: nil,
	}
	server.ListenAndServe()
}
