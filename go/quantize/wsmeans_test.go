package quantize

import (
	"testing"
	"github.com/jcc620/material-color-utilities/go/utils"
)

func TestWsmeansFullImage(t *testing.T) {
	pixels := make([]utils.Argb, 12544)
	for i := range pixels {
		pixels[i] = utils.Argb(i % 8000)
	}
	QuantizeWsmeans(pixels, nil, 128)
}

func TestWsmeansOneRed(t *testing.T) {
	pixels := []utils.Argb{0xffff0000}
	result := QuantizeWsmeans(pixels, nil, 256)
	if len(result.ColorToCount) != 1 {
		t.Errorf("len(result.ColorToCount) = %d, want 1", len(result.ColorToCount))
	}
	if result.ColorToCount[0xffff0000] != 1 {
		t.Errorf("result.ColorToCount[0xffff0000] = %d, want 1", result.ColorToCount[0xffff0000])
	}
}

func TestWsmeansFiveBlue(t *testing.T) {
	pixels := []utils.Argb{0xff0000ff, 0xff0000ff, 0xff0000ff, 0xff0000ff, 0xff0000ff}
	result := QuantizeWsmeans(pixels, nil, 256)
	if len(result.ColorToCount) != 1 {
		t.Errorf("len(result.ColorToCount) = %d, want 1", len(result.ColorToCount))
	}
	if result.ColorToCount[0xff0000ff] != 5 {
		t.Errorf("result.ColorToCount[0xff0000ff] = %d, want 5", result.ColorToCount[0xff0000ff])
	}
}
