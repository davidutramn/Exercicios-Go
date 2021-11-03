package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string `json:"title"` // the field must begin with a capital letter
	Year  int    `json:"year"`
}

func main() {
	movies := []Movie{
		{"Whatever", 2020},
		{"Whatever 2", 2021},
	}
	fmt.Println(movies)

	data, err := json.Marshal(movies) // transform into a json format
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	var m2 []Movie
	err = json.Unmarshal(data, &m2) // ignore the fields not present in the type of m2
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m2)
}
