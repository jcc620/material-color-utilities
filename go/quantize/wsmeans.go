package quantize

import (
	"math"
	"math/rand"
	"sort"
	"github.com/jcc620/material-color-utilities/go/utils"
)

const (
	maxIterations = 100
	minDeltaE     = 3.0
)

type QuantizerResult struct {
	ColorToCount             map[utils.Argb]uint32
	InputPixelToClusterPixel map[utils.Argb]utils.Argb
}

type swatch struct {
	argb       utils.Argb
	population int
}

type distanceToIndex struct {
	distance float64
	index    int
}

func QuantizeWsmeans(inputPixels []utils.Argb, startingClusters []utils.Argb, maxColorsCount int) QuantizerResult {
	if maxColorsCount == 0 || len(inputPixels) == 0 {
		return QuantizerResult{
			ColorToCount:             make(map[utils.Argb]uint32),
			InputPixelToClusterPixel: make(map[utils.Argb]utils.Argb),
		}
	}

	if maxColorsCount > 256 {
		maxColorsCount = 256
	}

	pixelToCount := make(map[utils.Argb]int)
	var pixels []utils.Argb
	var points []Lab
	for _, pixel := range inputPixels {
		if count, ok := pixelToCount[pixel]; ok {
			pixelToCount[pixel] = count + 1
		} else {
			pixels = append(pixels, pixel)
			points = append(points, LabFromInt(pixel))
			pixelToCount[pixel] = 1
		}
	}

	clusterCount := len(points)
	if maxColorsCount < clusterCount {
		clusterCount = maxColorsCount
	}

	if len(startingClusters) > 0 {
		if len(startingClusters) < clusterCount {
			clusterCount = len(startingClusters)
		}
	}

	var clusters []Lab
	for _, argb := range startingClusters {
		clusters = append(clusters, LabFromInt(argb))
	}

	r := rand.New(rand.NewSource(42688))
	additionalClustersNeeded := clusterCount - len(clusters)
	if len(startingClusters) == 0 && additionalClustersNeeded > 0 {
		for i := 0; i < additionalClustersNeeded; i++ {
			l := r.Float64() * 100.0
			a := r.Float64()*200.0 - 100.0
			b := r.Float64()*200.0 - 100.0
			clusters = append(clusters, Lab{L: l, A: a, B: b})
		}
	}

	clusterIndices := make([]int, len(points))
	for i := range points {
		clusterIndices[i] = r.Intn(clusterCount)
	}

	distanceToIndexMatrix := make([][]distanceToIndex, clusterCount)
	for i := range distanceToIndexMatrix {
		distanceToIndexMatrix[i] = make([]distanceToIndex, clusterCount)
	}

	indexMatrix := make([][]int, clusterCount)
	for i := range indexMatrix {
		indexMatrix[i] = make([]int, clusterCount)
	}

	pixelCountSums := make([]int, 256)

	for iteration := 0; iteration < maxIterations; iteration++ {
		for i := 0; i < clusterCount; i++ {
			distanceToIndexMatrix[i][i].distance = 0
			distanceToIndexMatrix[i][i].index = i
			for j := i + 1; j < clusterCount; j++ {
				dist := clusters[i].DeltaE(clusters[j])
				distanceToIndexMatrix[j][i].distance = dist
				distanceToIndexMatrix[j][i].index = i
				distanceToIndexMatrix[i][j].distance = dist
				distanceToIndexMatrix[i][j].index = j
			}

			row := make([]distanceToIndex, clusterCount)
			copy(row, distanceToIndexMatrix[i])
			sort.Slice(row, func(a, b int) bool {
				return row[a].distance < row[b].distance
			})

			for j := 0; j < clusterCount; j++ {
				indexMatrix[i][j] = row[j].index
			}
		}

		colorMoved := false
		for i := 0; i < len(points); i++ {
			point := points[i]
			previousClusterIndex := clusterIndices[i]
			previousCluster := clusters[previousClusterIndex]
			previousDistance := point.DeltaE(previousCluster)
			minimumDistance := previousDistance
			newClusterIndex := -1

			for j := 0; j < clusterCount; j++ {
				if distanceToIndexMatrix[previousClusterIndex][j].distance >= 4*previousDistance {
					continue
				}
				dist := point.DeltaE(clusters[j])
				if dist < minimumDistance {
					minimumDistance = dist
					newClusterIndex = j
				}
			}

			if newClusterIndex != -1 {
				distanceChange := math.Abs(math.Sqrt(minimumDistance) - math.Sqrt(previousDistance))
				if distanceChange > minDeltaE {
					colorMoved = true
					clusterIndices[i] = newClusterIndex
				}
			}
		}

		if !colorMoved && iteration != 0 {
			break
		}

		componentASums := make([]float64, 256)
		componentBSums := make([]float64, 256)
		componentCSums := make([]float64, 256)
		for i := 0; i < clusterCount; i++ {
			pixelCountSums[i] = 0
		}

		for i := 0; i < len(points); i++ {
			clusterIndex := clusterIndices[i]
			point := points[i]
			count := pixelToCount[pixels[i]]

			pixelCountSums[clusterIndex] += count
			componentASums[clusterIndex] += point.L * float64(count)
			componentBSums[clusterIndex] += point.A * float64(count)
			componentCSums[clusterIndex] += point.B * float64(count)
		}

		for i := 0; i < clusterCount; i++ {
			count := pixelCountSums[i]
			if count == 0 {
				clusters[i] = Lab{L: 0, A: 0, B: 0}
				continue
			}
			clusters[i] = Lab{
				L: componentASums[i] / float64(count),
				A: componentBSums[i] / float64(count),
				B: componentCSums[i] / float64(count),
			}
		}
	}

	var swatches []swatch
	var allClusterArgbs []utils.Argb
	for i := 0; i < clusterCount; i++ {
		possibleNewCluster := IntFromLab(clusters[i])
		allClusterArgbs = append(allClusterArgbs, possibleNewCluster)

		count := pixelCountSums[i]
		if count == 0 {
			continue
		}

		useNewCluster := true
		for j := range swatches {
			if swatches[j].argb == possibleNewCluster {
				swatches[j].population += count
				useNewCluster = false
				break
			}
		}

		if useNewCluster {
			swatches = append(swatches, swatch{argb: possibleNewCluster, population: count})
		}
	}

	sort.Slice(swatches, func(i, j int) bool {
		return swatches[i].population > swatches[j].population
	})

	colorToCount := make(map[utils.Argb]uint32)
	for _, s := range swatches {
		colorToCount[s.argb] = uint32(s.population)
	}

	inputPixelToClusterPixel := make(map[utils.Argb]utils.Argb)
	for i := range points {
		pixel := pixels[i]
		clusterIndex := clusterIndices[i]
		clusterArgb := allClusterArgbs[clusterIndex]
		inputPixelToClusterPixel[pixel] = clusterArgb
	}

	return QuantizerResult{
		ColorToCount:             colorToCount,
		InputPixelToClusterPixel: inputPixelToClusterPixel,
	}
}
