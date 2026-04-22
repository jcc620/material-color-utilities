package temperature

import (
	"math"
	"sort"
	"material-color-utilities/cam"
	"material-color-utilities/quantize"
	"material-color-utilities/utils"
)

type TemperatureCache struct {
	input                   *cam.Hct
	precomputedComplement  *cam.Hct
	precomputedHctsByTemp  []*cam.Hct
	precomputedHctsByHue   []*cam.Hct
	precomputedTempsByHct  map[utils.Argb]float64
}

func NewTemperatureCache(input *cam.Hct) *TemperatureCache {
	return &TemperatureCache{
		input: input,
	}
}

func (c *TemperatureCache) GetComplement() *cam.Hct {
	if c.precomputedComplement != nil {
		return c.precomputedComplement
	}

	coldest := c.getColdest()
	coldestHue := coldest.Hue()
	coldestTemp := c.getTempsByHct()[coldest.ToInt()]

	warmest := c.getWarmest()
	warmestHue := warmest.Hue()
	warmestTemp := c.getTempsByHct()[warmest.ToInt()]
	range_ := warmestTemp - coldestTemp
	startHueIsColdestToWarmest := isBetween(c.input.Hue(), coldestHue, warmestHue)
	startHue := coldestHue
	if startHueIsColdestToWarmest {
		startHue = warmestHue
	}
	endHue := warmestHue
	if startHueIsColdestToWarmest {
		endHue = coldestHue
	}
	directionOfRotation := 1.0
	smallestError := 1000.0
	answer := c.getHctsByHue()[int(math.Round(c.input.Hue()))%361]

	complementRelativeTemp := 1.0 - c.GetRelativeTemperature(c.input)
	for hueAddend := 0.0; hueAddend <= 360.0; hueAddend += 1.0 {
		hue := utils.SanitizeDegreesDouble(startHue + directionOfRotation*hueAddend)
		if !isBetween(hue, startHue, endHue) {
			continue
		}
		possibleAnswer := c.getHctsByHue()[int(math.Round(hue))%361]
		relativeTemp := (c.getTempsByHct()[possibleAnswer.ToInt()] - coldestTemp) / range_
		error := math.Abs(complementRelativeTemp - relativeTemp)
		if error < smallestError {
			smallestError = error
			answer = possibleAnswer
		}
	}
	c.precomputedComplement = answer
	return answer
}

func (c *TemperatureCache) GetAnalogousColors() []*cam.Hct {
	return c.GetAnalogousColorsWithCount(5, 12)
}

func (c *TemperatureCache) GetAnalogousColorsWithCount(count, divisions int) []*cam.Hct {
	startHue := int(math.Round(c.input.Hue()))
	startHct := c.getHctsByHue()[startHue%361]
	lastTemp := c.GetRelativeTemperature(startHct)

	var allColors []*cam.Hct
	allColors = append(allColors, startHct)

	absoluteTotalTempDelta := 0.0
	for i := 0; i < 360; i++ {
		hue := utils.SanitizeDegreesInt(startHue + i)
		hct := c.getHctsByHue()[hue%361]
		temp := c.GetRelativeTemperature(hct)
		tempDelta := math.Abs(temp - lastTemp)
		lastTemp = temp
		absoluteTotalTempDelta += tempDelta
	}

	hueAddend := 1
	tempStep := absoluteTotalTempDelta / float64(divisions)
	totalTempDelta := 0.0
	lastTemp = c.GetRelativeTemperature(startHct)
	for len(allColors) < divisions {
		hue := utils.SanitizeDegreesInt(startHue + hueAddend)
		hct := c.getHctsByHue()[hue%361]
		temp := c.GetRelativeTemperature(hct)
		tempDelta := math.Abs(temp - lastTemp)
		totalTempDelta += tempDelta

		desiredTotalTempDeltaForIndex := float64(len(allColors)) * tempStep
		indexSatisfied := totalTempDelta >= desiredTotalTempDeltaForIndex
		indexAddend := 1
		for indexSatisfied && len(allColors) < divisions {
			allColors = append(allColors, hct)
			desiredTotalTempDeltaForIndex = float64(len(allColors)+indexAddend) * tempStep
			indexSatisfied = totalTempDelta >= desiredTotalTempDeltaForIndex
			indexAddend++
		}
		lastTemp = temp
		hueAddend++

		if hueAddend > 360 {
			for len(allColors) < divisions {
				allColors = append(allColors, hct)
			}
			break
		}
	}

	var answers []*cam.Hct
	answers = append(answers, c.input)

	ccwCount := int(math.Floor((float64(count) - 1.0) / 2.0))
	for i := 1; i < (ccwCount + 1); i++ {
		index := 0 - i
		for index < 0 {
			index = len(allColors) + index
		}
		if index >= len(allColors) {
			index = index % len(allColors)
		}
		answers = append([]*cam.Hct{allColors[index]}, answers...)
	}

	cwCount := count - ccwCount - 1
	for i := 1; i < (cwCount + 1); i++ {
		index := i
		if index >= len(allColors) {
			index = index % len(allColors)
		}
		answers = append(answers, allColors[index])
	}

	return answers
}

