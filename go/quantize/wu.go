package quantize

import (
	"github.com/jcc620/material-color-utilities/go/utils"
)

type box struct {
	r0, r1, g0, g1, b0, b1, vol int
}

type direction int

const (
	directionRed direction = iota
	directionGreen
	directionBlue
)

const (
	indexBits  = 5
	indexCount = (1 << indexBits) + 1
	totalSize  = indexCount * indexCount * indexCount
	maxColors  = 256
)

func getIndex(r, g, b int) int {
	return (r << (indexBits * 2)) + (r << (indexBits + 1)) + (g << indexBits) + r + g + b
}

func constructHistogram(pixels []utils.Argb, weights, mr, mg, mb []int64, moments []float64) {
	for _, pixel := range pixels {
		red := utils.RedFromInt(pixel)
		green := utils.GreenFromInt(pixel)
		blue := utils.BlueFromInt(pixel)

		bitsToRemove := 8 - indexBits
		indexR := (red >> uint(bitsToRemove)) + 1
		indexG := (green >> uint(bitsToRemove)) + 1
		indexB := (blue >> uint(bitsToRemove)) + 1
		index := getIndex(indexR, indexG, indexB)

		weights[index]++
		mr[index] += int64(red)
		mg[index] += int64(green)
		mb[index] += int64(blue)
		moments[index] += float64(red*red + green*green + blue*blue)
	}
}

func computeMoments(weights, mr, mg, mb []int64, moments []float64) {
	for r := 1; r < indexCount; r++ {
		area := make([]int64, indexCount)
		areaR := make([]int64, indexCount)
		areaG := make([]int64, indexCount)
		areaB := make([]int64, indexCount)
		area2 := make([]float64, indexCount)
		for g := 1; g < indexCount; g++ {
			var line int64
			var lineR int64
			var lineG int64
			var lineB int64
			var line2 float64
			for b := 1; b < indexCount; b++ {
				index := getIndex(r, g, b)
				line += weights[index]
				lineR += mr[index]
				lineG += mg[index]
				lineB += mb[index]
				line2 += moments[index]

				area[b] += line
				areaR[b] += lineR
				areaG[b] += lineG
				areaB[b] += lineB
				area2[b] += line2

				previousIndex := getIndex(r-1, g, b)
				weights[index] = weights[previousIndex] + area[b]
				mr[index] = mr[previousIndex] + areaR[b]
				mg[index] = mg[previousIndex] + areaG[b]
				mb[index] = mb[previousIndex] + areaB[b]
				moments[index] = moments[previousIndex] + area2[b]
			}
		}
	}
}

func top(cube box, dir direction, position int, moment []int64) int64 {
	switch dir {
	case directionRed:
		return moment[getIndex(position, cube.g1, cube.b1)] -
			moment[getIndex(position, cube.g1, cube.b0)] -
			moment[getIndex(position, cube.g0, cube.b1)] +
			moment[getIndex(position, cube.g0, cube.b0)]
	case directionGreen:
		return moment[getIndex(cube.r1, position, cube.b1)] -
			moment[getIndex(cube.r1, position, cube.b0)] -
			moment[getIndex(cube.r0, position, cube.b1)] +
			moment[getIndex(cube.r0, position, cube.b0)]
	case directionBlue:
		return moment[getIndex(cube.r1, cube.g1, position)] -
			moment[getIndex(cube.r1, cube.g0, position)] -
			moment[getIndex(cube.r0, cube.g1, position)] +
			moment[getIndex(cube.r0, cube.g0, position)]
	default:
		return 0
	}
}

