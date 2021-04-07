package factory

const (
	sport = iota
	suv
)

type car struct {
	model, color string
}

func newCar(model int) *car {
	switch model {
	case 0:
		return &car{model: "sport", color: ""}
	case 1:
		return &car{model: "suv", color: ""}
	default:
		panic("unsupported model")
	}
}