func (c *TemperatureCache) GetRelativeTemperature(hct *cam.Hct) float64 {
	warmest := c.getWarmest()
	coldest := c.getColdest()
	range_ := c.getTempsByHct()[warmest.ToInt()] - c.getTempsByHct()[coldest.ToInt()]
	differenceFromColdest := c.getTempsByHct()[hct.ToInt()] - c.getTempsByHct()[coldest.ToInt()]
	if range_ == 0 {
		return 0.5
	}
	return differenceFromColdest / range_
}

func RawTemperature(color *cam.Hct) float64 {
	lab := quantize.LabFromInt(color.ToInt())
	hue := utils.SanitizeDegreesDouble(math.Atan2(lab.B, lab.A) * 180.0 / utils.Pi)
	chroma := math.Hypot(lab.A, lab.B)
	return -0.5 + 0.02*math.Pow(chroma, 1.07)*math.Cos(utils.SanitizeDegreesDouble(hue-50.0)*utils.Pi/180.0)
}

func (c *TemperatureCache) getColdest() *cam.Hct {
	return c.getHctsByTemp()[0]
}

func (c *TemperatureCache) getWarmest() *cam.Hct {
	hcts := c.getHctsByTemp()
	return hcts[len(hcts)-1]
}

func (c *TemperatureCache) getHctsByHue() []*cam.Hct {
	if c.precomputedHctsByHue != nil {
		return c.precomputedHctsByHue
	}
	var hcts []*cam.Hct
	for hue := 0.0; hue <= 360.0; hue += 1.0 {
		hcts = append(hcts, cam.NewHct(hue, c.input.Chroma(), c.input.Tone()))
	}
	c.precomputedHctsByHue = hcts
	return hcts
}

func (c *TemperatureCache) getHctsByTemp() []*cam.Hct {
	if c.precomputedHctsByTemp != nil {
		return c.precomputedHctsByTemp
	}

	hcts := make([]*cam.Hct, len(c.getHctsByHue()))
	copy(hcts, c.getHctsByHue())
	hcts = append(hcts, c.input)
	tempsByHct := c.getTempsByHct()
	sort.Slice(hcts, func(i, j int) bool {
		return tempsByHct[hcts[i].ToInt()] < tempsByHct[hcts[j].ToInt()]
	})
	c.precomputedHctsByTemp = hcts
	return hcts
}

func (c *TemperatureCache) getTempsByHct() map[utils.Argb]float64 {
	if c.precomputedTempsByHct != nil {
		return c.precomputedTempsByHct
	}

	allHcts := make([]*cam.Hct, len(c.getHctsByHue()))
	copy(allHcts, c.getHctsByHue())
	allHcts = append(allHcts, c.input)

	temps := make(map[utils.Argb]float64)
	for _, hct := range allHcts {
		temps[hct.ToInt()] = RawTemperature(hct)
	}

	c.precomputedTempsByHct = temps
	return temps
}

func isBetween(angle, a, b float64) bool {
	if a < b {
		return a <= angle && angle <= b
	}
	return a <= angle || angle <= b
}
