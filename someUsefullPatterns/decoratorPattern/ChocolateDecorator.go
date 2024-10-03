package decoratorPattern

type ChocolateDecorator struct {
	CoffeeDecorator
}

func (choc ChocolateDecorator) Description() string {
	return choc.CoffeeDecorator.Description() + " with Chocolate"
}

func (choc ChocolateDecorator) Cost() float64 {
	return choc.CoffeeDecorator.Cost() + 2.00
}
