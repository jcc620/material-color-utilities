package quantize

import (
	"testing"
	"github.com/jcc620/material-color-utilities/go/utils"
)

func TestCelebiFullImage(t *testing.T) {
	pixels := make([]utils.Argb, 12544)
	for i := range pixels {
		pixels[i] = utils.Argb(i % 8000)
	}
	QuantizeCelebi(pixels, 128)
}

func TestCelebiOneRed(t *testing.T) {
	pixels := []utils.Argb{0xffff0000}
	result := QuantizeCelebi(pixels, 256)
	if len(result.ColorToCount) != 1 {
		t.Errorf("len(result.ColorToCount) = %d, want 1", len(result.ColorToCount))
	}
	if result.ColorToCount[0xffff0000] != 1 {
		t.Errorf("result.ColorToCount[0xffff0000] = %d, want 1", result.ColorToCount[0xffff0000])
	}
}

func TestCelebiTwoRedThreeGreen(t *testing.T) {
	pixels := []utils.Argb{0xffff0000, 0xffff0000, 0xff00ff00, 0xff00ff00, 0xff00ff00}
	result := QuantizeCelebi(pixels, 256)
	if len(result.ColorToCount) != 2 {
		t.Errorf("len(result.ColorToCount) = %d, want 2", len(result.ColorToCount))
	}
	if result.ColorToCount[0xffff0000] != 2 {
		t.Errorf("result.ColorToCount[0xffff0000] = %d, want 2", result.ColorToCount[0xffff0000])
	}
	if result.ColorToCount[0xff00ff00] != 3 {
		t.Errorf("result.ColorToCount[0xff00ff00] = %d, want 3", result.ColorToCount[0xff00ff00])
	}
}