func bottom(cube box, dir direction, moment []int64) int64 {
	switch dir {
	case directionRed:
		return -moment[getIndex(cube.r0, cube.g1, cube.b1)] +
			moment[getIndex(cube.r0, cube.g1, cube.b0)] +
			moment[getIndex(cube.r0, cube.g0, cube.b1)] -
			moment[getIndex(cube.r0, cube.g0, cube.b0)]
	case directionGreen:
		return -moment[getIndex(cube.r1, cube.g0, cube.b1)] +
			moment[getIndex(cube.r1, cube.g0, cube.b0)] +
			moment[getIndex(cube.r0, cube.g0, cube.b1)] -
			moment[getIndex(cube.r0, cube.g0, cube.b0)]
	case directionBlue:
		return -moment[getIndex(cube.r1, cube.g1, cube.b0)] +
			moment[getIndex(cube.r1, cube.g0, cube.b0)] +
			moment[getIndex(cube.r0, cube.g1, cube.b0)] -
			moment[getIndex(cube.r0, cube.g0, cube.b0)]
	default:
		return 0
	}
}

func vol(cube box, moment []int64) int64 {
	return moment[getIndex(cube.r1, cube.g1, cube.b1)] -
		moment[getIndex(cube.r1, cube.g1, cube.b0)] -
		moment[getIndex(cube.r1, cube.g0, cube.b1)] +
		moment[getIndex(cube.r1, cube.g0, cube.b0)] -
		moment[getIndex(cube.r0, cube.g1, cube.b1)] +
		moment[getIndex(cube.r0, cube.g1, cube.b0)] +
		moment[getIndex(cube.r0, cube.g0, cube.b1)] -
		moment[getIndex(cube.r0, cube.g0, cube.b0)]
}

func variance(cube box, weights, mr, mg, mb []int64, moments []float64) float64 {
	dr := float64(vol(cube, mr))
	dg := float64(vol(cube, mg))
	db := float64(vol(cube, mb))
	xx := moments[getIndex(cube.r1, cube.g1, cube.b1)] -
		moments[getIndex(cube.r1, cube.g1, cube.b0)] -
		moments[getIndex(cube.r1, cube.g0, cube.b1)] +
		moments[getIndex(cube.r1, cube.g0, cube.b0)] -
		moments[getIndex(cube.r0, cube.g1, cube.b1)] +
		moments[getIndex(cube.r0, cube.g1, cube.b0)] +
		moments[getIndex(cube.r0, cube.g0, cube.b1)] -
		moments[getIndex(cube.r0, cube.g0, cube.b0)]
	hypotenuse := dr*dr + dg*dg + db*db
	volume := float64(vol(cube, weights))
	return xx - hypotenuse/volume
}

func maximize(cube box, dir direction, first, last int, wholeW, wholeR, wholeG, wholeB int64, weights, mr, mg, mb []int64) (float64, int) {
	bottomR := bottom(cube, dir, mr)
	bottomG := bottom(cube, dir, mg)
	bottomB := bottom(cube, dir, mb)
	bottomW := bottom(cube, dir, weights)

	max := 0.0
	cut := -1

	for i := first; i < last; i++ {
		halfR := bottomR + top(cube, dir, i, mr)
		halfG := bottomG + top(cube, dir, i, mg)
		halfB := bottomB + top(cube, dir, i, mb)
		halfW := bottomW + top(cube, dir, i, weights)
		if halfW == 0 {
			continue
		}

		temp := (float64(halfR)*float64(halfR) + float64(halfG)*float64(halfG) + float64(halfB)*float64(halfB)) / float64(halfW)

		halfR = wholeR - halfR
		halfG = wholeG - halfG
		halfB = wholeB - halfB
		halfW = wholeW - halfW
		if halfW == 0 {
			continue
		}
		temp += (float64(halfR)*float64(halfR) + float64(halfG)*float64(halfG) + float64(halfB)*float64(halfB)) / float64(halfW)

		if temp > max {
			max = temp
			cut = i
		}
	}
	return max, cut
}

