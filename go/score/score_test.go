package score

import (
	"testing"
	"material-color-utilities/utils"
)

func TestScorePrioritizesChroma(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff000000: 1, 0xffffffff: 1, 0xff0000ff: 1,
	}

	opts := DefaultScoreOptions()
	opts.Desired = 4
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 1 {
		t.Errorf("len(ranked) = %d, want 1", len(ranked))
	}
	if ranked[0] != 0xff0000ff {
		t.Errorf("ranked[0] = %x, want 0xff0000ff", ranked[0])
	}
}

func TestScorePrioritizesChromaWhenProportionsEqual(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xffff0000: 1, 0xff00ff00: 1, 0xff0000ff: 1,
	}

	opts := DefaultScoreOptions()
	opts.Desired = 4
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xffff0000 {
		t.Errorf("ranked[0] = %x, want 0xffff0000", ranked[0])
	}
	if ranked[1] != 0xff00ff00 {
		t.Errorf("ranked[1] = %x, want 0xff00ff00", ranked[1])
	}
	if ranked[2] != 0xff0000ff {
		t.Errorf("ranked[2] = %x, want 0xff0000ff", ranked[2])
	}
}

func TestScoreGeneratesGblueWhenNoColorsAvailable(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{0xff000000: 1}

	opts := DefaultScoreOptions()
	opts.Desired = 4
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 1 {
		t.Errorf("len(ranked) = %d, want 1", len(ranked))
	}
	if ranked[0] != 0xff4285f4 {
		t.Errorf("ranked[0] = %x, want 0xff4285f4", ranked[0])
	}
}

func TestScoreDedupesNearbyHues(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff008772: 1, // H 180 C 42 T 50
		0xff318477: 1, // H 184 C 35 T 50
	}

	opts := DefaultScoreOptions()
	opts.Desired = 4
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 1 {
		t.Errorf("len(ranked) = %d, want 1", len(ranked))
	}
	if ranked[0] != 0xff008772 {
		t.Errorf("ranked[0] = %x, want 0xff008772", ranked[0])
	}
}

func TestScoreMaximizesHueDistance(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff008772: 1, // H 180 C 42 T 50
		0xff008587: 1, // H 198 C 50 T 50
		0xff007ebc: 1, // H 245 C 50 T 50
	}

	opts := DefaultScoreOptions()
	opts.Desired = 2
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 2 {
		t.Errorf("len(ranked) = %d, want 2", len(ranked))
	}
	if ranked[0] != 0xff007ebc {
		t.Errorf("ranked[0] = %x, want 0xff007ebc", ranked[0])
	}
	if ranked[1] != 0xff008772 {
		t.Errorf("ranked[1] = %x, want 0xff008772", ranked[1])
	}
}

func TestGeneratedScenarioOne(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff7ea16d: 67,
		0xffd8ccae: 67,
		0xff835c0d: 49,
	}

	opts := ScoreOptions{Desired: 3, FallbackColorArgb: 0xff8d3819, Filter: false}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xff7ea16d {
		t.Errorf("ranked[0] = %x, want 0xff7ea16d", ranked[0])
	}
	if ranked[1] != 0xffd8ccae {
		t.Errorf("ranked[1] = %x, want 0xffd8ccae", ranked[1])
	}
	if ranked[2] != 0xff835c0d {
		t.Errorf("ranked[2] = %x, want 0xff835c0d", ranked[2])
	}
}

func TestGeneratedScenarioTwo(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xffd33881: 14,
		0xff3205cc: 77,
		0xff0b48cf: 36,
		0xffa08f5d: 81,
	}

	opts := ScoreOptions{Desired: 4, FallbackColorArgb: 0xff7d772b, Filter: true}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xff3205cc {
		t.Errorf("ranked[0] = %x, want 0xff3205cc", ranked[0])
	}
	if ranked[1] != 0xffa08f5d {
		t.Errorf("ranked[1] = %x, want 0xffa08f5d", ranked[1])
	}
	if ranked[2] != 0xffd33881 {
		t.Errorf("ranked[2] = %x, want 0xffd33881", ranked[2])
	}
}

func TestGeneratedScenarioThree(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xffbe94a6: 23,
		0xffc33fd7: 42,
		0xff899f36: 90,
		0xff94c574: 82,
	}

	opts := ScoreOptions{Desired: 3, FallbackColorArgb: 0xffaa79a4, Filter: true}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xff94c574 {
		t.Errorf("ranked[0] = %x, want 0xff94c574", ranked[0])
	}
	if ranked[1] != 0xffc33fd7 {
		t.Errorf("ranked[1] = %x, want 0xffc33fd7", ranked[1])
	}
	if ranked[2] != 0xffbe94a6 {
		t.Errorf("ranked[2] = %x, want 0xffbe94a6", ranked[2])
	}
}

