package cam

import (
	"math"
	"testing"
	"material-color-utilities/utils"
)

func TestHctLimitedToSRGB(t *testing.T) {
	hct := NewHct(120.0, 200.0, 50.0)
	argb := hct.ToInt()

	cam := CamFromInt(argb)
	if math.Abs(cam.Hue-hct.Hue()) > 1e-9 {
		t.Errorf("Expected hue %f, got %f", hct.Hue(), cam.Hue)
	}
	if math.Abs(cam.Chroma-hct.Chroma()) > 1e-9 {
		t.Errorf("Expected chroma %f, got %f", hct.Chroma(), cam.Chroma)
	}
	if math.Abs(utils.LstarFromArgb(argb)-hct.Tone()) > 1e-9 {
		t.Errorf("Expected tone %f, got %f", hct.Tone(), utils.LstarFromArgb(argb))
	}
}

func TestHctTruncatesColors(t *testing.T) {
	hct := NewHct(120.0, 60.0, 50.0)
	chroma := hct.Chroma()
	if chroma >= 60.0 {
		t.Errorf("Expected chroma < 60.0, got %f", chroma)
	}

	hct.SetTone(180.0)
	if hct.Chroma() >= chroma {
		t.Errorf("Expected chroma < %f, got %f", chroma, hct.Chroma())
	}
}

func isOnBoundary(rgbComponent int) bool {
	return rgbComponent == 0 || rgbComponent == 255
}

func colorIsOnBoundary(argb utils.Argb) bool {
	return isOnBoundary(utils.RedFromInt(argb)) || isOnBoundary(utils.GreenFromInt(argb)) || isOnBoundary(utils.BlueFromInt(argb))
}

func TestHctCorrectness(t *testing.T) {
	hues := []float64{15, 45, 75, 105, 135, 165, 195, 225, 255, 285, 315, 345}
	chromas := []float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	tones := []float64{20, 30, 40, 50, 60, 70, 80}

	for _, hue := range hues {
		for _, chroma := range chromas {
			for _, tone := range tones {
				color := NewHct(hue, chroma, tone)

				if chroma > 0 {
					if math.Abs(color.Hue()-hue) > 4.0 && math.Abs(color.Hue()-hue) < 356.0 {
						t.Errorf("Hue mismatch: expected %f, got %f (hue=%f, chroma=%f, tone=%f)", hue, color.Hue(), hue, chroma, tone)
					}
				}

				if color.Chroma() > chroma+2.5 {
					t.Errorf("Chroma too high: expected <= %f, got %f (hue=%f, chroma=%f, tone=%f)", chroma+2.5, color.Chroma(), hue, chroma, tone)
				}

				if color.Chroma() < chroma-2.5 {
					if !colorIsOnBoundary(color.ToInt()) {
						t.Errorf("Chroma too low and not on boundary: expected >= %f, got %f (hue=%f, chroma=%f, tone=%f)", chroma-2.5, color.Chroma(), hue, chroma, tone)
					}
				}

				if math.Abs(color.Tone()-tone) > 0.5 {
					t.Errorf("Tone mismatch: expected %f, got %f (hue=%f, chroma=%f, tone=%f)", tone, color.Tone(), hue, chroma, tone)
				}
			}
		}
	}
}
