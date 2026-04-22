package cam

import (
	"material-color-utilities/utils"
)

type Hct struct {
	hue    float64
	chroma float64
	tone   float64
	argb   utils.Argb
}

func NewHct(hue, chroma, tone float64) *Hct {
	h := &Hct{
		hue:    hue,
		chroma: chroma,
		tone:   tone,
	}
	h.setInternalState(SolveToInt(hue, chroma, tone))
	return h
}

func NewHctFromInt(argb utils.Argb) *Hct {
	h := &Hct{}
	h.setInternalState(argb)
	return h
}

func (h *Hct) Hue() float64 {
	return h.hue
}

func (h *Hct) Chroma() float64 {
	return h.chroma
}

func (h *Hct) Tone() float64 {
	return h.tone
}

func (h *Hct) ToInt() utils.Argb {
	return h.argb
}

func (h *Hct) SetHue(newHue float64) {
	h.setInternalState(SolveToInt(newHue, h.chroma, h.tone))
}

func (h *Hct) SetChroma(newChroma float64) {
	h.setInternalState(SolveToInt(h.hue, newChroma, h.tone))
}

func (h *Hct) SetTone(newTone float64) {
	h.setInternalState(SolveToInt(h.hue, h.chroma, newTone))
}

func (h *Hct) Less(other *Hct) bool {
	return h.hue < other.hue
}

func (h *Hct) setInternalState(argb utils.Argb) {
	h.argb = argb
	cam := CamFromInt(argb)
	h.hue = cam.Hue
	h.chroma = cam.Chroma
	h.tone = utils.LstarFromArgb(argb)
}
