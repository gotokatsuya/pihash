package pihash

import (
	"image"
	"math"
)

type Difference struct {
	Size int
}

func NewDifference() *Difference {
	return &Difference{
		Size: 8,
	}
}

func (d *Difference) Hash(src image.Image) uint8 {
	src = ResizeImage(src, d.Size+1, d.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X
	var (
		hash uint8
		one  uint8 = 1
	)
	for i := 0; i < maxY; i++ {
		lr, lg, lb, _ := src.At(0, i).RGBA()
		left := uint8(math.Floor(float64((lr + lg + lb) / 3)))
		for j := 1; j < maxX; j++ {
			rr, rg, rb, _ := src.At(j, i).RGBA()
			right := uint8(math.Floor(float64((rr + rg + rb) / 3)))
			if left > right {
				hash |= one
			}
			left = right
			one = one << 1
		}
	}
	return hash
}
