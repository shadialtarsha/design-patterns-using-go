package factory

import "fmt"

type Dog interface {
	bark()
}

type dog struct {
	sound string
	age   int
}

func (d *dog) bark() {
	fmt.Printf("Dog's sound is %s\n", d.sound)
}

type tiredDog struct {
	sound string
	age   int
}

func (td *tiredDog) bark() {
	fmt.Printf("Tired Dog's sound is %s\n", td.sound)
}

func newDog(sound string, age int) Dog {
	if age > 5 {
		return &tiredDog{sound: sound, age: age}
	}
	return &dog{sound: sound, age: age}
}
