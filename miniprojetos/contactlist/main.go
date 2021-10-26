package main

import "fmt"

func main() {
	var option int
	contactList := make(ContactList)
	for option != 9 {
		showOptions()
		option = readOption()

		if option == 1 {
			// List all contacts
			fmt.Println("Contacts:\n ")
			contacts := contactList.listAll()
			if len(contacts) == 0 {
				fmt.Println("No contacts saved")
				continue
			}
			for _, v := range contacts {
				fmt.Println(v)
			}
		} else if option == 2 {
			// Show contact info
			name := readName()
			contact, err := contactList.contactInfo(name)
			if err != nil {
				fmt.Println("Contact not found")
				continue
			}
			fmt.Println("Name:", contact.name)
			fmt.Println("Phone:", contact.phone)
		} else if option == 3 {
			// Create new contact
			contact := readContact()
			if err := contactList.createContact(contact); err != nil {
				fmt.Println("Error:", err)
			}
		} else if option == 4 {
			// Update contact
			name := readName()
			if _, err := contactList.contactInfo(name); err != nil {
				fmt.Println("Contact not found")
				continue
			}
			phone := readPhone()
			if err := contactList.updateContact(name, phone); err != nil {
				fmt.Println("Error:", err)
			}
		} else if option == 5 {
			// Delete contact
			name := readName()
			if err := contactList.deleteContact(name); err != nil {
				fmt.Println("Error:", err)
			}
		} else if option == 9 {
			// Exit the program
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func showOptions() {
	fmt.Println(
		"-=-=-=-=-=-=-=-=-=-=-=-=-=-=\n",
		"Welcome! \n\n",
		"1 - List all contacts\n",
		"2 - Show contact info\n",
		"3 - Create new contact\n",
		"4 - Update contact\n",
		"5 - Delete contact\n\n",
		"9 - Exit\n",
		"\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=",
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
