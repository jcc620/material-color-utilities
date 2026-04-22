package score

import (
	"math"
	"sort"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/utils"
)

const (
	targetChroma            = 48.0
	weightProportion        = 0.7
	weightChromaAbove      = 0.3
	weightChromaBelow      = 0.1
	cutoffChroma            = 5.0
	cutoffExcitedProportion = 0.01
)

type ScoreOptions struct {
	Desired           int
	FallbackColorArgb utils.Argb
	Filter            bool
}

func DefaultScoreOptions() ScoreOptions {
	return ScoreOptions{
		Desired:           4,
		FallbackColorArgb: 0xff4285f4,
		Filter:            true,
	}
}

type scoredHCT struct {
	hct   *cam.Hct
	score float64
}

func RankedSuggestions(argbToPopulation map[utils.Argb]uint32, options ScoreOptions) []utils.Argb {
	var colorsHct []*cam.Hct
	huePopulation := make([]uint32, 360)
	var populationSum float64
	for argb, population := range argbToPopulation {
		hct := cam.NewHctFromInt(argb)
		colorsHct = append(colorsHct, hct)
		hue := int(math.Floor(hct.Hue()))
		huePopulation[utils.SanitizeDegreesInt(hue)] += population
		populationSum += float64(population)
	}

	hueExcitedProportions := make([]float64, 360)
	for hue := 0; hue < 360; hue++ {
		proportion := float64(huePopulation[hue]) / populationSum
		for i := hue - 14; i < hue+16; i++ {
			neighborHue := utils.SanitizeDegreesInt(i)
			hueExcitedProportions[neighborHue] += proportion
		}
	}

	var scoredHcts []scoredHCT
	for _, hct := range colorsHct {
		hue := utils.SanitizeDegreesInt(int(math.Round(hct.Hue())))
		proportion := hueExcitedProportions[hue]
		if options.Filter && (hct.Chroma() < cutoffChroma || proportion <= cutoffExcitedProportion) {
			continue
		}

		proportionScore := proportion * 100.0 * weightProportion
		chromaWeight := weightChromaBelow
		if hct.Chroma() >= targetChroma {
			chromaWeight = weightChromaAbove
		}
		chromaScore := (hct.Chroma() - targetChroma) * chromaWeight
		score := proportionScore + chromaScore
		scoredHcts = append(scoredHcts, scoredHCT{hct, score})
	}

	sort.Slice(scoredHcts, func(i, j int) bool {
		return scoredHcts[i].score > scoredHcts[j].score
	})

	var chosenColors []*cam.Hct
	for differenceDegrees := 90; differenceDegrees >= 15; differenceDegrees-- {
		chosenColors = nil
		for _, entry := range scoredHcts {
			hct := entry.hct
			duplicateHue := false
			for _, chosenHct := range chosenColors {
				if utils.DiffDegrees(hct.Hue(), chosenHct.Hue()) < float64(differenceDegrees) {
					duplicateHue = true
					break
				}
			}
			if !duplicateHue {
				chosenColors = append(chosenColors, hct)
				if len(chosenColors) >= options.Desired {
					break
				}
			}
		}
		if len(chosenColors) >= options.Desired {
			break
		}
	}

	var colors []utils.Argb
	if len(chosenColors) == 0 {
		colors = append(colors, options.FallbackColorArgb)
	} else {
		for _, chosenHct := range chosenColors {
			colors = append(colors, chosenHct.ToInt())
		}
	}
	return colors
}
