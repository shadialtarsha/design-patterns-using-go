package prototype

import "fmt"

func Example() {
	shadi := Person{
		Name: "Shadi",
		Address: &Address{
			StreetName: "123 berlin",
			City:       "Berlin",
			Country:    "DE",
		},
		Friends: []string{"Bob", "James"},
	}

	michael := shadi.deepCopy()
	michael.Name = "michael"
	michael.Address.Country = "IR"
	michael.Friends[1] = "Carl"

	fmt.Println(shadi, *shadi.Address)
	fmt.Println(*michael, *michael.Address)
}
