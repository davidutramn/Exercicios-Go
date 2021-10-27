package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":3000", nil)
}