func cut(box1 *box, box2 *box, weights, mr, mg, mb []int64) bool {
	wholeR := vol(*box1, mr)
	wholeG := vol(*box1, mg)
	wholeB := vol(*box1, mb)
	wholeW := vol(*box1, weights)

	maxR, cutR := maximize(*box1, directionRed, box1.r0+1, box1.r1, wholeW, wholeR, wholeG, wholeB, weights, mr, mg, mb)
	maxG, cutG := maximize(*box1, directionGreen, box1.g0+1, box1.g1, wholeW, wholeR, wholeG, wholeB, weights, mr, mg, mb)
	maxB, cutB := maximize(*box1, directionBlue, box1.b0+1, box1.b1, wholeW, wholeR, wholeG, wholeB, weights, mr, mg, mb)

	var dir direction
	if maxR >= maxG && maxR >= maxB {
		dir = directionRed
		if cutR < 0 {
			return false
		}
	} else if maxG >= maxR && maxG >= maxB {
		dir = directionGreen
	} else {
		dir = directionBlue
	}

	box2.r1 = box1.r1
	box2.g1 = box1.g1
	box2.b1 = box1.b1

	if dir == directionRed {
		box2.r0 = cutR
		box1.r1 = cutR
		box2.g0 = box1.g0
		box2.b0 = box1.b0
	} else if dir == directionGreen {
		box2.r0 = box1.r0
		box2.g0 = cutG
		box1.g1 = cutG
		box2.b0 = box1.b0
	} else {
		box2.r0 = box1.r0
		box2.g0 = box1.g0
		box2.b0 = cutB
		box1.b1 = cutB
	}

	box1.vol = (box1.r1 - box1.r0) * (box1.g1 - box1.g0) * (box1.b1 - box1.b0)
	box2.vol = (box2.r1 - box2.r0) * (box2.g1 - box2.g0) * (box2.b1 - box2.b0)
	return true
}

func QuantizeWu(pixels []utils.Argb, maxColorsCount int) []utils.Argb {
	if maxColorsCount <= 0 || maxColorsCount > 256 || len(pixels) == 0 {
		return nil
	}

	weights := make([]int64, totalSize)
	mr := make([]int64, totalSize)
	mg := make([]int64, totalSize)
	mb := make([]int64, totalSize)
	moments := make([]float64, totalSize)
	constructHistogram(pixels, weights, mr, mg, mb, moments)
	computeMoments(weights, mr, mg, mb, moments)

	cubes := make([]box, maxColors)
	cubes[0].r0, cubes[0].g0, cubes[0].b0 = 0, 0, 0
	cubes[0].r1, cubes[0].g1, cubes[0].b1 = indexCount-1, indexCount-1, indexCount-1

	volumeVariance := make([]float64, maxColors)
	next := 0
	for i := 1; i < maxColorsCount; i++ {
		if cut(&cubes[next], &cubes[i], weights, mr, mg, mb) {
			if cubes[next].vol > 1 {
				volumeVariance[next] = variance(cubes[next], weights, mr, mg, mb, moments)
			} else {
				volumeVariance[next] = 0.0
			}
			if cubes[i].vol > 1 {
				volumeVariance[i] = variance(cubes[i], weights, mr, mg, mb, moments)
			} else {
				volumeVariance[i] = 0.0
			}
		} else {
			volumeVariance[next] = 0.0
			i--
		}

		next = 0
		temp := volumeVariance[0]
		for j := 1; j <= i; j++ {
			if volumeVariance[j] > temp {
				temp = volumeVariance[j]
				next = j
			}
		}
		if temp <= 0.0 {
			maxColorsCount = i + 1
			break
		}
	}

	outColors := make([]utils.Argb, 0, maxColorsCount)
	for i := 0; i < maxColorsCount; i++ {
		weight := vol(cubes[i], weights)
		if weight > 0 {
			red := int(vol(cubes[i], mr) / weight)
			green := int(vol(cubes[i], mg) / weight)
			blue := int(vol(cubes[i], mb) / weight)
			argb := utils.ArgbFromRgb(red, green, blue)
			outColors = append(outColors, argb)
		}
	}

	return outColors
}
