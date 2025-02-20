package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type yourTypeNameHere int

const (
	SidesTriangle yourTypeNameHere = 3
	SidesSquare   yourTypeNameHere = 4
	SidesCircle   yourTypeNameHere = 0
)

func CalcSquare(sideLen float64, sidesNum yourTypeNameHere) float64 {

	if sidesNum == SidesTriangle {

		return (sideLen * sideLen * math.Sqrt(3)) / 4

	} else if sidesNum == SidesSquare {

		return (sideLen * sideLen)

	} else if sidesNum == SidesCircle {

		return math.Pi * (sideLen * sideLen)

	} else {

		return 0

	}
}

func Main() {

	fmt.Print("Площадь треугольника 10", CalcSquare(10, 3))
	fmt.Print("Площадь квадрата 10", CalcSquare(10, 4))
	fmt.Print("Площадь круга 10", CalcSquare(10, 0))
}
