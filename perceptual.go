package pihash

import (
	"image"
	"image/color"
	"math"
	"sort"
)

type Perceptual struct {
	Size int
}

func NewPerceptual() *Perceptual {
	return &Perceptual{
		Size: 64,
	}
}

func (p *Perceptual) Hash(src image.Image) uint64 {
	src = GetResizedGrayscaledImage(src, p.Size, p.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X

	var (
		row  = make([]uint64, maxX)
		rows [][]uint64
	)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			y, _, _ := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			row[j] = uint64(y)
		}
		rows = append(rows, discreteCosineTransformation(row))
	}

	var (
		matrix [][]uint64
		col    = make([]uint64, maxY)
	)
	for j := 0; j < maxX; j++ {
		for i := 0; i < maxY; i++ {
			col[i] = rows[i][j]
		}
		matrix = append(matrix, discreteCosineTransformation(col))
	}

	var pixels []uint64
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			pixels = append(pixels, matrix[i][j])
		}
	}

	median := getMedian(pixels)
	var (
		hash uint64
		one  uint64 = 1
	)
	for _, pixel := range pixels {
		if pixel > median {
			hash |= one
		}
		one = one << 1
	}
	return hash
}

func discreteCosineTransformation(pixels []uint64) []uint64 {
	var (
		size        = len(pixels)
		transformed = make([]uint64, size)
	)
	for i := 0; i < size; i++ {
		var sum float64
		for j := 0; j < size; j++ {
			x := (float64(i) * math.Pi * (float64(j) + 0.5) / float64(size))
			sum += float64(pixels[j]) * math.Cos(x)
		}
		if sum != 0 {
			sum *= math.Sqrt(float64(2) / float64(size))
		}
		if i == 0 {
			sum *= (float64(1) / math.Sqrt(float64(2)))
		}
		transformed[i] = uint64(sum)
	}
	return transformed
}

func getMedian(pixels []uint64) (median uint64) {
	sort.Sort(UInt64Slice(pixels))
	middle := uint64(math.Floor(float64(len(pixels)) / float64(2)))
	if len(pixels)%2 == 0 {
		median = pixels[middle]
	} else {
		low := pixels[middle]
		high := pixels[middle+1]
		median = (low + high) / 2
	}
	return median
}
