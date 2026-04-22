package dynamiccolor

import "github.com/jcc620/material-color-utilities/go/utils"

type ContrastCurve struct {
	Low, Normal, Medium, High float64
}

func (c ContrastCurve) Get(contrastLevel float64) float64 {
	if contrastLevel <= -1.0 {
		return c.Low
	} else if contrastLevel < 0.0 {
		return utils.Lerp(c.Low, c.Normal, (contrastLevel - (-1.0)) / 1.0)
	} else if contrastLevel < 0.5 {
		return utils.Lerp(c.Normal, c.Medium, (contrastLevel - 0.0) / 0.5)
	} else if contrastLevel < 1.0 {
		return utils.Lerp(c.Medium, c.High, (contrastLevel - 0.5) / 0.5)
	} else {
		return c.High
	}
}
