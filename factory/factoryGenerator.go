package factory

type employee struct {
	name, position string
	annualIncome   int
}

// The functional approach of factory generators.
func newEmployeeFactory(position string, income int) func(name string) *employee {
	return func(name string) *employee {
		return &employee{name: name, position: position, annualIncome: income}
	}
}

// The structural approach of factory generators.
type employeeFactory struct {
	position     string
	annualIncome int
}

func newEmployeeFactoryStruct(position string, income int) *employeeFactory {
	return &employeeFactory{position: position, annualIncome: income}
}

func (ef *employeeFactory) createEmployee(name string) *employee {
	return &employee{name: name, position: ef.position, annualIncome: ef.annualIncome}
}
