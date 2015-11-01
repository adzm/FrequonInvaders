package fourier

import (
	"github.com/ArchRobison/Gophetica/nimble"
)

const (
	clutSize   = 128            // Size of Clut along either axis.  Power of 2 to speed up indexing.
	clutCenter = clutSize / 2   // Clut indices corresponding to (0,0)
	clutRadius = clutCenter - 1 // Distance from center representing magnitude of 1.
)

func clutCoor(k int) (z float32) {
	const (
		clutScale  = 1.0 / clutRadius
		clutOffset = -clutCenter * clutScale
	)
	return float32(k)*clutScale + clutOffset
}

var clut [clutSize][clutSize]nimble.Pixel

type colorMap interface {
	Color(x, y float32) (r, g, b float32)
}

var currentMap colorMap

func setColoring(cm colorMap) {
	if cm == currentMap {
		return
	}
	currentMap = cm

	for i := 0; i < clutSize; i++ {
		y := clutCoor(i)
		for j := 1; j < clutSize; j++ {
			x := clutCoor(j)
			clut[i][j] = nimble.RGB(cm.Color(x, y))
		}
	}
}