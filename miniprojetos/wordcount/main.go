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

func wordcount(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordlist := make(map[string]int)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordlist[strings.ToLower(scanner.Text())]++
	}

	for word, count := range wordlist {
		fmt.Printf("%q: %d\n", word, count)
	}
}
