package point

import "math"

type Point struct {
	x, y float64
}

// NewPoint - конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// DistanceTo возвращает расстояние до другой точки
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y

	distance := math.Sqrt(dx*dx + dy*dy)

	return distance
}
