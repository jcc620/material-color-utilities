package cam

import (
	"math"
	"testing"
	"github.com/jcc620/material-color-utilities/go/utils"
)

var (
	RED   utils.Argb = 0xffff0000
	GREEN utils.Argb = 0xff00ff00
	BLUE  utils.Argb = 0xff0000ff
	WHITE utils.Argb = 0xffffffff
	BLACK utils.Argb = 0xff000000
)

func TestCamRed(t *testing.T) {
	cam := CamFromInt(RED)
	if math.Abs(cam.Hue-27.408) > 1e-3 {
		t.Errorf("cam.Hue = %f, want 27.408", cam.Hue)
	}
	if math.Abs(cam.Chroma-113.357) > 1e-3 {
		t.Errorf("cam.Chroma = %f, want 113.357", cam.Chroma)
	}
	if math.Abs(cam.J-46.445) > 1e-3 {
		t.Errorf("cam.J = %f, want 46.445", cam.J)
	}
}

func TestCamGreen(t *testing.T) {
	cam := CamFromInt(GREEN)
	if math.Abs(cam.Hue-142.139) > 1e-3 {
		t.Errorf("cam.Hue = %f, want 142.139", cam.Hue)
	}
	if math.Abs(cam.Chroma-108.410) > 1e-3 {
		t.Errorf("cam.Chroma = %f, want 108.410", cam.Chroma)
	}
}

func TestCamBlue(t *testing.T) {
	cam := CamFromInt(BLUE)
	if math.Abs(cam.Hue-282.788) > 1e-3 {
		t.Errorf("cam.Hue = %f, want 282.788", cam.Hue)
	}
}

func TestCamWhite(t *testing.T) {
	cam := CamFromInt(WHITE)
	if math.Abs(cam.J-100.0) > 1e-3 {
		t.Errorf("cam.J = %f, want 100.0", cam.J)
	}
}

func TestCamBlack(t *testing.T) {
	cam := CamFromInt(BLACK)
	if math.Abs(cam.J-0.0) > 1e-3 {
		t.Errorf("cam.J = %f, want 0.0", cam.J)
	}
}

func TestCamRoundTrip(t *testing.T) {
	colors := []utils.Argb{RED, GREEN, BLUE}
	for _, color := range colors {
		cam := CamFromInt(color)
		recovered := IntFromCam(cam)
		if recovered != color {
			t.Errorf("IntFromCam(CamFromInt(%x)) = %x, want %x", color, recovered, color)
		}
	}
}
