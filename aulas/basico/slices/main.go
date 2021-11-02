package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)

	m := months[1:4]               // shares the same underlying array as months
	fmt.Println(m, len(m), cap(m)) // len == 3, cap == cap(months)

	m = m[:]       // slice can use the length or m or the capacity of m (capacity of months)
	fmt.Println(m) // same values as before

	m = m[:cap(m)]
	fmt.Println(m) // uses the value of the underlying array

	fmt.Println("Before", months)
	reverse(months[:len(months)/2]) // changes only the elements passed
	fmt.Println("After", months)
	fmt.Println(m) // also changes all the references

	// fmt.Println(months == m) // error, slices can only be compared to nil
	fmt.Println(equal(months, m)) // manual comparison unless it is a []byte (bytes.Equal)
}

func reverse(s []string) {
	// s is an alias for the underlying array
	// the changes in s will reflect in the original slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
