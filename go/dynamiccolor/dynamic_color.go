package dynamiccolor

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/contrast"
	"github.com/jcc620/material-color-utilities/go/palettes"
	"github.com/jcc620/material-color-utilities/go/utils"
)

type DynamicColor struct {
	Name             string
	Palette          func(scheme *DynamicScheme) *palettes.TonalPalette
	Tone             func(scheme *DynamicScheme) float64
	IsBackground     bool
	Background       func(scheme *DynamicScheme) *DynamicColor
	SecondBackground func(scheme *DynamicScheme) *DynamicColor
	ContrastCurve    *ContrastCurve
	ToneDeltaPair    func(scheme *DynamicScheme) ToneDeltaPair
}

func NewDynamicColor(
	name string,
	palette func(scheme *DynamicScheme) *palettes.TonalPalette,
	tone func(scheme *DynamicScheme) float64,
	isBackground bool,
	background func(scheme *DynamicScheme) *DynamicColor,
	secondBackground func(scheme *DynamicScheme) *DynamicColor,
	contrastCurve *ContrastCurve,
	toneDeltaPair func(scheme *DynamicScheme) ToneDeltaPair,
) *DynamicColor {
	return &DynamicColor{
		Name:             name,
		Palette:          palette,
		Tone:             tone,
		IsBackground:     isBackground,
		Background:       background,
		SecondBackground: secondBackground,
		ContrastCurve:    contrastCurve,
		ToneDeltaPair:    toneDeltaPair,
	}
}

func FromPalette(
	name string,
	palette func(scheme *DynamicScheme) *palettes.TonalPalette,
	tone func(scheme *DynamicScheme) float64,
) *DynamicColor {
	return &DynamicColor{
		Name:    name,
		Palette: palette,
		Tone:    tone,
	}
}

func (c *DynamicColor) GetArgb(scheme *DynamicScheme) utils.Argb {
	return c.Palette(scheme).Get(c.GetTone(scheme))
}

func (c *DynamicColor) GetHct(scheme *DynamicScheme) *cam.Hct {
	return cam.NewHctFromInt(c.GetArgb(scheme))
}

