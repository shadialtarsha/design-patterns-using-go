package builder

type person struct {
	streetAddress, postcode, city string
	companyName, position         string
	annualIncome                  int
}

type personBuilder struct {
	person *person
}

type personAddressBuilder struct {
	personBuilder
}

type personJobBuilder struct {
	personBuilder
}

func (pb *personBuilder) lives() *personAddressBuilder {
	return &personAddressBuilder{*pb}
}

func (pab *personAddressBuilder) at(street string) *personAddressBuilder {
	pab.person.streetAddress = street
	return pab
}

func (pab *personAddressBuilder) in(city string) *personAddressBuilder {
	pab.person.city = city
	return pab
}
func (pab *personAddressBuilder) withPostcode(postcode string) *personAddressBuilder {
	pab.person.postcode = postcode
	return pab
}

func (pb *personBuilder) works() *personJobBuilder {
	return &personJobBuilder{*pb}
}

func (pjb *personJobBuilder) at(company string) *personJobBuilder {
	pjb.person.companyName = company
	return pjb
}

func (pjb *personJobBuilder) as(position string) *personJobBuilder {
	pjb.person.position = position
	return pjb
}
func (pjb *personJobBuilder) earning(income int) *personJobBuilder {
	pjb.person.annualIncome = income
	return pjb
}

func newPersonBuilder() *personBuilder {
	return &personBuilder{&person{}}
}

func (pb *personBuilder) build() *person {
	return pb.person
}
