package polyhedra

import "math"

func RadToDeg(rad float64) float64 {
	return rad * 180.0 / math.Phi
}

func DegToRad(rad float64) float64 {
	return rad * math.Phi / 180.0
}
