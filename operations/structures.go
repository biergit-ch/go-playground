package operations

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Addresses []Address
}

type Address struct {
	Street 	string
	City  	string
	Plz     int
	Country string
}
