package main

import (
	"fmt"

	"github.com/jaswdr/faker/v2"
)

// ok, so to generate the row, i would take the field.

// type Field struct {
// 	Name      string
// 	FieldType string
// }
//
// type Table struct {
// 	name   string
// 	fields []Field
// }

type Person struct {
	UUID    string
	Name    string
	Address string
	Phone   string
	Email   string
}

func generatePerson() Person {
	f := faker.New()

	return Person{
		UUID:    f.UUID().V4(),
		Name:    f.Person().Name(),
		Address: f.Address().Address(),
		Phone:   f.Phone().Number(),
		Email:   f.Internet().Email(),
	}
}

type Product struct {
	UUID        string
	Name        string
	Description string
}

func generateProduct() Product {
	f := faker.New()

	return Product{
		UUID:        f.UUID().V4(),
		Name:        f.Lorem().Text(50),
		Description: f.Lorem().Text(500),
	}
}

// Later, make sure orders have a product from db (fk)

type Order struct {
	UUID    string
	Address string
	Phone   string
	Email   string
}

func generateOrder() Order {
	f := faker.New()

	return Order{
		UUID:    f.UUID().V4(),
		Address: f.Address().Address(),
		Phone:   f.Phone().Number(),
		Email:   f.Internet().Email(),
	}
}

func main() {
	person := generatePerson()
	product := generateProduct()
	order := generateOrder()
	fmt.Printf("New Person generated: \n %v \n\n", person)
	fmt.Printf("New Product generated: \n %v \n\n", product)
	fmt.Printf("New Order generated: \n %v \n\n", order)
}
