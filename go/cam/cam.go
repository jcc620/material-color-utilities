package cam

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/utils"
)

type Cam struct {
	Hue    float64
	Chroma float64
	J      float64
	Q      float64
	M      float64
	S      float64
	Jstar  float64
	Astar  float64
	Bstar  float64
}

func CamFromInt(argb utils.Argb) Cam {
	return CamFromIntAndViewingConditions(argb, DefaultViewingConditions)
}

func CamFromIntAndViewingConditions(argb utils.Argb, vc ViewingConditions) Cam {
	red := utils.RedFromInt(argb)
	green := utils.GreenFromInt(argb)
	blue := utils.BlueFromInt(argb)
	redL := utils.Linearized(red)
	greenL := utils.Linearized(green)
	blueL := utils.Linearized(blue)
	x := 0.41233895*redL + 0.35762064*greenL + 0.18051042*blueL
	y := 0.2126*redL + 0.7152*greenL + 0.0722*blueL
	z := 0.01932141*redL + 0.11916382*greenL + 0.95034478*blueL

	rC := 0.401288*x + 0.650173*y - 0.051461*z
	gC := -0.250268*x + 1.204414*y + 0.045854*z
	bC := -0.002079*x + 0.048952*y + 0.953127*z

	rD := vc.RgbD[0] * rC
	gD := vc.RgbD[1] * gC
	bD := vc.RgbD[2] * bC

	rAF := math.Pow(vc.Fl*math.Abs(rD)/100.0, 0.42)
	gAF := math.Pow(vc.Fl*math.Abs(gD)/100.0, 0.42)
	bAF := math.Pow(vc.Fl*math.Abs(bD)/100.0, 0.42)
	rA := utils.Signum(rD) * 400.0 * rAF / (rAF + 27.13)
	gA := utils.Signum(gD) * 400.0 * gAF / (gAF + 27.13)
	bA := utils.Signum(bD) * 400.0 * bAF / (bAF + 27.13)

	a := (11.0*rA + -12.0*gA + bA) / 11.0
	b := (rA + gA - 2.0*bA) / 9.0
	u := (20.0*rA + 20.0*gA + 21.0*bA) / 20.0
	p2 := (40.0*rA + 20.0*gA + bA) / 20.0

	radians := math.Atan2(b, a)
	degrees := radians * 180.0 / utils.Pi
	hue := utils.SanitizeDegreesDouble(degrees)
	hueRadians := hue * utils.Pi / 180.0
	ac := p2 * vc.Nbb

	j := 100.0 * math.Pow(ac/vc.Aw, vc.C*vc.Z)
	q := (4.0 / vc.C) * math.Sqrt(j/100.0) * (vc.Aw + 4.0) * vc.FlRoot
	huePrime := hue
	if hue < 20.14 {
		huePrime += 360
	}
	eHue := 0.25 * (math.Cos(huePrime*utils.Pi/180.0+2.0) + 3.8)
	p1 := 50000.0 / 13.0 * eHue * vc.Nc * vc.Ncb
	t := p1 * math.Sqrt(a*a+b*b) / (u + 0.305)
	alpha := math.Pow(t, 0.9) * math.Pow(1.64-math.Pow(0.29, vc.BackgroundYToWhitePointY), 0.73)
	c := alpha * math.Sqrt(j/100.0)
	m := c * vc.FlRoot
	s := 50.0 * math.Sqrt((alpha*vc.C)/(vc.Aw+4.0))
	jstar := (1.0 + 100.0*0.007) * j / (1.0 + 0.007*j)
	mstar := 1.0 / 0.0228 * math.Log(1.0+0.0228*m)
	astar := mstar * math.Cos(hueRadians)
	bstar := mstar * math.Sin(hueRadians)

	return Cam{
		Hue:    hue,
		Chroma: c,
		J:      j,
		Q:      q,
		M:      m,
		S:      s,
		Jstar:  jstar,
		Astar:  astar,
		Bstar:  bstar,
	}
}

func IntFromCam(cam Cam) utils.Argb {
	return IntFromCamAndViewingConditions(cam, DefaultViewingConditions)
}

