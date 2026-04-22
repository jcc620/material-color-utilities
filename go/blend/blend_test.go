package blend

import (
	"testing"
	"material-color-utilities/utils"
)

func TestRedToBlue(t *testing.T) {
	blended := HctHue(0xffff0000, 0xff0000ff, 0.8)
	if got, want := utils.HexFromArgb(blended), "ff905eff"; got != want {
		t.Errorf("HctHue(0xffff0000, 0xff0000ff, 0.8) = %s, want %s", got, want)
	}
}
