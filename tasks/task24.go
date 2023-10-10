package tasks

import (
	"fmt"
	"wb-level-1/point"
)

func Task24() {
	// создаем 2 точки
	p1 := point.NewPoint(1, 2)
	p2 := point.NewPoint(-2, 3)

	// получаем расстояние между точками
	d := p1.DistanceTo(p2)

	fmt.Printf("Расстояние: %.2f", d)
}
