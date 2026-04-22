package utils

import (
	"fmt"
	"math"
)

type Argb uint32

type Vec3 struct {
	A, B, C float64
}

const Pi = 3.141592653589793

func RedFromInt(argb Argb) int {
	return int((argb & 0x00ff0000) >> 16)
}

func GreenFromInt(argb Argb) int {
	return int((argb & 0x0000ff00) >> 8)
}

func BlueFromInt(argb Argb) int {
	return int(argb & 0x000000ff)
}

func AlphaFromInt(argb Argb) int {
	return int((argb & 0xff000000) >> 24)
}

func ArgbFromRgb(red, green, blue int) Argb {
	return Argb((0xFF << 24) | ((red & 0xff) << 16) | ((green & 0xff) << 8) | (blue & 0xff))
}

func ArgbFromLinrgb(linrgb Vec3) Argb {
	r := Delinearized(linrgb.A)
	g := Delinearized(linrgb.B)
	b := Delinearized(linrgb.C)
	return Argb((0xFF << 24) | ((r & 0xff) << 16) | ((g & 0xff) << 8) | (b & 0xff))
}

func Delinearized(rgbComponent float64) int {
	normalized := rgbComponent / 100.0
	var delinearized float64
	if normalized <= 0.0031308 {
		delinearized = normalized * 12.92
	} else {
		delinearized = 1.055*math.Pow(normalized, 1.0/2.4) - 0.055
	}
	res := int(math.Round(delinearized * 255.0))
	if res < 0 {
		return 0
	}
	if res > 255 {
		return 255
	}
	return res
}

func Linearized(rgbComponent int) float64 {
	normalized := float64(rgbComponent) / 255.0
	if normalized <= 0.040449936 {
		return normalized / 12.92 * 100.0
	} else {
		return math.Pow((normalized+0.055)/1.055, 2.4) * 100.0
	}
}

func IsOpaque(argb Argb) bool {
	return AlphaFromInt(argb) == 255
}

func LstarFromArgb(argb Argb) float64 {
	red := RedFromInt(argb)
	green := GreenFromInt(argb)
	blue := BlueFromInt(argb)
	redL := Linearized(red)
	greenL := Linearized(green)
	blueL := Linearized(blue)
	y := 0.2126*redL + 0.7152*greenL + 0.0722*blueL
	return LstarFromY(y)
}

func YFromLstar(lstar float64) float64 {
	const ke = 8.0
	if lstar > ke {
		cubeRoot := (lstar + 16.0) / 116.0
		cube := cubeRoot * cubeRoot * cubeRoot
		return cube * 100.0
	} else {
		return lstar / (24389.0 / 27.0) * 100.0
	}
}

func LstarFromY(y float64) float64 {
	const e = 216.0 / 24389.0
	yNormalized := y / 100.0
	if yNormalized <= e {
		return (24389.0 / 27.0) * yNormalized
	} else {
		return 116.0*math.Pow(yNormalized, 1.0/3.0) - 16.0
	}
}

func SanitizeDegreesInt(degrees int) int {
	if degrees < 0 {
		return (degrees % 360) + 360
	} else if degrees >= 360 {
		return degrees % 360
	} else {
		return degrees
	}
}

func SanitizeDegreesDouble(degrees float64) float64 {
	if degrees < 0.0 {
		return math.Mod(degrees, 360.0) + 360.0
	} else if degrees >= 360.0 {
		return math.Mod(degrees, 360.0)
	} else {
		return degrees
	}
}

func DiffDegrees(a, b float64) float64 {
	return 180.0 - math.Abs(math.Abs(a-b)-180.0)
}

func RotationDirection(from, to float64) float64 {
	increasingDifference := SanitizeDegreesDouble(to - from)
	if increasingDifference <= 180.0 {
		return 1.0
	}
	return -1.0
}

func HexFromArgb(argb Argb) string {
	return fmt.Sprintf("%08x", argb)
}

func IntFromLstar(lstar float64) Argb {
	y := YFromLstar(lstar)
	component := Delinearized(y)
	return ArgbFromRgb(component, component, component)
}

func Signum(num float64) float64 {
	if num < 0 {
		return -1
	} else if num == 0 {
		return 0
	} else {
		return 1
	}
}

func Lerp(start, stop, amount float64) float64 {
	return (1.0-amount)*start + amount*stop
}

func MatrixMultiply(input Vec3, matrix [3][3]float64) Vec3 {
	a := input.A*matrix[0][0] + input.B*matrix[0][1] + input.C*matrix[0][2]
	b := input.A*matrix[1][0] + input.B*matrix[1][1] + input.C*matrix[1][2]
	c := input.A*matrix[2][0] + input.B*matrix[2][1] + input.C*matrix[2][2]
	return Vec3{A: a, B: b, C: c}
}
