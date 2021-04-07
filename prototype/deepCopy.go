package prototype

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetName, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// Copying through serialization.
func (p *Person) deepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	result := Person{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&result)

	return &result
}
