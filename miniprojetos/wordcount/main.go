package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	wordcount("file.txt")
}

func wordcount(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordlist := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, word := range strings.Fields(scanner.Text()) {
			wordlist[word]++
		}
	}

	for word, count := range wordlist {
		fmt.Printf("%q: %d\n", word, count)
	}
}
