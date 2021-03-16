package specification

type Color string
type Size int

type Product struct {
	name  string
	color Color
	size  Size
}

// Specification pattern implementation.
// Perfect implementation for Open-Close principle.
// Spec type is open for extension but closed for modification.
type Spec interface {
	isSatisfied(p Product) bool
}

func Filter(products []Product, s Spec) []Product {
	filteredProdcuts := []Product{}
	for _, v := range products {
		if s.isSatisfied(v) {
			filteredProdcuts = append(filteredProdcuts, v)
		}
	}
	return filteredProdcuts
}

type ColorFilter struct {
	color Color
}

func (cf *ColorFilter) isSatisfied(p Product) bool {
	return p.color == cf.color
}

type SizeFilter struct {
	size Size
}

func (sf *SizeFilter) isSatisfied(p Product) bool {
	return p.size == sf.size
}

// This is an example of the composite design pattern.
type ColorAndSizeFilter struct {
	colorFilter ColorFilter
	sizeFilter  SizeFilter
}

func (csf *ColorAndSizeFilter) isSatisfied(p Product) bool {
	return csf.colorFilter.isSatisfied(p) && csf.sizeFilter.isSatisfied(p)
}
