package quantize

import (
	"math"
	"material-color-utilities/utils"
)

type Lab struct {
	L, A, B float64
}

func (l Lab) DeltaE(other Lab) float64 {
	dl := l.L - other.L
	da := l.A - other.A
	db := l.B - other.B
	return dl*dl + da*da + db*db
}

var whitePointD65 = [3]float64{95.047, 100.0, 108.883}

func IntFromLab(lab Lab) utils.Argb {
	e := 216.0 / 24389.0
	kappa := 24389.0 / 27.0
	ke := 8.0

	fy := (lab.L + 16.0) / 116.0
	fx := (lab.A / 500.0) + fy
	fz := fy - (lab.B / 200.0)
	fx3 := fx * fx * fx
	xNormalized := (116.0*fx - 16.0) / kappa
	if fx3 > e {
		xNormalized = fx3
	}
	yNormalized := lab.L / kappa
	if lab.L > ke {
		yNormalized = fy * fy * fy
	}
	fz3 := fz * fz * fz
	zNormalized := (116.0*fz - 16.0) / kappa
	if fz3 > e {
		zNormalized = fz3
	}
	x := xNormalized * whitePointD65[0]
	y := yNormalized * whitePointD65[1]
	z := zNormalized * whitePointD65[2]

	rL := 3.2406*x - 1.5372*y - 0.4986*z
	gL := -0.9689*x + 1.8758*y + 0.0415*z
	bL := 0.0557*x - 0.2040*y + 1.0570*z

	red := utils.Delinearized(rL)
	green := utils.Delinearized(gL)
	blue := utils.Delinearized(bL)

	return utils.ArgbFromRgb(red, green, blue)
}

func LabFromInt(argb utils.Argb) Lab {
	red := utils.RedFromInt(argb)
	green := utils.GreenFromInt(argb)
	blue := utils.BlueFromInt(argb)
	redL := utils.Linearized(red)
	greenL := utils.Linearized(green)
	blueL := utils.Linearized(blue)
	x := 0.41233895*redL + 0.35762064*greenL + 0.18051042*blueL
	y := 0.2126*redL + 0.7152*greenL + 0.0722*blueL
	z := 0.01932141*redL + 0.11916382*greenL + 0.95034478*blueL

	yNormalized := y / whitePointD65[1]
	e := 216.0 / 24389.0
	kappa := 24389.0 / 27.0
	var fy float64
	if yNormalized > e {
		fy = math.Pow(yNormalized, 1.0/3.0)
	} else {
		fy = (kappa*yNormalized + 16) / 116
	}

	xNormalized := x / whitePointD65[0]
	var fx float64
	if xNormalized > e {
		fx = math.Pow(xNormalized, 1.0/3.0)
	} else {
		fx = (kappa*xNormalized + 16) / 116
	}

	zNormalized := z / whitePointD65[2]
	var fz float64
	if zNormalized > e {
		fz = math.Pow(zNormalized, 1.0/3.0)
	} else {
		fz = (kappa*zNormalized + 16) / 116
	}

	l := 116.0*fy - 16
	a := 500.0 * (fx - fy)
	b := 200.0 * (fy - fz)
	return Lab{L: l, A: a, B: b}
}
