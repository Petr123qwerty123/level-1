package tasks

import (
	"fmt"
	"wb-level-1/point"
)

func Task24() {
	p1 := point.NewPoint(1, 2)
	p2 := point.NewPoint(-2, 3)

	d := p1.DistanceTo(p2)

	fmt.Printf("Расстояние: %.2f", d)
}
