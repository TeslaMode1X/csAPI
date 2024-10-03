package decoratorPattern

type CoffeeDecorator struct {
	Coffee Coffee
}

func (d CoffeeDecorator) Cost() float64 {
	return d.Coffee.Cost()
}

func (d CoffeeDecorator) Description() string {
	return d.Coffee.Description()
}
