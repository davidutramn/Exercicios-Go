package main

import "fmt"

type Person struct {
	Name string // export the field if it begins with a capital letter
	Age  int
}

func main() {
	var p1 Person // start with the zero values of each type
	fmt.Println(p1)

	p1 = Person{"Davi", 22}
	fmt.Println(p1)

	p1 = Person{Name: "Davi"}
	fmt.Println(p1)

	p1 = Person{Name: "Davi", Age: 21}
	fmt.Println(p1)

	getOlder(&p1.Age)
	fmt.Println(p1)

	pp := new(Person) // structs are commonly dealt with throught pointers
	fmt.Println(pp)

	pp = &Person{Name: "Davi", Age: 22}
	fmt.Println(pp)

	fmt.Println(p1 == *pp)

	p1.sayHello()
}

func getOlder(age *int) {
	*(age)++
}

func (p Person) sayHello() {
	fmt.Println("Hello, my name is", p.Name)
}
