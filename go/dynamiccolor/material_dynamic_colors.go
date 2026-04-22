package dynamiccolor

import (
	"math"
	"material-color-utilities/cam"
	"material-color-utilities/palettes"
)

type MaterialDynamicColors struct{}

func (m MaterialDynamicColors) highestSurface(s *DynamicScheme) *DynamicColor {
	if s.IsDark {
		return m.SurfaceBright()
	}
	return m.SurfaceDim()
}

func (m MaterialDynamicColors) PrimaryPaletteKeyColor() *DynamicColor {
	return FromPalette(
		"primary_palette_key_color",
		func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		func(s *DynamicScheme) float64 { return s.PrimaryPalette.KeyColor().Tone() },
	)
}

func (m MaterialDynamicColors) SecondaryPaletteKeyColor() *DynamicColor {
	return FromPalette(
		"secondary_palette_key_color",
		func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		func(s *DynamicScheme) float64 { return s.SecondaryPalette.KeyColor().Tone() },
	)
}

func (m MaterialDynamicColors) TertiaryPaletteKeyColor() *DynamicColor {
	return FromPalette(
		"tertiary_palette_key_color",
		func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		func(s *DynamicScheme) float64 { return s.TertiaryPalette.KeyColor().Tone() },
	)
}

func (m MaterialDynamicColors) NeutralPaletteKeyColor() *DynamicColor {
	return FromPalette(
		"neutral_palette_key_color",
		func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		func(s *DynamicScheme) float64 { return s.NeutralPalette.KeyColor().Tone() },
	)
}

func (m MaterialDynamicColors) NeutralVariantPaletteKeyColor() *DynamicColor {
	return FromPalette(
		"neutral_variant_palette_key_color",
		func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralVariantPalette },
		func(s *DynamicScheme) float64 { return s.NeutralVariantPalette.KeyColor().Tone() },
	)
}

func (m MaterialDynamicColors) Background() *DynamicColor {
	return &DynamicColor{
		Name:    "background",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 6.0
			}
			return 98.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) OnBackground() *DynamicColor {
	return &DynamicColor{
		Name:    "on_background",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 90.0
			}
			return 10.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.Background() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 3.0, Medium: 4.5, High: 7.0},
	}
}

func (m MaterialDynamicColors) Surface() *DynamicColor {
	return &DynamicColor{
		Name:    "surface",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 6.0
			}
			return 98.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceDim() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_dim",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 6.0
			}
			return ContrastCurve{Low: 87.0, Normal: 87.0, Medium: 80.0, High: 75.0}.Get(s.ContrastLevel)
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceBright() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_bright",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 24.0, Normal: 24.0, Medium: 29.0, High: 34.0}.Get(s.ContrastLevel)
			}
			return 98.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceContainerLowest() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_container_lowest",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 4.0, Normal: 4.0, Medium: 2.0, High: 0.0}.Get(s.ContrastLevel)
			}
			return 100.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceContainerLow() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_container_low",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 10.0, Normal: 10.0, Medium: 11.0, High: 12.0}.Get(s.ContrastLevel)
			}
			return ContrastCurve{Low: 96.0, Normal: 96.0, Medium: 96.0, High: 95.0}.Get(s.ContrastLevel)
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 12.0, Normal: 12.0, Medium: 16.0, High: 20.0}.Get(s.ContrastLevel)
			}
			return ContrastCurve{Low: 94.0, Normal: 94.0, Medium: 92.0, High: 90.0}.Get(s.ContrastLevel)
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceContainerHigh() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_container_high",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 17.0, Normal: 17.0, Medium: 21.0, High: 25.0}.Get(s.ContrastLevel)
			}
			return ContrastCurve{Low: 92.0, Normal: 92.0, Medium: 88.0, High: 85.0}.Get(s.ContrastLevel)
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) SurfaceContainerHighest() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_container_highest",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return ContrastCurve{Low: 22.0, Normal: 22.0, Medium: 26.0, High: 30.0}.Get(s.ContrastLevel)
			}
			return ContrastCurve{Low: 90.0, Normal: 90.0, Medium: 84.0, High: 80.0}.Get(s.ContrastLevel)
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) OnSurface() *DynamicColor {
	return &DynamicColor{
		Name:    "on_surface",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 90.0
			}
			return 10.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) SurfaceVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralVariantPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 30.0
			}
			return 90.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) OnSurfaceVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "on_surface_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralVariantPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 80.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) InverseSurface() *DynamicColor {
	return &DynamicColor{
		Name:    "inverse_surface",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 90.0
			}
			return 20.0
		},
	}
}

