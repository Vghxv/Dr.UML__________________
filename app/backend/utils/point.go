package utils

import (
	"fmt"
	"math"

	"Dr.uml/backend/utils/duerror"
)

type Point struct {
	X int
	Y int
}

// FromString Accepts only strings of format "<X>, <Y>" as input.
func FromString(str string) (Point, duerror.DUError) {
	var x, y int
	_, err := fmt.Sscanf(str, "%d, %d", &x, &y)
	if err != nil {
		return Point{}, duerror.NewInvalidArgumentError("Invalid point format: " + str)
	}
	return Point{X: x, Y: y}, nil
}

func (p Point) Magnitude() (float64, duerror.DUError) {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y)), nil
}

func (p Point) MagnitudeInt() (int, duerror.DUError) {
	mag, _ := p.Magnitude()
	return int(mag), nil
}

func EqualPoints(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func AddPoints(p1, p2 Point) Point {
	return Point{X: p1.X + p2.X, Y: p1.Y + p2.Y}
}

func SubPoints(p1, p2 Point) Point {
	return Point{X: p1.X - p2.X, Y: p1.Y - p2.Y}
}
