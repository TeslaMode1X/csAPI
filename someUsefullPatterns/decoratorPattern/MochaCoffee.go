package decoratorPattern

type MochaCoffee struct{}

func (m MochaCoffee) Cost() float64 {
	return 5.00
}

func (m MochaCoffee) Description() string {
	return "Mocha Coffee"
}
