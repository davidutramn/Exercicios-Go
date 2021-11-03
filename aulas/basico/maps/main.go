package main

import (
	"fmt"
	"sort"
)

func main() {
	// how to create ages map
	var ages map[string]int
	ages = make(map[string]int)
	ages = map[string]int{}
	ages = map[string]int{"Davi": 22}

	// var a map[string]int
	// a["r"] = 1 // assignment to entry in nil map

	ages["Lucas"] = 1
	ages["João"] += 1 // or
	ages["Pedro"]++   // or

	fmt.Println(ages)

	if zeroValue, ok := ages["Joãoo"]; !ok {
		fmt.Println("Joãoo is not a key", zeroValue)
	}

	delete(ages, "Pedro") // safe operation

	// no order
	for key, value := range ages {
		fmt.Println(key, value)
	}

	names := make([]string, 0, len(ages))
	for k := range ages {
		names = append(names, k)
	}
	sort.Strings(names) // ordering
	for _, name := range names {
		fmt.Println(name, ages[name])
	}

	a := map[string]int{}
	// fmt.Println(ages == a) // error: map can only be compared to nil
	fmt.Println(equal(ages, a))

}

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}
