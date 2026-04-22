package palettes

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/utils"
)

type TonalPalette struct {
	hue      float64
	chroma   float64
	keyColor *cam.Hct
}

func NewTonalPaletteFromInt(argb utils.Argb) *TonalPalette {
	c := cam.CamFromInt(argb)
	return &TonalPalette{
		hue:      c.Hue,
		chroma:   c.Chroma,
		keyColor: NewKeyColor(c.Hue, c.Chroma).Create(),
	}
}

func NewTonalPaletteFromHct(hct *cam.Hct) *TonalPalette {
	return &TonalPalette{
		hue:      hct.Hue(),
		chroma:   hct.Chroma(),
		keyColor: hct,
	}
}

func NewTonalPaletteFromHueChroma(hue, chroma float64) *TonalPalette {
	return &TonalPalette{
		hue:      hue,
		chroma:   chroma,
		keyColor: NewKeyColor(hue, chroma).Create(),
	}
}

func NewTonalPaletteFromHueChromaKey(hue, chroma float64, keyColor *cam.Hct) *TonalPalette {
	return &TonalPalette{
		hue:      hue,
		chroma:   chroma,
		keyColor: keyColor,
	}
}

func (p *TonalPalette) Get(tone float64) utils.Argb {
	return cam.SolveToInt(p.hue, p.chroma, tone)
}

func (p *TonalPalette) Hue() float64 {
	return p.hue
}

func (p *TonalPalette) Chroma() float64 {
	return p.chroma
}

func (p *TonalPalette) KeyColor() *cam.Hct {
	return p.keyColor
}

type KeyColor struct {
	hue              float64
	requestedChroma float64
	chromaCache      map[float64]float64
}

func NewKeyColor(hue, requestedChroma float64) *KeyColor {
	return &KeyColor{
		hue:              hue,
		requestedChroma: requestedChroma,
		chromaCache:      make(map[float64]float64),
	}
}

func (k *KeyColor) Create() *cam.Hct {
	const pivotTone = 50
	const toneStepSize = 1
	const epsilon = 0.01

	lowerTone := 0
	upperTone := 100
	for lowerTone < upperTone {
		midTone := (lowerTone + upperTone) / 2
		isAscending := k.maxChroma(float64(midTone)) < k.maxChroma(float64(midTone+toneStepSize))
		sufficientChroma := k.maxChroma(float64(midTone)) >= k.requestedChroma-epsilon

		if sufficientChroma {
			if math.Abs(float64(lowerTone-pivotTone)) < math.Abs(float64(upperTone-pivotTone)) {
				upperTone = midTone
			} else {
				if lowerTone == midTone {
					return cam.NewHct(k.hue, k.requestedChroma, float64(lowerTone))
				}
				lowerTone = midTone
			}
		} else {
			if isAscending {
				lowerTone = midTone + toneStepSize
			} else {
				upperTone = midTone
			}
		}
	}
	return cam.NewHct(k.hue, k.requestedChroma, float64(lowerTone))
}

func (k *KeyColor) maxChroma(tone float64) float64 {
	if v, ok := k.chromaCache[tone]; ok {
		return v
	}
	chroma := cam.NewHct(k.hue, 200.0, tone).Chroma()
	k.chromaCache[tone] = chroma
	return chroma
}
