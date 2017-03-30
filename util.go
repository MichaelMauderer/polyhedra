package geometry

import "math"

func RadToDeg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}

func DegToRad(rad float64) float64 {
	return rad * math.Pi / 180.0
}