func IntFromCamAndViewingConditions(cam Cam, vc ViewingConditions) utils.Argb {
	alpha := 0.0
	if cam.Chroma != 0.0 && cam.J != 0.0 {
		alpha = cam.Chroma / math.Sqrt(cam.J/100.0)
	}
	t := math.Pow(alpha/math.Pow(1.64-math.Pow(0.29, vc.BackgroundYToWhitePointY), 0.73), 1.0/0.9)
	hRad := cam.Hue * utils.Pi / 180.0
	eHue := 0.25 * (math.Cos(hRad+2.0) + 3.8)
	ac := vc.Aw * math.Pow(cam.J/100.0, 1.0/vc.C/vc.Z)
	p1 := eHue * (50000.0 / 13.0) * vc.Nc * vc.Ncb
	p2 := ac / vc.Nbb
	hSin := math.Sin(hRad)
	hCos := math.Cos(hRad)
	gamma := 23.0 * (p2 + 0.305) * t / (23.0*p1 + 11.0*t*hCos + 108.0*t*hSin)
	a := gamma * hCos
	b := gamma * hSin
	rA := (460.0*p2 + 451.0*a + 288.0*b) / 1403.0
	gA := (460.0*p2 - 891.0*a - 261.0*b) / 1403.0
	bA := (460.0*p2 - 220.0*a - 6300.0*b) / 1403.0

	rCBase := math.Max(0, (27.13*math.Abs(rA))/(400.0-math.Abs(rA)))
	rC := utils.Signum(rA) * (100.0 / vc.Fl) * math.Pow(rCBase, 1.0/0.42)
	gCBase := math.Max(0, (27.13*math.Abs(gA))/(400.0-math.Abs(gA)))
	gC := utils.Signum(gA) * (100.0 / vc.Fl) * math.Pow(gCBase, 1.0/0.42)
	bCBase := math.Max(0, (27.13*math.Abs(bA))/(400.0-math.Abs(bA)))
	bC := utils.Signum(bA) * (100.0 / vc.Fl) * math.Pow(bCBase, 1.0/0.42)

	rX := rC / vc.RgbD[0]
	gX := gC / vc.RgbD[1]
	bX := bC / vc.RgbD[2]
	x := 1.86206786*rX - 1.01125463*gX + 0.14918677*bX
	y := 0.38752654*rX + 0.62144744*gX - 0.00897398*bX
	z := -0.01584150*rX - 0.03412294*gX + 1.04996444*bX

	rL := 3.2406*x - 1.5372*y - 0.4986*z
	gL := -0.9689*x + 1.8758*y + 0.0415*z
	bL := 0.0557*x - 0.2040*y + 1.0570*z

	red := utils.Delinearized(rL)
	green := utils.Delinearized(gL)
	blue := utils.Delinearized(bL)

	return utils.ArgbFromRgb(red, green, blue)
}

func CamFromXyzAndViewingConditions(x, y, z float64, vc ViewingConditions) Cam {
	rC := 0.401288*x + 0.650173*y - 0.051461*z
	gC := -0.250268*x + 1.204414*y + 0.045854*z
	bC := -0.002079*x + 0.048952*y + 0.953127*z

	rD := vc.RgbD[0] * rC
	gD := vc.RgbD[1] * gC
	bD := vc.RgbD[2] * bC

	rAF := math.Pow(vc.Fl*math.Abs(rD)/100.0, 0.42)
	gAF := math.Pow(vc.Fl*math.Abs(gD)/100.0, 0.42)
	bAF := math.Pow(vc.Fl*math.Abs(bD)/100.0, 0.42)
	rA := utils.Signum(rD) * 400.0 * rAF / (rAF + 27.13)
	gA := utils.Signum(gD) * 400.0 * gAF / (gAF + 27.13)
	bA := utils.Signum(bD) * 400.0 * bAF / (bAF + 27.13)

	a := (11.0*rA + -12.0*gA + bA) / 11.0
	b := (rA + gA - 2.0*bA) / 9.0
	p2 := (40.0*rA + 20.0*gA + bA) / 20.0

	radians := math.Atan2(b, a)
	degrees := radians * 180.0 / utils.Pi
	hue := utils.SanitizeDegreesDouble(degrees)
	ac := p2 * vc.Nbb

	j := 100.0 * math.Pow(ac/vc.Aw, vc.C*vc.Z)
	return CamFromJchAndViewingConditions(j, utils.Signum(ac)*math.Pow(j/100.0, 1.0/vc.C/vc.Z), hue, vc)
}

