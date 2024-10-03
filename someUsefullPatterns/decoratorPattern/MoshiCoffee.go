package decoratorPattern

type MoshiCoffee struct{}

func (m MoshiCoffee) Cost() float64 {
	return 10.00
}

func (m MoshiCoffee) Description() string {
	return "Moshi Coffee"
}
