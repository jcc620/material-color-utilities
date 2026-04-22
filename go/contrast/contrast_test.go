package contrast

import (
	"math"
	"testing"
)

func TestRatioOfTonesOutOfBoundsInput(t *testing.T) {
	if got, want := RatioOfTones(-10.0, 110.0), 21.0; math.Abs(got-want) > 0.001 {
		t.Errorf("RatioOfTones(-10.0, 110.0) = %f, want %f", got, want)
	}
}

func TestLighterImpossibleRatioErrors(t *testing.T) {
	if got, want := Lighter(90.0, 10.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Lighter(90.0, 10.0) = %f, want %f", got, want)
	}
}

func TestLighterOutOfBoundsInputAboveErrors(t *testing.T) {
	if got, want := Lighter(110.0, 2.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Lighter(110.0, 2.0) = %f, want %f", got, want)
	}
}

func TestLighterOutOfBoundsInputBelowErrors(t *testing.T) {
	if got, want := Lighter(-10.0, 2.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Lighter(-10.0, 2.0) = %f, want %f", got, want)
	}
}

func TestLighterUnsafeReturnsMaxTone(t *testing.T) {
	if got, want := LighterUnsafe(100.0, 2.0), 100.0; math.Abs(got-want) > 0.001 {
		t.Errorf("LighterUnsafe(100.0, 2.0) = %f, want %f", got, want)
	}
}

func TestDarkerImpossibleRatioErrors(t *testing.T) {
	if got, want := Darker(10.0, 20.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Darker(10.0, 20.0) = %f, want %f", got, want)
	}
}

func TestDarkerOutOfBoundsInputAboveErrors(t *testing.T) {
	if got, want := Darker(110.0, 2.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Darker(110.0, 2.0) = %f, want %f", got, want)
	}
}

func TestDarkerOutOfBoundsInputBelowErrors(t *testing.T) {
	if got, want := Darker(-10.0, 2.0), -1.0; math.Abs(got-want) > 0.001 {
		t.Errorf("Darker(-10.0, 2.0) = %f, want %f", got, want)
	}
}

func TestDarkerUnsafeReturnsMinTone(t *testing.T) {
	if got, want := DarkerUnsafe(0.0, 2.0), 0.0; math.Abs(got-want) > 0.001 {
		t.Errorf("DarkerUnsafe(0.0, 2.0) = %f, want %f", got, want)
	}
}
