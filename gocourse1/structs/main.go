package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	// also possible to just use the bare struct name
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p) // %+v : display field names and values in a struct
}

func (p *person) updateName(newFirstName string) {
	// also possible to do (*p).firstName but this is not necessary
	p.firstName = newFirstName
}

func main() {
	// alex := person{"Alex", "Anderson"} // not recommended
	// joe := person{firstName: "Joe", lastName: "Bloggs"} // better struct declaration

	var amy person // create new struct with zero values

	amy.firstName = "Amy"
	amy.lastName = "Jones"
	amy.contactInfo.email = "amy@awesomecoder.com"
	amy.contactInfo.zipCode = 94101

	amy.updateName("Amelia")
	amy.print()

	// struct with embedded structs -- declare values together
	// don't forget those trailing comments!
	letitia := person{
		firstName: "Letitia",
		lastName:  "Baker",
		contactInfo: contactInfo{
			zipCode: 94101,
			email:   "titia@awesomecoder.com",
		},
	}

	letitia.print()
}
