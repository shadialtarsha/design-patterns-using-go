package specification

import "fmt"

func Example() {
	apple := Product{name: "Apple", color: "green", size: 1}
	house := Product{name: "House", color: "green", size: 3}
	shoes := Product{name: "Shoes", color: "red", size: 3}

	products := []Product{apple, house, shoes}

	fmt.Println("Color Filter")
	greenColorFilter := ColorFilter{color: "green"}
	fmt.Println(Filter(products, &greenColorFilter)) // [{name: "Apple", color: "green", size: 1}, {name: "House", color: "green", size: 3}]

	fmt.Println("Size Filter")
	bigSizeFilter := SizeFilter{size: 3}
	fmt.Println(Filter(products, &bigSizeFilter)) // [{name: "House", color: "green", size: 3}, {name: "Shoes", color: "red", size: 3}]

	fmt.Println("Size And Color Filter")
	bigSizeAndGreenColorFilter := ColorAndSizeFilter{colorFilter: greenColorFilter, sizeFilter: bigSizeFilter}
	fmt.Println(Filter(products, &bigSizeAndGreenColorFilter)) // {name: "House", color: "green", size: 3}
}
