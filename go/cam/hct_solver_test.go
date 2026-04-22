package cam

import (
	"testing"
	"material-color-utilities/utils"
)

func TestHctSolverRed(t *testing.T) {
	color := utils.Argb(0xFFFE0315)
	cam := CamFromInt(color)
	tone := utils.LstarFromArgb(color)
	recovered := SolveToInt(cam.Hue, cam.Chroma, tone)
	if recovered != color {
		t.Errorf("SolveToInt = %x, want %x", recovered, color)
	}
}

func TestHctSolverGreen(t *testing.T) {
	color := utils.Argb(0xFF15FE03)
	cam := CamFromInt(color)
	tone := utils.LstarFromArgb(color)
	recovered := SolveToInt(cam.Hue, cam.Chroma, tone)
	if recovered != color {
		t.Errorf("SolveToInt = %x, want %x", recovered, color)
	}
}

func TestHctSolverBlue(t *testing.T) {
	color := utils.Argb(0xFF0315FE)
	cam := CamFromInt(color)
	tone := utils.LstarFromArgb(color)
	recovered := SolveToInt(cam.Hue, cam.Chroma, tone)
	if recovered != color {
		t.Errorf("SolveToInt = %x, want %x", recovered, color)
	}
}

// Exhaustive test is omitted as it takes too long for standard test runs.