func TestGeneratedScenarioFour(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xffdf241c: 85, 0xff685859: 44, 0xffd06d5f: 34,
		0xff561c54: 27, 0xff713090: 88,
	}

	opts := ScoreOptions{Desired: 5, FallbackColorArgb: 0xff58c19c, Filter: false}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 2 {
		t.Errorf("len(ranked) = %d, want 2", len(ranked))
	}
	if ranked[0] != 0xffdf241c {
		t.Errorf("ranked[0] = %x, want 0xffdf241c", ranked[0])
	}
	if ranked[1] != 0xff561c54 {
		t.Errorf("ranked[1] = %x, want 0xff561c54", ranked[1])
	}
}

func TestGeneratedScenarioFive(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xffbe66f8: 41, 0xff4bbda9: 88, 0xff80f6f9: 44,
		0xffab8017: 43, 0xffe89307: 65,
	}

	opts := ScoreOptions{Desired: 3, FallbackColorArgb: 0xff916691, Filter: false}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xffab8017 {
		t.Errorf("ranked[0] = %x, want 0xffab8017", ranked[0])
	}
	if ranked[1] != 0xff4bbda9 {
		t.Errorf("ranked[1] = %x, want 0xff4bbda9", ranked[1])
	}
	if ranked[2] != 0xffbe66f8 {
		t.Errorf("ranked[2] = %x, want 0xffbe66f8", ranked[2])
	}
}

func TestGeneratedScenarioSix(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff18ea8f: 93, 0xff327593: 18, 0xff066a18: 53,
		0xfffa8a23: 74, 0xff04ca1f: 62,
	}

	opts := ScoreOptions{Desired: 2, FallbackColorArgb: 0xff4c377a, Filter: false}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 2 {
		t.Errorf("len(ranked) = %d, want 2", len(ranked))
	}
	if ranked[0] != 0xff18ea8f {
		t.Errorf("ranked[0] = %x, want 0xff18ea8f", ranked[0])
	}
	if ranked[1] != 0xfffa8a23 {
		t.Errorf("ranked[1] = %x, want 0xfffa8a23", ranked[1])
	}
}

func TestGeneratedScenarioSeven(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff2e05ed: 23, 0xff153e55: 90, 0xff9ab220: 23,
		0xff153379: 66, 0xff68bcc3: 81,
	}

	opts := ScoreOptions{Desired: 2, FallbackColorArgb: 0xfff588dc, Filter: true}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 2 {
		t.Errorf("len(ranked) = %d, want 2", len(ranked))
	}
	if ranked[0] != 0xff2e05ed {
		t.Errorf("ranked[0] = %x, want 0xff2e05ed", ranked[0])
	}
	if ranked[1] != 0xff9ab220 {
		t.Errorf("ranked[1] = %x, want 0xff9ab220", ranked[1])
	}
}

func TestGeneratedScenarioEight(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff816ec5: 24,
		0xff6dcb94: 19,
		0xff3cae91: 98,
		0xff5b542f: 25,
	}

	opts := ScoreOptions{Desired: 1, FallbackColorArgb: 0xff84b0fd, Filter: false}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 1 {
		t.Errorf("len(ranked) = %d, want 1", len(ranked))
	}
	if ranked[0] != 0xff3cae91 {
		t.Errorf("ranked[0] = %x, want 0xff3cae91", ranked[0])
	}
}

func TestGeneratedScenarioNine(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff206f86: 52, 0xff4a620d: 96, 0xfff51401: 85,
		0xff2b8ebf: 3, 0xff277766: 59,
	}

	opts := ScoreOptions{Desired: 3, FallbackColorArgb: 0xff02b415, Filter: true}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xfff51401 {
		t.Errorf("ranked[0] = %x, want 0xfff51401", ranked[0])
	}
	if ranked[1] != 0xff4a620d {
		t.Errorf("ranked[1] = %x, want 0xff4a620d", ranked[1])
	}
	if ranked[2] != 0xff2b8ebf {
		t.Errorf("ranked[2] = %x, want 0xff2b8ebf", ranked[2])
	}
}

func TestGeneratedScenarioTen(t *testing.T) {
	argbToPopulation := map[utils.Argb]uint32{
		0xff8b1d99: 54,
		0xff27effe: 43,
		0xff6f558d: 2,
		0xff77fdf2: 78,
	}

	opts := ScoreOptions{Desired: 4, FallbackColorArgb: 0xff5e7a10, Filter: true}
	ranked := RankedSuggestions(argbToPopulation, opts)

	if len(ranked) != 3 {
		t.Errorf("len(ranked) = %d, want 3", len(ranked))
	}
	if ranked[0] != 0xff27effe {
		t.Errorf("ranked[0] = %x, want 0xff27effe", ranked[0])
	}
	if ranked[1] != 0xff8b1d99 {
		t.Errorf("ranked[1] = %x, want 0xff8b1d99", ranked[1])
	}
	if ranked[2] != 0xff6f558d {
		t.Errorf("ranked[2] = %x, want 0xff6f558d", ranked[2])
	}
}
