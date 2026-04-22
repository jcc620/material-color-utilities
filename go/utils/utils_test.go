package utils

import (
	"math"
	"testing"
)

var kMatrix = [3][3]float64{
	{1, 2, 3},
	{-4, 5, -6},
	{-7, -8, -9},
}

func TestArgbFromRgb(t *testing.T) {
	if got := ArgbFromRgb(0, 0, 0); got != 0xff000000 {
		t.Errorf("ArgbFromRgb(0, 0, 0) = %x, want %x", got, 0xff000000)
	}
	if got := ArgbFromRgb(255, 255, 255); got != 0xffffffff {
		t.Errorf("ArgbFromRgb(255, 255, 255) = %x, want %x", got, 0xffffffff)
	}
	if got := ArgbFromRgb(50, 150, 250); got != 0xff3296fa {
		t.Errorf("ArgbFromRgb(50, 150, 250) = %x, want %x", got, 0xff3296fa)
	}
}

func TestSignum(t *testing.T) {
	if got := Signum(0.001); got != 1 {
		t.Errorf("Signum(0.001) = %f, want 1", got)
	}
	if got := Signum(3.0); got != 1 {
		t.Errorf("Signum(3.0) = %f, want 1", got)
	}
	if got := Signum(100.0); got != 1 {
		t.Errorf("Signum(100.0) = %f, want 1", got)
	}
	if got := Signum(-0.002); got != -1 {
		t.Errorf("Signum(-0.002) = %f, want -1", got)
	}
	if got := Signum(-4.0); got != -1 {
		t.Errorf("Signum(-4.0) = %f, want -1", got)
	}
	if got := Signum(-101.0); got != -1 {
		t.Errorf("Signum(-101.0) = %f, want -1", got)
	}
	if got := Signum(0.0); got != 0 {
		t.Errorf("Signum(0.0) = %f, want 0", got)
	}
}

func TestRotationDirection(t *testing.T) {
	tests := []struct {
		from, to float64
		want     float64
	}{
		{0.0, 30.0, 1.0},
		{0.0, 60.0, 1.0},
		{0.0, 150.0, 1.0},
		{90.0, 240.0, 1.0},
		{300.0, 30.0, 1.0},
		{270.0, 60.0, 1.0},
		{360.0 * 2, 15.0, 1.0},
		{360.0*3 + 15.0, -360.0*4 + 30.0, 1.0},
		{30.0, 0.0, -1.0},
		{60.0, 0.0, -1.0},
		{150.0, 0.0, -1.0},
		{240.0, 90.0, -1.0},
		{30.0, 300.0, -1.0},
		{60.0, 270.0, -1.0},
		{15.0, -360.0 * 2, -1.0},
		{-360.0*4 + 270.0, 360.0*5 + 180.0, -1.0},
	}
	for _, tt := range tests {
		if got := RotationDirection(tt.from, tt.to); got != tt.want {
			t.Errorf("RotationDirection(%f, %f) = %f, want %f", tt.from, tt.to, got, tt.want)
		}
	}
}

