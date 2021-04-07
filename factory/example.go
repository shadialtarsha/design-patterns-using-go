package factory

import "fmt"

func Example() {
	p := newPerson("Shadi", 27)
	fmt.Println(*p)

	d := newDog("Wof Wof!", 2)
	d.bark()
	td := newDog("Wo!", 10)
	td.bark()

	developerFactory := newEmployeeFactory("developer", 80000)
	managerFactory := newEmployeeFactory("manager", 100000)
	developer := developerFactory("shadi")
	manager := managerFactory("bob")
	fmt.Println(*developer)
	fmt.Println(*manager)

	john := NewMainOfficeWorker("John", 100)
	jane := NewAuxOfficeWorker("Jane", 200)
	fmt.Println(john, *john.Office)
	fmt.Println(jane, *jane.Office)
}
