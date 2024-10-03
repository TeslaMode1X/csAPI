package flyWeightPattern

import "fmt"

type Sniper struct{}

func (s Sniper) Report(positionX, positionY int) {
	fmt.Printf("Sniper at position (%d, %d)\n", positionX, positionY)
}