func TestDiffDegrees(t *testing.T) {
	tests := []struct {
		a, b float64
		want float64
	}{
		{0.0, 30.0, 30.0},
		{0.0, 60.0, 60.0},
		{0.0, 150.0, 150.0},
		{90.0, 240.0, 150.0},
		{300.0, 30.0, 90.0},
		{270.0, 60.0, 150.0},
		{30.0, 0.0, 30.0},
		{60.0, 0.0, 60.0},
		{150.0, 0.0, 150.0},
		{240.0, 90.0, 150.0},
		{30.0, 300.0, 90.0},
		{60.0, 270.0, 150.0},
	}
	for _, tt := range tests {
		if got := DiffDegrees(tt.a, tt.b); got != tt.want {
			t.Errorf("DiffDegrees(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestAngleSanitation(t *testing.T) {
	if got := SanitizeDegreesInt(30); got != 30 {
		t.Errorf("SanitizeDegreesInt(30) = %d, want 30", got)
	}
	if got := SanitizeDegreesInt(240); got != 240 {
		t.Errorf("SanitizeDegreesInt(240) = %d, want 240", got)
	}
	if got := SanitizeDegreesInt(360); got != 0 {
		t.Errorf("SanitizeDegreesInt(360) = %d, want 0", got)
	}
	if got := SanitizeDegreesInt(-30); got != 330 {
		t.Errorf("SanitizeDegreesInt(-30) = %d, want 330", got)
	}
	if got := SanitizeDegreesInt(-750); got != 330 {
		t.Errorf("SanitizeDegreesInt(-750) = %d, want 330", got)
	}
	if got := SanitizeDegreesInt(-54321); got != 39 {
		t.Errorf("SanitizeDegreesInt(-54321) = %d, want 39", got)
	}

	doubleTests := []struct {
		in   float64
		want float64
	}{
		{30.0, 30.0},
		{240.0, 240.0},
		{360.0, 0.0},
		{-30.0, 330.0},
		{-750.0, 330.0},
		{-54321.0, 39.0},
		{360.125, 0.125},
		{-11111.11, 48.89},
	}
	for _, tt := range doubleTests {
		if got := SanitizeDegreesDouble(tt.in); math.Abs(got-tt.want) > 1e-4 {
			t.Errorf("SanitizeDegreesDouble(%f) = %f, want %f", tt.in, got, tt.want)
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	v1 := MatrixMultiply(Vec3{A: 1, B: 3, C: 5}, kMatrix)
	if math.Abs(v1.A-22) > 1e-4 || math.Abs(v1.B+19) > 1e-4 || math.Abs(v1.C+76) > 1e-4 {
		t.Errorf("MatrixMultiply({1, 3, 5}, kMatrix) = %+v, want {22, -19, -76}", v1)
	}

	v2 := MatrixMultiply(Vec3{A: -11.1, B: 22.2, C: -33.3}, kMatrix)
	if math.Abs(v2.A+66.6) > 1e-4 || math.Abs(v2.B-355.2) > 1e-4 || math.Abs(v2.C-199.8) > 1e-4 {
		t.Errorf("MatrixMultiply({-11.1, 22.2, -33.3}, kMatrix) = %+v, want {-66.6, 355.2, 199.8}", v2)
	}
}

func TestAlphaFromInt(t *testing.T) {
	if got := AlphaFromInt(0xff123456); got != 0xff {
		t.Errorf("AlphaFromInt(0xff123456) = %x, want ff", got)
	}
}

func TestRedFromInt(t *testing.T) {
	if got := RedFromInt(0xff123456); got != 0x12 {
		t.Errorf("RedFromInt(0xff123456) = %x, want 12", got)
	}
}

func TestGreenFromInt(t *testing.T) {
	if got := GreenFromInt(0xff123456); got != 0x34 {
		t.Errorf("GreenFromInt(0xff123456) = %x, want 34", got)
	}
}

func TestBlueFromInt(t *testing.T) {
	if got := BlueFromInt(0xff123456); got != 0x56 {
		t.Errorf("BlueFromInt(0xff123456) = %x, want 56", got)
	}
}

func TestIsOpaque(t *testing.T) {
	if !IsOpaque(0xff123456) {
		t.Error("IsOpaque(0xff123456) = false, want true")
	}
	if IsOpaque(0xf0123456) {
		t.Error("IsOpaque(0xf0123456) = true, want false")
	}
}

func TestLinearizedComponents(t *testing.T) {
	tests := []struct {
		in   int
		want float64
	}{
		{0, 0.0},
		{1, 0.0303527},
		{2, 0.0607054},
		{8, 0.242822},
		{9, 0.273174},
		{16, 0.518152},
		{32, 1.44438},
		{64, 5.12695},
		{128, 21.5861},
		{255, 100.0},
	}
	for _, tt := range tests {
		if got := Linearized(tt.in); math.Abs(got-tt.want) > 1e-4 {
			t.Errorf("Linearized(%d) = %f, want %f", tt.in, got, tt.want)
		}
	}
}

func TestDelinearizedComponents(t *testing.T) {
	tests := []struct {
		in   float64
		want int
	}{
		{0.0, 0},
		{0.0303527, 1},
		{0.0607054, 2},
		{0.242822, 8},
		{0.273174, 9},
		{0.518152, 16},
		{1.44438, 32},
		{5.12695, 64},
		{21.5861, 128},
		{100.0, 255},
		{25.0, 137},
		{50.0, 188},
		{75.0, 225},
		{-1.0, 0},
		{101.0, 255},
	}
	for _, tt := range tests {
		if got := Delinearized(tt.in); got != tt.want {
			t.Errorf("Delinearized(%f) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestLstarFromArgb(t *testing.T) {
	if got := LstarFromArgb(0xff89bce1); math.Abs(got-74.011) > 1e-3 {
		t.Errorf("LstarFromArgb(0xff89bce1) = %f, want 74.011", got)
	}
}

func TestHexFromArgb(t *testing.T) {
	if got := HexFromArgb(0xff89bce1); got != "ff89bce1" {
		t.Errorf("HexFromArgb(0xff89bce1) = %s, want ff89bce1", got)
	}
}

func TestIntFromLstar(t *testing.T) {
	tests := []struct {
		in   float64
		want Argb
	}{
		{0.0, 0xff000000},
		{50.0, 0xff777777},
		{100.0, 0xffffffff},
		{-1.0, 0xff000000},
		{101.0, 0xffffffff},
	}
	for _, tt := range tests {
		if got := IntFromLstar(tt.in); got != tt.want {
			t.Errorf("IntFromLstar(%f) = %x, want %x", tt.in, got, tt.want)
		}
	}
}

func TestYFromLstar(t *testing.T) {
	tests := []struct {
		in   float64
		want float64
	}{
		{0.0, 0.0},
		{50.0, 18.4186518},
		{100.0, 100.0},
	}
	for _, tt := range tests {
		if got := YFromLstar(tt.in); math.Abs(got-tt.want) > 1e-5 {
			t.Errorf("YFromLstar(%f) = %f, want %f", tt.in, got, tt.want)
		}
	}
}

func TestLstarFromY(t *testing.T) {
	tests := []struct {
		in   float64
		want float64
	}{
		{0.0, 0.0},
		{18.4186518, 50.0},
		{100.0, 100.0},
	}
	for _, tt := range tests {
		if got := LstarFromY(tt.in); math.Abs(got-tt.want) > 1e-5 {
			t.Errorf("LstarFromY(%f) = %f, want %f", tt.in, got, tt.want)
		}
	}
}
