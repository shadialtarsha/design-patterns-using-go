package main

import (
	"github.com/shadialtarsha/design-patterns-using-go/builder"
	"github.com/shadialtarsha/design-patterns-using-go/factory"
	"github.com/shadialtarsha/design-patterns-using-go/specification"
)

func main() {
	// Specification Pattern
	specification.Example()
	// Builder Pattern
	builder.Example()
	// Factory Pattern
	factory.Example()
}
