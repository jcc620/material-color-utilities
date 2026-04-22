package palettes

import (
	"math"
	"testing"
	"material-color-utilities/utils"
)

func TestTonesBlue(t *testing.T) {
	color := utils.Argb(0xff0000ff)
	tonalPalette := NewTonalPaletteFromInt(color)
	tests := []struct {
		tone float64
		want string
	}{
		{100, "ffffffff"},
		{95, "fff1efff"},
		{90, "ffe0e0ff"},
		{80, "ffbec2ff"},
		{70, "ff9da3ff"},
		{60, "ff7c84ff"},
		{50, "ff5a64ff"},
		{40, "ff343dff"},
		{30, "ff0000ef"},
		{20, "ff0001ac"},
		{10, "ff00006e"},
		{0, "ff000000"},
	}
	for _, tt := range tests {
		if got := utils.HexFromArgb(tonalPalette.Get(tt.tone)); got != tt.want {
			t.Errorf("TonalPalette.Get(%f) = %s, want %s", tt.tone, got, tt.want)
		}
	}
}

func TestKeyColorExactChromaAvailable(t *testing.T) {
	palette := NewTonalPaletteFromHueChroma(50.0, 60.0)
	result := palette.KeyColor()

	if math.Abs(result.Hue()-50.0) > 10.0 {
		t.Errorf("result.Hue() = %f, want 50.0 (near)", result.Hue())
	}
	if math.Abs(result.Chroma()-60.0) > 0.5 {
		t.Errorf("result.Chroma() = %f, want 60.0 (near)", result.Chroma())
	}
	if result.Tone() <= 0 || result.Tone() >= 100 {
		t.Errorf("result.Tone() = %f, want in (0, 100)", result.Tone())
	}
}

func TestKeyColorUnusuallyHighChroma(t *testing.T) {
	palette := NewTonalPaletteFromHueChroma(149.0, 200.0)
	result := palette.KeyColor()

	if math.Abs(result.Hue()-149.0) > 10.0 {
		t.Errorf("result.Hue() = %f, want 149.0 (near)", result.Hue())
	}
	if result.Chroma() <= 89.0 {
		t.Errorf("result.Chroma() = %f, want > 89.0", result.Chroma())
	}
	if result.Tone() <= 0 || result.Tone() >= 100 {
		t.Errorf("result.Tone() = %f, want in (0, 100)", result.Tone())
	}
}

func TestKeyColorUnusuallyLowChroma(t *testing.T) {
	palette := NewTonalPaletteFromHueChroma(50.0, 3.0)
	result := palette.KeyColor()

	if math.Abs(result.Hue()-50.0) > 10.0 {
		t.Errorf("result.Hue() = %f, want 50.0 (near)", result.Hue())
	}
	if math.Abs(result.Chroma()-3.0) > 0.5 {
		t.Errorf("result.Chroma() = %f, want 3.0 (near)", result.Chroma())
	}
	if math.Abs(result.Tone()-50.0) > 0.5 {
		t.Errorf("result.Tone() = %f, want 50.0 (near)", result.Tone())
	}
}
