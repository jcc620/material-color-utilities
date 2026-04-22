package contrast

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/utils"
)

const (
	contrastRatioEpsilon        = 0.04
	luminanceGamutMapTolerance = 0.4
)

func ratioOfYs(y1, y2 float64) float64 {
	lighter := y1
	if y2 > y1 {
		lighter = y2
	}
	darker := y1
	if lighter == y1 {
		darker = y2
	}
	return (lighter + 5.0) / (darker + 5.0)
}

func RatioOfTones(toneA, toneB float64) float64 {
	if toneA < 0 {
		toneA = 0
	} else if toneA > 100 {
		toneA = 100
	}
	if toneB < 0 {
		toneB = 0
	} else if toneB > 100 {
		toneB = 100
	}
	return ratioOfYs(utils.YFromLstar(toneA), utils.YFromLstar(toneB))
}

func Lighter(tone, ratio float64) float64 {
	if tone < 0.0 || tone > 100.0 {
		return -1.0
	}

	darkY := utils.YFromLstar(tone)
	lightY := ratio*(darkY+5.0) - 5.0
	realContrast := ratioOfYs(lightY, darkY)
	delta := math.Abs(realContrast - ratio)
	if realContrast < ratio && delta > contrastRatioEpsilon {
		return -1
	}

	value := utils.LstarFromY(lightY) + luminanceGamutMapTolerance
	if value < 0 || value > 100 {
		return -1
	}
	return value
}

func Darker(tone, ratio float64) float64 {
	if tone < 0.0 || tone > 100.0 {
		return -1.0
	}

	lightY := utils.YFromLstar(tone)
	darkY := ((lightY + 5.0) / ratio) - 5.0
	realContrast := ratioOfYs(lightY, darkY)

	delta := math.Abs(realContrast - ratio)
	if realContrast < ratio && delta > contrastRatioEpsilon {
		return -1
	}

	value := utils.LstarFromY(darkY) - luminanceGamutMapTolerance
	if value < 0 || value > 100 {
		return -1
	}
	return value
}

func LighterUnsafe(tone, ratio float64) float64 {
	lighterSafe := Lighter(tone, ratio)
	if lighterSafe < 0.0 {
		return 100.0
	}
	return lighterSafe
}

func DarkerUnsafe(tone, ratio float64) float64 {
	darkerSafe := Darker(tone, ratio)
	if darkerSafe < 0.0 {
		return 0.0
	}
	return darkerSafe
}
