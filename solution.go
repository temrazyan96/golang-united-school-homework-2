package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare intCustomType = 4
	SidesCircle intCustomType = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var result float64 
	switch(sidesNum) {
		case SidesTriangle:
			result = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
		case SidesSquare:
			result = math.Pow(sideLen, 2)
		case SidesCircle:
			result = math.Pi * math.Pow(sideLen, 2)
	}

	return result
}
