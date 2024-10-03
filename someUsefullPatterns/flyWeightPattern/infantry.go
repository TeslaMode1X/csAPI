package flyWeightPattern

import "fmt"

type Infantry struct{}

func (i Infantry) Report(positionX, positionY int) {
	fmt.Printf("Infantry at position (%d, %d)\n", positionX, positionY)
}
