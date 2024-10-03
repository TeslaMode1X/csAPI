package decoratorPattern

type Coffee interface {
	Cost() float64
	Description() string
}

// decorator Pattern
//mocha := decoratorPattern.MochaCoffee{}
//moshi := decoratorPattern.MoshiCoffee{}
//
//milkMocha := decoratorPattern.MilkDecorator{CoffeeDecorator: decoratorPattern.CoffeeDecorator{Coffee: mocha}}
//chocolateMocha := decoratorPattern.ChocolateDecorator{CoffeeDecorator: decoratorPattern.CoffeeDecorator{Coffee: mocha}}
//
//milkMoshi := decoratorPattern.MilkDecorator{CoffeeDecorator: decoratorPattern.CoffeeDecorator{Coffee: moshi}}
//chocolateMoshi := decoratorPattern.ChocolateDecorator{CoffeeDecorator: decoratorPattern.CoffeeDecorator{Coffee: moshi}}
//
//fmt.Println(milkMocha.Description())
//fmt.Println(milkMoshi.Description())
//
//fmt.Println(chocolateMocha.Description())
//fmt.Println(chocolateMoshi.Description())