func (c *DynamicColor) GetTone(scheme *DynamicScheme) float64 {
	decreasingContrast := scheme.ContrastLevel < 0

	if c.ToneDeltaPair != nil {
		toneDeltaPair := c.ToneDeltaPair(scheme)
		roleA := toneDeltaPair.RoleA
		roleB := toneDeltaPair.RoleB
		delta := toneDeltaPair.Delta
		polarity := toneDeltaPair.Polarity
		stayTogether := toneDeltaPair.StayTogether

		bg := c.Background(scheme)
		bgTone := bg.GetTone(scheme)

		aIsNearer := (polarity == TonePolarityNearer ||
			(polarity == TonePolarityLighter && !scheme.IsDark) ||
			(polarity == TonePolarityDarker && scheme.IsDark))
		nearer := roleA
		if !aIsNearer {
			nearer = roleB
		}
		farther := roleB
		if !aIsNearer {
			farther = roleA
		}
		amNearer := c.Name == nearer.Name
		expansionDir := 1.0
		if !scheme.IsDark {
			expansionDir = -1.0
		}

		nContrast := nearer.ContrastCurve.Get(scheme.ContrastLevel)
		fContrast := farther.ContrastCurve.Get(scheme.ContrastLevel)

		nInitialTone := nearer.Tone(scheme)
		nTone := nInitialTone
		if contrast.RatioOfTones(bgTone, nInitialTone) < nContrast {
			nTone = ForegroundTone(bgTone, nContrast)
		}

		fInitialTone := farther.Tone(scheme)
		fTone := fInitialTone
		if contrast.RatioOfTones(bgTone, fInitialTone) < fContrast {
			fTone = ForegroundTone(bgTone, fContrast)
		}

		if decreasingContrast {
			nTone = ForegroundTone(bgTone, nContrast)
			fTone = ForegroundTone(bgTone, fContrast)
		}

		if (fTone-nTone)*expansionDir < delta {
			fTone = nTone + delta*expansionDir
			if fTone < 0 {
				fTone = 0
			} else if fTone > 100 {
				fTone = 100
			}

			if (fTone-nTone)*expansionDir < delta {
				nTone = fTone - delta*expansionDir
				if nTone < 0 {
					nTone = 0
				} else if nTone > 100 {
					nTone = 100
				}
			}
		}

		if nTone >= 50 && nTone < 60 {
			if expansionDir > 0 {
				nTone = 60
				fTone = math.Max(fTone, nTone+delta*expansionDir)
			} else {
				nTone = 49
				fTone = math.Min(fTone, nTone+delta*expansionDir)
			}
		} else if fTone >= 50 && fTone < 60 {
			if stayTogether {
				if expansionDir > 0 {
					nTone = 60
					fTone = math.Max(fTone, nTone+delta*expansionDir)
				} else {
					nTone = 49
					fTone = math.Min(fTone, nTone+delta*expansionDir)
				}
			} else {
				if expansionDir > 0 {
					fTone = 60
				} else {
					fTone = 49
				}
			}
		}

		if amNearer {
			return nTone
		}
		return fTone
	}

	answer := c.Tone(scheme)

	if c.Background == nil {
		return answer
	}

	bgTone := c.Background(scheme).GetTone(scheme)
	desiredRatio := c.ContrastCurve.Get(scheme.ContrastLevel)

	if contrast.RatioOfTones(bgTone, answer) < desiredRatio {
		answer = ForegroundTone(bgTone, desiredRatio)
	}

	if decreasingContrast {
		answer = ForegroundTone(bgTone, desiredRatio)
	}

	if c.IsBackground && answer >= 50 && answer < 60 {
		if contrast.RatioOfTones(49, bgTone) >= desiredRatio {
			answer = 49
		} else {
			answer = 60
		}
	}

	if c.SecondBackground != nil {
		bgTone1 := c.Background(scheme).GetTone(scheme)
		bgTone2 := c.SecondBackground(scheme).GetTone(scheme)

		upper := math.Max(bgTone1, bgTone2)
		lower := math.Min(bgTone1, bgTone2)

		if contrast.RatioOfTones(upper, answer) >= desiredRatio &&
			contrast.RatioOfTones(lower, answer) >= desiredRatio {
			return answer
		}

		lightOption := contrast.Lighter(upper, desiredRatio)
		darkOption := contrast.Darker(lower, desiredRatio)

		var availables []float64
		if lightOption != -1 {
			availables = append(availables, lightOption)
		}
		if darkOption != -1 {
			availables = append(availables, darkOption)
		}

		prefersLight := TonePrefersLightForeground(bgTone1) || TonePrefersLightForeground(bgTone2)
		if prefersLight {
			if lightOption < 0 {
				return 100
			}
			return lightOption
		}
		if len(availables) == 1 {
			return availables[0]
		}
		if darkOption < 0 {
			return 0
		}
		return darkOption
	}

	return answer
}

func ForegroundTone(bgTone, ratio float64) float64 {
	lighterTone := contrast.LighterUnsafe(bgTone, ratio)
	darkerTone := contrast.DarkerUnsafe(bgTone, ratio)
	lighterRatio := contrast.RatioOfTones(lighterTone, bgTone)
	darkerRatio := contrast.RatioOfTones(darkerTone, bgTone)
	preferLighter := TonePrefersLightForeground(bgTone)

	if preferLighter {
		negligibleDifference := math.Abs(lighterRatio-darkerRatio) < 0.1 && lighterRatio < ratio && darkerRatio < ratio
		if lighterRatio >= ratio || lighterRatio >= darkerRatio || negligibleDifference {
			return lighterTone
		}
		return darkerTone
	} else {
		if darkerRatio >= ratio || darkerRatio >= lighterRatio {
			return darkerTone
		}
		return lighterTone
	}
}

func EnableLightForeground(tone float64) float64 {
	if TonePrefersLightForeground(tone) && !ToneAllowsLightForeground(tone) {
		return 49.0
	}
	return tone
}

func TonePrefersLightForeground(tone float64) bool {
	return math.Round(tone) < 60
}

func ToneAllowsLightForeground(tone float64) bool {
	return math.Round(tone) <= 49
}
