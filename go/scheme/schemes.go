package scheme

import (
	"math"
	"material-color-utilities/cam"
	"material-color-utilities/dynamiccolor"
	"material-color-utilities/palettes"
	"material-color-utilities/utils"
)

func NewSchemeTonalSpot(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantTonalSpot,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 36.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 16.0),
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()+60.0), 24.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 6.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 8.0),
		nil,
	)
}

func NewSchemeVibrant(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantVibrant,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 200.0),
		*palettes.NewTonalPaletteFromHueChroma(dynamiccolor.GetRotatedHue(sourceColorHct, []float64{0, 41, 61, 101, 131, 181, 251, 301, 360}, []float64{18, 15, 10, 12, 15, 18, 15, 12, 12}), 24.0),
		*palettes.NewTonalPaletteFromHueChroma(dynamiccolor.GetRotatedHue(sourceColorHct, []float64{0, 41, 61, 101, 131, 181, 251, 301, 360}, []float64{35, 30, 20, 25, 30, 35, 30, 25, 25}), 32.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 10.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 12.0),
		nil,
	)
}

func NewSchemeExpressive(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantExpressive,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()+240.0), 40.0),
		*palettes.NewTonalPaletteFromHueChroma(dynamiccolor.GetRotatedHue(sourceColorHct, []float64{0, 21, 51, 121, 151, 191, 271, 321, 360}, []float64{15, 15, 10, 12, 15, 18, 15, 12, 15}), 24.0),
		*palettes.NewTonalPaletteFromHueChroma(dynamiccolor.GetRotatedHue(sourceColorHct, []float64{0, 21, 51, 121, 151, 191, 271, 321, 360}, []float64{35, 30, 20, 25, 30, 35, 30, 25, 35}), 32.0),
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()+15.0), 8.0),
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()+15.0), 12.0),
		nil,
	)
}

func NewSchemeContent(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantContent,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), math.Max(sourceColorHct.Chroma()-32.0, sourceColorHct.Chroma()*0.5)),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()/8.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()/8.0+4.0),
		nil,
	)
}

func NewSchemeMonochrome(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantMonochrome,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		nil,
	)
}

func NewSchemeNeutral(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantNeutral,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 12.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 8.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 16.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 2.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 2.0),
		nil,
	)
}

func NewSchemeFidelity(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantFidelity,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), math.Max(sourceColorHct.Chroma()-32.0, sourceColorHct.Chroma()*0.5)),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()/8.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), sourceColorHct.Chroma()/8.0+4.0),
		nil,
	)
}

func NewSchemeRainbow(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantRainbow,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 16.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 16.0),
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()+60.0), 24.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 0.0),
		nil,
	)
}

func NewSchemeFruitSalad(sourceColorHct *cam.Hct, isDark bool, contrastLevel float64) *dynamiccolor.DynamicScheme {
	return dynamiccolor.NewDynamicScheme(
		sourceColorHct,
		dynamiccolor.VariantFruitSalad,
		contrastLevel,
		isDark,
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()-50.0), 48.0),
		*palettes.NewTonalPaletteFromHueChroma(utils.SanitizeDegreesDouble(sourceColorHct.Hue()-50.0), 36.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 36.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 10.0),
		*palettes.NewTonalPaletteFromHueChroma(sourceColorHct.Hue(), 12.0),
		nil,
	)
}
