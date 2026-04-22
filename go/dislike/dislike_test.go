package dislike

import (
	"testing"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/utils"
)

func TestMonkSkinToneScaleColorsLiked(t *testing.T) {
	colors := []utils.Argb{
		0xfff6ede4, 0xfff3e7db, 0xfff7ead0, 0xffeadaba,
		0xffd7bd96, 0xffa07e56, 0xff825c43, 0xff604134,
		0xff3a312a, 0xff292420,
	}
	for _, argb := range colors {
		if IsDisliked(cam.NewHctFromInt(argb)) {
			t.Errorf("IsDisliked(%x) = true, want false", argb)
		}
	}
}

func TestBileColorsDisliked(t *testing.T) {
	colors := []utils.Argb{
		0xff95884B, 0xff716B40, 0xffB08E00, 0xff4C4308,
		0xff464521,
	}
	for _, argb := range colors {
		if !IsDisliked(cam.NewHctFromInt(argb)) {
			t.Errorf("IsDisliked(%x) = false, want true", argb)
		}
	}
}

func TestBileColorsFixed(t *testing.T) {
	colors := []utils.Argb{
		0xff95884B, 0xff716B40, 0xffB08E00, 0xff4C4308,
		0xff464521,
	}
	for _, argb := range colors {
		bileColor := cam.NewHctFromInt(argb)
		if !IsDisliked(bileColor) {
			t.Errorf("IsDisliked(%x) = false, want true", argb)
		}
		fixedBileColor := FixIfDisliked(bileColor)
		if IsDisliked(fixedBileColor) {
			t.Errorf("IsDisliked(FixIfDisliked(%x)) = true, want false", argb)
		}
	}
}

func TestTone67Liked(t *testing.T) {
	color := cam.NewHct(100.0, 50.0, 67.0)
	if IsDisliked(color) {
		t.Error("IsDisliked(Hct(100, 50, 67)) = true, want false")
	}
	if FixIfDisliked(color).ToInt() != color.ToInt() {
		t.Errorf("FixIfDisliked(Hct(100, 50, 67)).ToInt() = %x, want %x", FixIfDisliked(color).ToInt(), color.ToInt())
	}
}
