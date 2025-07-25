package util

// 0 -> 0.5, 90 -> 0.75, 179 -> 0.99.., 180 -> 0.0, 270 -> 0.25
func GetPercentage(degrees int) float64 {
	switch {
	case degrees == 180:
		return 1
	case degrees == -180:
		return 0
	}
	return float64((degrees+180)%360) / 360
}