func (m MaterialDynamicColors) InverseOnSurface() *DynamicColor {
	return &DynamicColor{
		Name:    "inverse_on_surface",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 20.0
			}
			return 95.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.InverseSurface() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) Outline() *DynamicColor {
	return &DynamicColor{
		Name:    "outline",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralVariantPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 60.0
			}
			return 50.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.5, Normal: 3.0, Medium: 4.5, High: 7.0},
	}
}

func (m MaterialDynamicColors) OutlineVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "outline_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralVariantPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 30.0
			}
			return 80.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
	}
}

func (m MaterialDynamicColors) Shadow() *DynamicColor {
	return &DynamicColor{
		Name:    "shadow",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone:    func(s *DynamicScheme) float64 { return 0.0 },
	}
}

func (m MaterialDynamicColors) Scrim() *DynamicColor {
	return &DynamicColor{
		Name:    "scrim",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.NeutralPalette },
		Tone:    func(s *DynamicScheme) float64 { return 0.0 },
	}
}

func (m MaterialDynamicColors) SurfaceTint() *DynamicColor {
	return &DynamicColor{
		Name:    "surface_tint",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 80.0
			}
			return 40.0
		},
		IsBackground: true,
	}
}

func (m MaterialDynamicColors) Primary() *DynamicColor {
	return &DynamicColor{
		Name:    "primary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 100.0
				}
				return 0.0
			}
			if s.IsDark {
				return 80.0
			}
			return 40.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 7.0},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.PrimaryContainer(), m.Primary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnPrimary() *DynamicColor {
	return &DynamicColor{
		Name:    "on_primary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 10.0
				}
				return 90.0
			}
			if s.IsDark {
				return 20.0
			}
			return 100.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.Primary() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) PrimaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "primary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return s.SourceColorHct.Tone()
			}
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 85.0
				}
				return 25.0
			}
			if s.IsDark {
				return 30.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.PrimaryContainer(), m.Primary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnPrimaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "on_primary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return ForegroundTone(m.PrimaryContainer().Tone(s), 4.5)
			}
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 0.0
				}
				return 100.0
			}
			if s.IsDark {
				return 90.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.PrimaryContainer() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) InversePrimary() *DynamicColor {
	return &DynamicColor{
		Name:    "inverse_primary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 40.0
			}
			return 80.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.InverseSurface() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 7.0},
	}
}

func (m MaterialDynamicColors) Secondary() *DynamicColor {
	return &DynamicColor{
		Name:    "secondary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 80.0
			}
			return 40.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 7.0},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.SecondaryContainer(), m.Secondary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnSecondary() *DynamicColor {
	return &DynamicColor{
		Name:    "on_secondary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 10.0
				}
				return 90.0
			}
			if s.IsDark {
				return 20.0
			}
			return 100.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.Secondary() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) SecondaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "secondary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			initialTone := 90.0
			if s.IsDark {
				initialTone = 30.0
			}
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 30.0
				}
				return 85.0
			}
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return FindDesiredChromaByTone(s.SecondaryPalette.Hue(), s.SecondaryPalette.Chroma(), initialTone, s.IsDark)
			}
			return initialTone
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.SecondaryContainer(), m.Secondary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnSecondaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "on_secondary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 90.0
				}
				return 10.0
			}
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return ForegroundTone(m.SecondaryContainer().Tone(s), 4.5)
			}
			if s.IsDark {
				return 90.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.SecondaryContainer() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) Tertiary() *DynamicColor {
	return &DynamicColor{
		Name:    "tertiary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 90.0
				}
				return 25.0
			}
			if s.IsDark {
				return 80.0
			}
			return 40.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 7.0},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.TertiaryContainer(), m.Tertiary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnTertiary() *DynamicColor {
	return &DynamicColor{
		Name:    "on_tertiary",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 10.0
				}
				return 90.0
			}
			if s.IsDark {
				return 20.0
			}
			return 100.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.Tertiary() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) TertiaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "tertiary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 60.0
				}
				return 49.0
			}
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return FindDesiredChromaByTone(s.TertiaryPalette.Hue(), s.TertiaryPalette.Chroma(), s.SourceColorHct.Tone(), s.IsDark)
			}
			if s.IsDark {
				return 30.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.TertiaryContainer(), m.Tertiary(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnTertiaryContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "on_tertiary_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 0.0
				}
				return 100.0
			}
			if s.Variant == VariantFidelity || s.Variant == VariantContent {
				return ForegroundTone(m.TertiaryContainer().Tone(s), 4.5)
			}
			if s.IsDark {
				return 90.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.TertiaryContainer() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) Error() *DynamicColor {
	return &DynamicColor{
		Name:    "error",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.ErrorPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 80.0
			}
			return 40.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 7.0},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.ErrorContainer(), m.Error(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnError() *DynamicColor {
	return &DynamicColor{
		Name:    "on_error",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.ErrorPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 20.0
			}
			return 100.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.Error() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) ErrorContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "error_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.ErrorPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 30.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.ErrorContainer(), m.Error(), 10.0, TonePolarityNearer, false)
		},
	}
}

