package compositePattern

import "fmt"

type Composite struct {
	name       string
	price      int
	components []Component
}

func (c *Composite) addComponent(component ...Component) {
	c.components = append(c.components, component...)
}

func (c *Composite) showPrice() {
	fmt.Println(c.name)
	for _, component := range c.components {
		component.showPrice()
	}
}
