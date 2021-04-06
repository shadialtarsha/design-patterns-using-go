package builder

type employee struct {
	name, position string
}

type employeeMod func(*employee)

type employeeBuilder struct {
	actions []employeeMod
}

func (b *employeeBuilder) called(name string) *employeeBuilder {
	b.actions = append(b.actions, func(e *employee) {
		e.name = name
	})
	return b
}

func (b *employeeBuilder) worksAs(position string) *employeeBuilder {
	b.actions = append(b.actions, func(e *employee) {
		e.position = position
	})
	return b
}

func (b *employeeBuilder) build() *employee {
	e := employee{}
	for _, a := range b.actions {
		a(&e)
	}
	return &e
}

// extend employeeBuilder would as simple as adding another method that add a new action.
