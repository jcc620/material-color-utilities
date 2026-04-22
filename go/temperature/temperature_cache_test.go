package temperature

import (
	"math"
	"testing"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/utils"
)

func TestRawTemperature(t *testing.T) {
	blueHct := cam.NewHctFromInt(0xff0000ff)
	blueTemp := RawTemperature(blueHct)
	if math.Abs(blueTemp - -1.393) > 0.001 {
		t.Errorf("RawTemperature(blue) = %f, want -1.393", blueTemp)
	}

	redHct := cam.NewHctFromInt(0xffff0000)
	redTemp := RawTemperature(redHct)
	if math.Abs(redTemp - 2.351) > 0.001 {
		t.Errorf("RawTemperature(red) = %f, want 2.351", redTemp)
	}

	greenHct := cam.NewHctFromInt(0xff00ff00)
	greenTemp := RawTemperature(greenHct)
	if math.Abs(greenTemp - -0.267) > 0.001 {
		t.Errorf("RawTemperature(green) = %f, want -0.267", greenTemp)
	}

	whiteHct := cam.NewHctFromInt(0xffffffff)
	whiteTemp := RawTemperature(whiteHct)
	if math.Abs(whiteTemp - -0.5) > 0.001 {
		t.Errorf("RawTemperature(white) = %f, want -0.5", whiteTemp)
	}

	blackHct := cam.NewHctFromInt(0xff000000)
	blackTemp := RawTemperature(blackHct)
	if math.Abs(blackTemp - -0.5) > 0.001 {
		t.Errorf("RawTemperature(black) = %f, want -0.5", blackTemp)
	}
}

func TestComplement(t *testing.T) {
	tests := []struct {
		in   uint32
		want uint32
	}{
		{0xff0000ff, 0xff9d0002},
		{0xffff0000, 0xff007bfc},
		{0xff00ff00, 0xffffd2c9},
		{0xffffffff, 0xffffffff},
		{0xff000000, 0xff000000},
	}
	for _, tt := range tests {
		got := NewTemperatureCache(cam.NewHctFromInt(utils.Argb(tt.in))).GetComplement().ToInt()
		if uint32(got) != tt.want {
			t.Errorf("GetComplement(%x) = %x, want %x", tt.in, got, tt.want)
		}
	}
}

func TestAnalogous(t *testing.T) {
	blueAnalogous := NewTemperatureCache(cam.NewHctFromInt(0xff0000ff)).GetAnalogousColors()
	blueWants := []uint32{0xff00590c, 0xff00564e, 0xff0000ff, 0xff6700cc, 0xff81009f}
	for i, want := range blueWants {
		if uint32(blueAnalogous[i].ToInt()) != want {
			t.Errorf("blueAnalogous[%d] = %x, want %x", i, blueAnalogous[i].ToInt(), want)
		}
	}

	redAnalogous := NewTemperatureCache(cam.NewHctFromInt(0xffff0000)).GetAnalogousColors()
	redWants := []uint32{0xfff60082, 0xfffc004c, 0xffff0000, 0xffd95500, 0xffaf7200}
	for i, want := range redWants {
		if uint32(redAnalogous[i].ToInt()) != want {
			t.Errorf("redAnalogous[%d] = %x, want %x", i, redAnalogous[i].ToInt(), want)
		}
	}

	greenAnalogous := NewTemperatureCache(cam.NewHctFromInt(0xff00ff00)).GetAnalogousColors()
	greenWants := []uint32{0xffcee900, 0xff92f500, 0xff00ff00, 0xff00fd6f, 0xff00fab3}
	for i, want := range greenWants {
		if uint32(greenAnalogous[i].ToInt()) != want {
			t.Errorf("greenAnalogous[%d] = %x, want %x", i, greenAnalogous[i].ToInt(), want)
		}
	}

	blackAnalogous := NewTemperatureCache(cam.NewHctFromInt(0xff000000)).GetAnalogousColors()
	for i := 0; i < 5; i++ {
		if uint32(blackAnalogous[i].ToInt()) != 0xff000000 {
			t.Errorf("blackAnalogous[%d] = %x, want 0xff000000", i, blackAnalogous[i].ToInt())
		}
	}

	whiteAnalogous := NewTemperatureCache(cam.NewHctFromInt(0xffffffff)).GetAnalogousColors()
	for i := 0; i < 5; i++ {
		if uint32(whiteAnalogous[i].ToInt()) != 0xffffffff {
			t.Errorf("whiteAnalogous[%d] = %x, want 0xffffffff", i, whiteAnalogous[i].ToInt())
		}
	}
}
