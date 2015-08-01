package szudzik

import (
	"math"
)

// As described in http://szudzik.com/ElegantPairing.pdf

func ElegantPair(x, y float64) (z float64) {
	if math.Max(x, y) != x {
		z = math.Pow(y, 2) + x
	} else {
		z = math.Pow(x, 2) + x + y
	}
	return
}

func ElegantUnpair(z float64) (x, y float64) {
	fz := math.Floor(math.Sqrt(z))
	c := z - math.Pow(fz, 2)
	if c < fz {
		x, y = c, fz
	} else {
		x, y = fz, c-fz
	}
	return
}
