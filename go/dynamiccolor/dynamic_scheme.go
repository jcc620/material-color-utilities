package dynamiccolor

import (
	"github.com/jcc620/material-color-utilities/go/cam"
	"github.com/jcc620/material-color-utilities/go/palettes"
	"github.com/jcc620/material-color-utilities/go/utils"
)

type DynamicScheme struct {
	SourceColorHct        *cam.Hct
	Variant               Variant
	IsDark                bool
	ContrastLevel         float64
	PrimaryPalette        palettes.TonalPalette
	SecondaryPalette      palettes.TonalPalette
	TertiaryPalette       palettes.TonalPalette
	NeutralPalette        palettes.TonalPalette
	NeutralVariantPalette palettes.TonalPalette
	ErrorPalette          palettes.TonalPalette
}

func NewDynamicScheme(
	sourceColorHct *cam.Hct,
	variant Variant,
	contrastLevel float64,
	isDark bool,
	primaryPalette palettes.TonalPalette,
	secondaryPalette palettes.TonalPalette,
	tertiaryPalette palettes.TonalPalette,
	neutralPalette palettes.TonalPalette,
	neutralVariantPalette palettes.TonalPalette,
	errorPalette *palettes.TonalPalette,
) *DynamicScheme {
	scheme := &DynamicScheme{
		SourceColorHct:        sourceColorHct,
		Variant:               variant,
		ContrastLevel:         contrastLevel,
		IsDark:                isDark,
		PrimaryPalette:        primaryPalette,
		SecondaryPalette:      secondaryPalette,
		TertiaryPalette:       tertiaryPalette,
		NeutralPalette:        neutralPalette,
		NeutralVariantPalette: neutralVariantPalette,
	}
	if errorPalette != nil {
		scheme.ErrorPalette = *errorPalette
	} else {
		scheme.ErrorPalette = *palettes.NewTonalPaletteFromHueChroma(25.0, 84.0)
	}
	return scheme
}

func GetRotatedHue(sourceColor *cam.Hct, hues []float64, rotations []float64) float64 {
	sourceHue := sourceColor.Hue()
	if len(rotations) == 1 {
		return utils.SanitizeDegreesDouble(sourceHue + rotations[0])
	}
	size := len(hues)
	for i := 0; i <= size-2; i++ {
		thisHue := hues[i]
		nextHue := hues[i+1]
		if thisHue < sourceHue && sourceHue < nextHue {
			return utils.SanitizeDegreesDouble(sourceHue + rotations[i])
		}
	}
	return sourceHue
}

func (s *DynamicScheme) SourceColorArgb() utils.Argb {
	return s.SourceColorHct.ToInt()
}

func (s *DynamicScheme) GetPrimaryPaletteKeyColor() utils.Argb {
	return MaterialDynamicColors{}.PrimaryPaletteKeyColor().GetArgb(s)
}

func (s *DynamicScheme) GetSecondaryPaletteKeyColor() utils.Argb {
	return MaterialDynamicColors{}.SecondaryPaletteKeyColor().GetArgb(s)
}

func (s *DynamicScheme) GetTertiaryPaletteKeyColor() utils.Argb {
	return MaterialDynamicColors{}.TertiaryPaletteKeyColor().GetArgb(s)
}

func (s *DynamicScheme) GetNeutralPaletteKeyColor() utils.Argb {
	return MaterialDynamicColors{}.NeutralPaletteKeyColor().GetArgb(s)
}

func (s *DynamicScheme) GetNeutralVariantPaletteKeyColor() utils.Argb {
	return MaterialDynamicColors{}.NeutralVariantPaletteKeyColor().GetArgb(s)
}

func (s *DynamicScheme) GetBackground() utils.Argb {
	return MaterialDynamicColors{}.Background().GetArgb(s)
}

func (s *DynamicScheme) GetOnBackground() utils.Argb {
	return MaterialDynamicColors{}.OnBackground().GetArgb(s)
}

