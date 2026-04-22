package quantize

import (
	"testing"
	"material-color-utilities/utils"
)

func TestWuFullImage(t *testing.T) {
	pixels := make([]utils.Argb, 12544)
	for i := range pixels {
		pixels[i] = utils.Argb(i % 8000)
	}
	QuantizeWu(pixels, 128)
}

func TestWuTwoRedThreeGreen(t *testing.T) {
	pixels := []utils.Argb{0xffff0000, 0xffff0000, 0xffff0000, 0xff00ff00, 0xff00ff00}
	result := QuantizeWu(pixels, 256)
	if len(result) != 2 {
		t.Errorf("len(result) = %d, want 2", len(result))
	}
}

func TestWuRedGreenBlue(t *testing.T) {
	pixels := []utils.Argb{0xffff0000, 0xff00ff00, 0xff0000ff}
	result := QuantizeWu(pixels, 256)
	if len(result) != 3 {
		t.Errorf("len(result) = %d, want 3", len(result))
	}
	// order might differ
}
