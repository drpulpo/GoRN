package model

type Order struct {
	Id         int
	Firstname  string
	Lastname   string
	Phone      string
	City       string
	PostalCode string
	State      string
	Country    string
	Crayon00   string
	Crayon01   string
	Crayon02   string
	Crayon03   string
	Crayon04   string
	Crayon05   string
	Crayon06   string
	Crayon07   string
	Crayon08   string
	Crayon09   string
	Crayon10   string
	Crayon11   string
	Crayon12   string
	Crayon13   string
	Crayon14   string
	Crayon15   string
	Crayon16   string
	Crayon17   string
	Crayon18   string
	Crayon19   string
}

// type Order struct {
// 	User         User
// 	OrderDetails []OrderDetil
// }

// type OrderDetil struct {
// 	Quantity int
// 	Product  Product
// }

// type Product struct {
// 	Name        string
// 	Description string
// 	Price       string
// }

// User struct contains user data
type User struct {
	UserName  string
	Password  string
	Firstname string
	Lastname  string
	Phone     string
	Address   Address
}

// Address struct contains address of a given user
type Address struct {
	City    string
	State   string
	Country string
}