func (s *DynamicScheme) GetSurface() utils.Argb {
	return MaterialDynamicColors{}.Surface().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceDim() utils.Argb {
	return MaterialDynamicColors{}.SurfaceDim().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceBright() utils.Argb {
	return MaterialDynamicColors{}.SurfaceBright().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceContainerLowest() utils.Argb {
	return MaterialDynamicColors{}.SurfaceContainerLowest().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceContainerLow() utils.Argb {
	return MaterialDynamicColors{}.SurfaceContainerLow().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceContainer() utils.Argb {
	return MaterialDynamicColors{}.SurfaceContainer().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceContainerHigh() utils.Argb {
	return MaterialDynamicColors{}.SurfaceContainerHigh().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceContainerHighest() utils.Argb {
	return MaterialDynamicColors{}.SurfaceContainerHighest().GetArgb(s)
}

func (s *DynamicScheme) GetOnSurface() utils.Argb {
	return MaterialDynamicColors{}.OnSurface().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceVariant() utils.Argb {
	return MaterialDynamicColors{}.SurfaceVariant().GetArgb(s)
}

func (s *DynamicScheme) GetOnSurfaceVariant() utils.Argb {
	return MaterialDynamicColors{}.OnSurfaceVariant().GetArgb(s)
}

func (s *DynamicScheme) GetInverseSurface() utils.Argb {
	return MaterialDynamicColors{}.InverseSurface().GetArgb(s)
}

func (s *DynamicScheme) GetInverseOnSurface() utils.Argb {
	return MaterialDynamicColors{}.InverseOnSurface().GetArgb(s)
}

func (s *DynamicScheme) GetOutline() utils.Argb {
	return MaterialDynamicColors{}.Outline().GetArgb(s)
}

func (s *DynamicScheme) GetOutlineVariant() utils.Argb {
	return MaterialDynamicColors{}.OutlineVariant().GetArgb(s)
}

func (s *DynamicScheme) GetShadow() utils.Argb {
	return MaterialDynamicColors{}.Shadow().GetArgb(s)
}

func (s *DynamicScheme) GetScrim() utils.Argb {
	return MaterialDynamicColors{}.Scrim().GetArgb(s)
}

func (s *DynamicScheme) GetSurfaceTint() utils.Argb {
	return MaterialDynamicColors{}.SurfaceTint().GetArgb(s)
}

func (s *DynamicScheme) GetPrimary() utils.Argb {
	return MaterialDynamicColors{}.Primary().GetArgb(s)
}

func (s *DynamicScheme) GetOnPrimary() utils.Argb {
	return MaterialDynamicColors{}.OnPrimary().GetArgb(s)
}

func (s *DynamicScheme) GetPrimaryContainer() utils.Argb {
	return MaterialDynamicColors{}.PrimaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetOnPrimaryContainer() utils.Argb {
	return MaterialDynamicColors{}.OnPrimaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetInversePrimary() utils.Argb {
	return MaterialDynamicColors{}.InversePrimary().GetArgb(s)
}

func (s *DynamicScheme) GetSecondary() utils.Argb {
	return MaterialDynamicColors{}.Secondary().GetArgb(s)
}

func (s *DynamicScheme) GetOnSecondary() utils.Argb {
	return MaterialDynamicColors{}.OnSecondary().GetArgb(s)
}

func (s *DynamicScheme) GetSecondaryContainer() utils.Argb {
	return MaterialDynamicColors{}.SecondaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetOnSecondaryContainer() utils.Argb {
	return MaterialDynamicColors{}.OnSecondaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetTertiary() utils.Argb {
	return MaterialDynamicColors{}.Tertiary().GetArgb(s)
}

func (s *DynamicScheme) GetOnTertiary() utils.Argb {
	return MaterialDynamicColors{}.OnTertiary().GetArgb(s)
}

func (s *DynamicScheme) GetTertiaryContainer() utils.Argb {
	return MaterialDynamicColors{}.TertiaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetOnTertiaryContainer() utils.Argb {
	return MaterialDynamicColors{}.OnTertiaryContainer().GetArgb(s)
}

func (s *DynamicScheme) GetError() utils.Argb {
	return MaterialDynamicColors{}.Error().GetArgb(s)
}

func (s *DynamicScheme) GetOnError() utils.Argb {
	return MaterialDynamicColors{}.OnError().GetArgb(s)
}

func (s *DynamicScheme) GetErrorContainer() utils.Argb {
	return MaterialDynamicColors{}.ErrorContainer().GetArgb(s)
}

func (s *DynamicScheme) GetOnErrorContainer() utils.Argb {
	return MaterialDynamicColors{}.OnErrorContainer().GetArgb(s)
}

func (s *DynamicScheme) GetPrimaryFixed() utils.Argb {
	return MaterialDynamicColors{}.PrimaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetPrimaryFixedDim() utils.Argb {
	return MaterialDynamicColors{}.PrimaryFixedDim().GetArgb(s)
}

func (s *DynamicScheme) GetOnPrimaryFixed() utils.Argb {
	return MaterialDynamicColors{}.OnPrimaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetOnPrimaryFixedVariant() utils.Argb {
	return MaterialDynamicColors{}.OnPrimaryFixedVariant().GetArgb(s)
}

func (s *DynamicScheme) GetSecondaryFixed() utils.Argb {
	return MaterialDynamicColors{}.SecondaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetSecondaryFixedDim() utils.Argb {
	return MaterialDynamicColors{}.SecondaryFixedDim().GetArgb(s)
}

func (s *DynamicScheme) GetOnSecondaryFixed() utils.Argb {
	return MaterialDynamicColors{}.OnSecondaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetOnSecondaryFixedVariant() utils.Argb {
	return MaterialDynamicColors{}.OnSecondaryFixedVariant().GetArgb(s)
}

func (s *DynamicScheme) GetTertiaryFixed() utils.Argb {
	return MaterialDynamicColors{}.TertiaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetTertiaryFixedDim() utils.Argb {
	return MaterialDynamicColors{}.TertiaryFixedDim().GetArgb(s)
}

func (s *DynamicScheme) GetOnTertiaryFixed() utils.Argb {
	return MaterialDynamicColors{}.OnTertiaryFixed().GetArgb(s)
}

func (s *DynamicScheme) GetOnTertiaryFixedVariant() utils.Argb {
	return MaterialDynamicColors{}.OnTertiaryFixedVariant().GetArgb(s)
}
