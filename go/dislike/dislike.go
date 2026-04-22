package dislike

import (
	"math"
	"github.com/jcc620/material-color-utilities/go/cam"
)

func IsDisliked(hct *cam.Hct) bool {
	roundedHue := math.Round(hct.Hue())

	huePasses := roundedHue >= 90.0 && roundedHue <= 111.0
	chromaPasses := math.Round(hct.Chroma()) > 16.0
	tonePasses := math.Round(hct.Tone()) < 65.0

	return huePasses && chromaPasses && tonePasses
}

func FixIfDisliked(hct *cam.Hct) *cam.Hct {
	if IsDisliked(hct) {
		return cam.NewHct(hct.Hue(), hct.Chroma(), 70.0)
	}

	return hct
}
