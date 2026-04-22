package cam

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/utils"
)

type ViewingConditions struct {
	AdaptingLuminance           float64
	BackgroundLstar             float64
	Surround                    float64
	DiscountingIlluminant      bool
	BackgroundYToWhitePointY    float64
	Aw                          float64
	Nbb                         float64
	Ncb                         float64
	C                           float64
	Nc                          float64
	Fl                          float64
	FlRoot                      float64
	Z                           float64
	WhitePoint                  [3]float64
	RgbD                        [3]float64
}

func CreateViewingConditions(whitePoint [3]float64, adaptingLuminance, backgroundLstar, surround float64, discountingIlluminant bool) ViewingConditions {
	backgroundYNormalized := utils.YFromLstar(backgroundLstar) / 100.0
	whitePointY := whitePoint[1]
	backgroundYToWhitePointY := backgroundYNormalized / (whitePointY / 100.0)
	
	rC := 0.401288*whitePoint[0] + 0.650173*whitePoint[1] - 0.051461*whitePoint[2]
	gC := -0.250268*whitePoint[0] + 1.204414*whitePoint[1] + 0.045854*whitePoint[2]
	bC := -0.002079*whitePoint[0] + 0.048952*whitePoint[1] + 0.953127*whitePoint[2]

	d := 1.0
	if !discountingIlluminant {
		d = surround * (1.0 - (1.0/3.6)*math.Exp((-adaptingLuminance-42.0)/92.0))
	}
	if d > 1.0 {
		d = 1.0
	} else if d < 0.0 {
		d = 0.0
	}

	rgbD := [3]float64{
		d*(100.0/rC) + 1.0 - d,
		d*(100.0/gC) + 1.0 - d,
		d*(100.0/bC) + 1.0 - d,
	}
	k := 1.0 / (5.0*adaptingLuminance + 1.0)
	k4 := k * k * k * k
	fl := k4*adaptingLuminance + 0.1*(1.0-k4)*(1.0-k4)*math.Pow(5.0*adaptingLuminance, 1.0/3.0)
	flRoot := math.Pow(fl, 0.25)
	n := backgroundYToWhitePointY
	z := 1.48 + math.Sqrt(n)
	nbb := 0.725 / math.Pow(n, 0.2)
	ncb := nbb
	rgbAFactors := [3]float64{
		math.Pow(fl*rgbD[0]*rC/100.0, 0.42),
		math.Pow(fl*rgbD[1]*gC/100.0, 0.42),
		math.Pow(fl*rgbD[2]*bC/100.0, 0.42),
	}
	rgbA := [3]float64{
		400.0 * rgbAFactors[0] / (rgbAFactors[0] + 27.13),
		400.0 * rgbAFactors[1] / (rgbAFactors[1] + 27.13),
		400.0 * rgbAFactors[2] / (rgbAFactors[2] + 27.13),
	}
	aw := (2.0*rgbA[0] + rgbA[1] + 0.05*rgbA[2]) * nbb
	
	c := surround 
	nc := surround 

	return ViewingConditions{
		AdaptingLuminance:        adaptingLuminance,
		BackgroundLstar:          backgroundLstar,
		Surround:                 surround,
		DiscountingIlluminant:   discountingIlluminant,
		BackgroundYToWhitePointY: backgroundYToWhitePointY,
		Aw:                       aw,
		Nbb:                      nbb,
		Ncb:                      ncb,
		C:                        c,
		Nc:                       nc,
		Fl:                       fl,
		FlRoot:                   flRoot,
		Z:                        z,
		WhitePoint:               whitePoint,
		RgbD:                     rgbD,
	}
}

var DefaultViewingConditions = ViewingConditions{
	AdaptingLuminance:        11.725676537,
	BackgroundLstar:          50.000000000,
	Surround:                 2.000000000,
	DiscountingIlluminant:   false,
	BackgroundYToWhitePointY: 0.184186503,
	Aw:                       29.981000900,
	Nbb:                      1.016919255,
	Ncb:                      1.016919255,
	C:                        0.689999998,
	Nc:                       1.000000000,
	Fl:                       0.388481468,
	FlRoot:                   0.789482653,
	Z:                        1.909169555,
	WhitePoint:               [3]float64{95.047, 100.0, 108.883},
	RgbD:                     [3]float64{1.021177769, 0.986307740, 0.933960497},
}
