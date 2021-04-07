package main

import (
	"fmt"

	"github.com/shadialtarsha/design-patterns-using-go/builder"
	"github.com/shadialtarsha/design-patterns-using-go/factory"
	"github.com/shadialtarsha/design-patterns-using-go/prototype"
	"github.com/shadialtarsha/design-patterns-using-go/specification"
)

func main() {
	// Specification Pattern
	fmt.Println("========== Specification ==========")
	specification.Example()
	// Builder Pattern
	fmt.Println("========== Builder ==========")
	builder.Example()
	// Factory Pattern
	fmt.Println("========== Factory ==========")
	factory.Example()
	// Prototype Pattern
	fmt.Println("========== Prototype ==========")
	prototype.Example()
}
