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

func (p *Perceptual) Hash(src image.Image) uint8 {
	src = ResizeImage(src, p.Size, p.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X

	var (
		row  = make([]uint8, maxX)
		rows = make([][]uint8, maxY*maxX)
	)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			y, _, _ := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			row[j] = y
		}
		rows[i] = discreteCosineTransformation(row)
	}

	var (
		matrix = make([][]uint8, maxY*maxX)
		col    = make([]uint8, maxX)
	)
	for j := 0; j < maxX; j++ {
		for i := 0; i < maxY; i++ {
			col[i] = rows[i][j]
		}
		matrix[j] = discreteCosineTransformation(col)
	}

	var pixels []uint8
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			pixels = append(pixels, matrix[i][j])
		}
	}

	medianValue := median(pixels)

	var (
		hash uint8
		one  uint8 = 1
	)
	for _, pixel := range pixels {
		if pixel > medianValue {
			hash |= one
		}
		one = one << 1
	}
	return hash
}

func discreteCosineTransformation(pixels []uint8) []uint8 {
	var (
		size        = len(pixels)
		transformed = make([]uint8, size)
	)
	for i := 0; i < size; i++ {
		var sum uint8
		for j := 0; j < size; j++ {
			v := (float64(i) * math.Pi * (float64(j) + 0.5) / float64(size))
			sum += pixels[j] * uint8(math.Cos(v))
		}
		sum *= uint8(math.Sqrt(float64(2 / size)))
		if i == 0 {
			sum *= 1 / uint8(math.Sqrt(float64(2)))
		}
		transformed[i] = sum
	}
	return transformed
}

func median(pixels []uint8) (median uint8) {
	sort.Sort(UInt8Slice(pixels))
	middle := uint8(math.Floor(float64(len(pixels) / 2)))
	if len(pixels)%2 == 0 {
		median = pixels[middle]
	} else {
		low := pixels[middle]
		high := pixels[middle+1]
		median = (low + high) / 2
	}
	return median
}