func CamFromJchAndViewingConditions(j, c, h float64, vc ViewingConditions) Cam {
	q := (4.0 / vc.C) * math.Sqrt(j/100.0) * (vc.Aw + 4.0) * vc.FlRoot
	m := c * vc.FlRoot
	alpha := 0.0
	if j != 0 {
		alpha = c / math.Sqrt(j/100.0)
	}
	s := 50.0 * math.Sqrt((alpha*vc.C)/(vc.Aw+4.0))
	hueRadians := h * utils.Pi / 180.0
	jstar := (1.0 + 100.0*0.007) * j / (1.0 + 0.007*j)
	mstar := 1.0 / 0.0228 * math.Log(1.0+0.0228*m)
	astar := mstar * math.Cos(hueRadians)
	bstar := mstar * math.Sin(hueRadians)
	return Cam{
		Hue:    h,
		Chroma: c,
		J:      j,
		Q:      q,
		M:      m,
		S:      s,
		Jstar:  jstar,
		Astar:  astar,
		Bstar:  bstar,
	}
}

func CamFromUcsAndViewingConditions(jstar, astar, bstar float64, vc ViewingConditions) Cam {
	m := math.Sqrt(astar*astar + bstar*bstar)
	m2 := (math.Exp(m*0.0228) - 1.0) / 0.0228
	c := m2 / vc.FlRoot
	h := math.Atan2(bstar, astar) * (180.0 / utils.Pi)
	if h < 0.0 {
		h += 360.0
	}
	j := jstar / (1.0 - (jstar-100.0)*0.007)
	return CamFromJchAndViewingConditions(j, c, h, vc)
}

func XyzInViewingConditions(cam Cam, vc ViewingConditions) utils.Vec3 {
	alpha := 0.0
	if cam.Chroma != 0.0 && cam.J != 0.0 {
		alpha = cam.Chroma / math.Sqrt(cam.J/100.0)
	}

	t := math.Pow(alpha/math.Pow(1.64-math.Pow(0.29, vc.BackgroundYToWhitePointY), 0.73), 1.0/0.9)
	hRad := cam.Hue * utils.Pi / 180.0

	eHue := 0.25 * (math.Cos(hRad+2.0) + 3.8)
	ac := vc.Aw * math.Pow(cam.J/100.0, 1.0/vc.C/vc.Z)
	p1 := eHue * (50000.0 / 13.0) * vc.Nc * vc.Ncb

	p2 := ac / vc.Nbb

	hSin := math.Sin(hRad)
	hCos := math.Cos(hRad)

	gamma := 23.0 * (p2 + 0.305) * t / (23.0*p1 + 11*t*hCos + 108.0*t*hSin)
	a := gamma * hCos
	b := gamma * hSin
	rA := (460.0*p2 + 451.0*a + 288.0*b) / 1403.0
	gA := (460.0*p2 - 891.0*a - 261.0*b) / 1403.0
	bA := (460.0*p2 - 220.0*a - 6300.0*b) / 1403.0

	rCBase := math.Max(0, (27.13*math.Abs(rA))/(400.0-math.Abs(rA)))
	rC := utils.Signum(rA) * (100.0 / vc.Fl) * math.Pow(rCBase, 1.0/0.42)
	gCBase := math.Max(0, (27.13*math.Abs(gA))/(400.0-math.Abs(gA)))
	gC := utils.Signum(gA) * (100.0 / vc.Fl) * math.Pow(gCBase, 1.0/0.42)
	bCBase := math.Max(0, (27.13*math.Abs(bA))/(400.0-math.Abs(bA)))
	bC := utils.Signum(bA) * (100.0 / vc.Fl) * math.Pow(bCBase, 1.0/0.42)
	rF := rC / vc.RgbD[0]
	gF := gC / vc.RgbD[1]
	bF := bC / vc.RgbD[2]

	x := 1.86206786*rF - 1.01125463*gF + 0.14918677*bF
	y := 0.38752654*rF + 0.62144744*gF - 0.00897398*bF
	z := -0.01584150*rF - 0.03412294*gF + 1.04996444*bF

	return utils.Vec3{A: x, B: y, C: z}
}

func InViewingConditions(hct *Hct, vc ViewingConditions) *Hct {
	cam16 := CamFromInt(hct.ToInt())
	viewedInVc := XyzInViewingConditions(cam16, vc)

	recastInVc := CamFromXyzAndViewingConditions(viewedInVc.A, viewedInVc.B, viewedInVc.C, DefaultViewingConditions)

	return NewHct(recastInVc.Hue, recastInVc.Chroma, utils.LstarFromY(viewedInVc.B))
}
