package compositePattern

import "fmt"

type Leaf struct {
	name  string
	price int
}

func (l Leaf) showPrice() {
	fmt.Printf("%s: %d tenge\n", l.name, l.price)
}