func (m MaterialDynamicColors) OnErrorContainer() *DynamicColor {
	return &DynamicColor{
		Name:    "on_error_container",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.ErrorPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.IsDark {
				return 90.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.ErrorContainer() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) PrimaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "primary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 100.0
				}
				return 10.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.PrimaryFixed(), m.PrimaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) PrimaryFixedDim() *DynamicColor {
	return &DynamicColor{
		Name:    "primary_fixed_dim",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 90.0
				}
				return 20.0
			}
			return 80.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.PrimaryFixed(), m.PrimaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) OnPrimaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "on_primary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 10.0
				}
				return 90.0
			}
			return 10.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.PrimaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) OnPrimaryFixedVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "on_primary_fixed_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.PrimaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				if s.IsDark {
					return 30.0
				}
				return 70.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.PrimaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) SecondaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "secondary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 80.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.SecondaryFixed(), m.SecondaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) SecondaryFixedDim() *DynamicColor {
	return &DynamicColor{
		Name:    "secondary_fixed_dim",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 70.0
			}
			return 80.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.SecondaryFixed(), m.SecondaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) OnSecondaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "on_secondary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone:    func(s *DynamicScheme) float64 { return 10.0 },
		Background:    func(s *DynamicScheme) *DynamicColor { return m.SecondaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) OnSecondaryFixedVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "on_secondary_fixed_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.SecondaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 25.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.SecondaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func (m MaterialDynamicColors) TertiaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "tertiary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 40.0
			}
			return 90.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.TertiaryFixed(), m.TertiaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) TertiaryFixedDim() *DynamicColor {
	return &DynamicColor{
		Name:    "tertiary_fixed_dim",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 30.0
			}
			return 80.0
		},
		IsBackground: true,
		Background:   func(s *DynamicScheme) *DynamicColor { return m.highestSurface(s) },
		ContrastCurve: &ContrastCurve{Low: 1.0, Normal: 1.0, Medium: 3.0, High: 4.5},
		ToneDeltaPair: func(s *DynamicScheme) ToneDeltaPair {
			return NewToneDeltaPair(m.TertiaryFixed(), m.TertiaryFixedDim(), 10.0, TonePolarityLighter, true)
		},
	}
}

func (m MaterialDynamicColors) OnTertiaryFixed() *DynamicColor {
	return &DynamicColor{
		Name:    "on_tertiary_fixed",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone:    func(s *DynamicScheme) float64 { return 10.0 },
		Background:    func(s *DynamicScheme) *DynamicColor { return m.TertiaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 4.5, Normal: 7.0, Medium: 11.0, High: 21.0},
	}
}

func (m MaterialDynamicColors) OnTertiaryFixedVariant() *DynamicColor {
	return &DynamicColor{
		Name:    "on_tertiary_fixed_variant",
		Palette: func(s *DynamicScheme) *palettes.TonalPalette { return &s.TertiaryPalette },
		Tone: func(s *DynamicScheme) float64 {
			if s.Variant == VariantMonochrome {
				return 0.0
			}
			return 30.0
		},
		Background:    func(s *DynamicScheme) *DynamicColor { return m.TertiaryFixedDim() },
		ContrastCurve: &ContrastCurve{Low: 3.0, Normal: 4.5, Medium: 7.0, High: 11.0},
	}
}

func FindDesiredChromaByTone(hue, chroma, tone float64, byDecreasingTone bool) float64 {
	answer := tone

	closestToChroma := cam.NewHct(hue, chroma, tone)
	if closestToChroma.Chroma() < chroma {
		chromaPeak := closestToChroma.Chroma()
		for closestToChroma.Chroma() < chroma {
			if byDecreasingTone {
				answer -= 1.0
			} else {
				answer += 1.0
			}
			if answer < 0 || answer > 100 {
				break
			}
			potentialSolution := cam.NewHct(hue, chroma, answer)
			if chromaPeak > potentialSolution.Chroma() {
				break
			}
			if math.Abs(potentialSolution.Chroma()-chroma) < 0.4 {
				break
			}

			potentialDelta := math.Abs(potentialSolution.Chroma() - chroma)
			currentDelta := math.Abs(closestToChroma.Chroma() - chroma)
			if potentialDelta < currentDelta {
				closestToChroma = potentialSolution
			}
			chromaPeak = math.Max(chromaPeak, potentialSolution.Chroma())
		}
	}

	return answer
}
