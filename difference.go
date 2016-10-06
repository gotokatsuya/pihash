package pihash

import "image"

type Difference struct {
	Size int
}

func NewDifference() *Difference {
	return &Difference{
		Size: 8,
	}
}

func (d *Difference) Hash(src image.Image) uint64 {
	src = GetResizedGrayscaledImage(src, d.Size+1, d.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X
	var (
		hash uint64
		one  uint64 = 1
	)
	for i := 0; i < maxY; i++ {
		left := wrapSumPixels(src.At(0, i).RGBA())
		for j := 1; j < maxX; j++ {
			right := wrapSumPixels(src.At(j, i).RGBA())
			if left > right {
				hash |= one
			}
			left = right
			one = one << 1
		}
	}
	return hash
}
