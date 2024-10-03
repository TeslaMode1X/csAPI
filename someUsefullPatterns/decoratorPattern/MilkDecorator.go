package decoratorPattern

type MilkDecorator struct {
	CoffeeDecorator
}

func (milk MilkDecorator) Description() string {
	return milk.CoffeeDecorator.Description() + " with Milk"
}

func (milk MilkDecorator) Cost() float64 {
	return milk.CoffeeDecorator.Cost() + 1.00
}
