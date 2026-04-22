package quantize

import (
	"material-color-utilities/utils"
)

func QuantizeCelebi(pixels []utils.Argb, maxColorsCount int) QuantizerResult {
	if maxColorsCount == 0 || len(pixels) == 0 {
		return QuantizerResult{
			ColorToCount:             make(map[utils.Argb]uint32),
			InputPixelToClusterPixel: make(map[utils.Argb]utils.Argb),
		}
	}

	if maxColorsCount > 256 {
		maxColorsCount = 256
	}

	opaquePixels := make([]utils.Argb, 0, len(pixels))
	for _, pixel := range pixels {
		if utils.IsOpaque(pixel) {
			opaquePixels = append(opaquePixels, pixel)
		}
	}

	wuResult := QuantizeWu(opaquePixels, maxColorsCount)
	return QuantizeWsmeans(opaquePixels, wuResult, maxColorsCount)
}
