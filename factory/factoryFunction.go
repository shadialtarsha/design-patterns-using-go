package factory

type person struct {
	name     string
	age      int
	eyeCount int
}

func newPerson(name string, age int) *person {
	return &person{name: name, age: age, eyeCount: 2}
}
