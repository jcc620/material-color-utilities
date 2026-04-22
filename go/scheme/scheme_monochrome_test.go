package scheme

import (
	"math"
	"testing"
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/dynamiccolor"
)

func TestSchemeMonochromeDark(t *testing.T) {
	scheme := NewSchemeMonochrome(cam.NewHctFromInt(0xff0000ff), true, 0.0)
	mdc := dynamiccolor.MaterialDynamicColors{}

	tests := []struct {
		name string
		hct  *cam.Hct
		want float64
	}{
		{"Primary", mdc.Primary().GetHct(scheme), 100.0},
		{"OnPrimary", mdc.OnPrimary().GetHct(scheme), 10.0},
		{"PrimaryContainer", mdc.PrimaryContainer().GetHct(scheme), 85.0},
		{"OnPrimaryContainer", mdc.OnPrimaryContainer().GetHct(scheme), 0.0},
		{"Secondary", mdc.Secondary().GetHct(scheme), 80.0},
		{"OnSecondary", mdc.OnSecondary().GetHct(scheme), 10.0},
		{"SecondaryContainer", mdc.SecondaryContainer().GetHct(scheme), 30.0},
		{"OnSecondaryContainer", mdc.OnSecondaryContainer().GetHct(scheme), 90.0},
		{"Tertiary", mdc.Tertiary().GetHct(scheme), 90.0},
		{"OnTertiary", mdc.OnTertiary().GetHct(scheme), 10.0},
		{"TertiaryContainer", mdc.TertiaryContainer().GetHct(scheme), 60.0},
		{"OnTertiaryContainer", mdc.OnTertiaryContainer().GetHct(scheme), 0.0},
	}

	for _, tt := range tests {
		if math.Abs(tt.hct.Tone()-tt.want) > 1.0 {
			t.Errorf("%s tone = %f, want %f", tt.name, tt.hct.Tone(), tt.want)
		}
	}
}

func TestSchemeMonochromeLight(t *testing.T) {
	scheme := NewSchemeMonochrome(cam.NewHctFromInt(0xff0000ff), false, 0.0)
	mdc := dynamiccolor.MaterialDynamicColors{}

	tests := []struct {
		name string
		hct  *cam.Hct
		want float64
	}{
		{"Primary", mdc.Primary().GetHct(scheme), 0.0},
		{"OnPrimary", mdc.OnPrimary().GetHct(scheme), 90.0},
		{"PrimaryContainer", mdc.PrimaryContainer().GetHct(scheme), 25.0},
		{"OnPrimaryContainer", mdc.OnPrimaryContainer().GetHct(scheme), 100.0},
		{"Secondary", mdc.Secondary().GetHct(scheme), 40.0},
		{"OnSecondary", mdc.OnSecondary().GetHct(scheme), 100.0},
		{"SecondaryContainer", mdc.SecondaryContainer().GetHct(scheme), 85.0},
		{"OnSecondaryContainer", mdc.OnSecondaryContainer().GetHct(scheme), 10.0},
		{"Tertiary", mdc.Tertiary().GetHct(scheme), 25.0},
		{"OnTertiary", mdc.OnTertiary().GetHct(scheme), 90.0},
		{"TertiaryContainer", mdc.TertiaryContainer().GetHct(scheme), 49.0},
		{"OnTertiaryContainer", mdc.OnTertiaryContainer().GetHct(scheme), 100.0},
	}

	for _, tt := range tests {
		if math.Abs(tt.hct.Tone()-tt.want) > 1.0 {
			t.Errorf("%s tone = %f, want %f", tt.name, tt.hct.Tone(), tt.want)
		}
	}
}
