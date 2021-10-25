package main

import "fmt"

func main() {
	var option int
	for option != 9 {
		showOptions()
		option = readOption()

		if option == 1 {
			fmt.Println("list")
		} else if option == 2 {
			fmt.Println("info")
		} else if option == 3 {
			fmt.Println("create")
		} else if option == 4 {
			fmt.Println("update")
		} else if option == 5 {
			fmt.Println("delete")
		} else {
			fmt.Println("Invalid option")
		}
	}
	fmt.Println("Goodbye!")
}

func showOptions() {
	fmt.Println(
		"Welcome! \n\n",
		"1 - List all contacts\n",
		"2 - Show contact info\n",
		"3 - Create new contact\n",
		"4 - Update contact\n",
		"5 - Delete contact\n\n",
		"9 - Exit\n ",
	)
}

func readOption() int {
	var option int
	fmt.Println("Choose an option:")
	fmt.Scan(&option)
	return option
}

func readName() string {
	var name string
	fmt.Println("Name:")
	fmt.Scan(&name)
	return name
}

func readPhone() string {
	var phone string
	fmt.Println("Phone:")
	fmt.Scan(&phone)
	return phone
}

func readContact() Contact {
	name := readName()
	phone := readPhone()
	return Contact{name, phone}
}
