package blend

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/utils"
)

func Harmonize(designColor, keyColor utils.Argb) utils.Argb {
	fromHct := cam.NewHctFromInt(designColor)
	toHct := cam.NewHctFromInt(keyColor)
	differenceDegrees := utils.DiffDegrees(fromHct.Hue(), toHct.Hue())
	rotationDegrees := math.Min(differenceDegrees*0.5, 15.0)
	outputHue := utils.SanitizeDegreesDouble(
		fromHct.Hue() +
			rotationDegrees*utils.RotationDirection(fromHct.Hue(), toHct.Hue()))
	fromHct.SetHue(outputHue)
	return fromHct.ToInt()
}

func HctHue(from, to utils.Argb, amount float64) utils.Argb {
	ucs := Cam16Ucs(from, to, amount)
	ucsHct := cam.NewHctFromInt(ucs)
	fromHct := cam.NewHctFromInt(from)
	fromHct.SetHue(ucsHct.Hue())
	return fromHct.ToInt()
}

func Cam16Ucs(from, to utils.Argb, amount float64) utils.Argb {
	fromCam := cam.CamFromInt(from)
	toCam := cam.CamFromInt(to)

	aj := fromCam.Jstar
	aa := fromCam.Astar
	ab := fromCam.Bstar

	bj := toCam.Jstar
	ba := toCam.Astar
	bb := toCam.Bstar

	jstar := aj + (bj-aj)*amount
	astar := aa + (ba-aa)*amount
	bstar := ab + (bb-ab)*amount

	blended := cam.CamFromUcsAndViewingConditions(jstar, astar, bstar, cam.DefaultViewingConditions)
	return cam.IntFromCam(blended)
}
