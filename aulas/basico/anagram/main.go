package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ovo", "ovo"))
	fmt.Println(isAnagram("ovo", "voo"))
	fmt.Println(isAnagram("lol", "loa"))
	fmt.Println(isAnagram("fried", "fired"))
	fmt.Println(isAnagram("listen", "silent"))
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	c1 := make(map[rune]int)
	c2 := make(map[rune]int)

	for _, v := range s1 {
		c1[v]++
	}

	for _, v := range s2 {
		c2[v]++
	}

	for k, v := range c1 {
		if v != c2[k] {
			return false
		}
	}

	return true
}
