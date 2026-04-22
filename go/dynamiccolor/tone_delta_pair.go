package dynamiccolor

type TonePolarity int

const (
	TonePolarityDarker TonePolarity = iota
	TonePolarityLighter
	TonePolarityNearer
	TonePolarityFarther
)

type ToneDeltaPair struct {
	RoleA         *DynamicColor
	RoleB         *DynamicColor
	Delta         float64
	Polarity      TonePolarity
	StayTogether bool
}

func NewToneDeltaPair(roleA, roleB *DynamicColor, delta float64, polarity TonePolarity, stayTogether bool) ToneDeltaPair {
	return ToneDeltaPair{
		RoleA:         roleA,
		RoleB:         roleB,
		Delta:         delta,
		Polarity:      polarity,
		StayTogether: stayTogether,
	}
}
