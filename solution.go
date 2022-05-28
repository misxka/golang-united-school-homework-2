package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type Amount int

func CalcSquare(sideLen float64, sidesNum Amount) float64 {
	const (
		sidesTriangle = 3
		sidesSquare   = 4
		sidesCircle   = 0
	)

	switch sidesNum {
	case sidesSquare:
		return math.Pow(sideLen, 2)
	case sidesTriangle:
		return math.Sqrt(3) * math.Pow(sideLen, 2) / 4
	case sidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}
}
