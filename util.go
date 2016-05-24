package pihash

import (
	"image"
	"sort"

	"github.com/disintegration/gift"
)

func ResizeImage(src image.Image, width, height int) image.Image {
	gi := gift.New(gift.Resize(width, height, gift.LanczosResampling))
	dst := image.NewRGBA(gi.Bounds(src.Bounds()))
	gi.Draw(dst, src)
	return dst
}

type UInt8Slice []uint8

func (p UInt8Slice) Len() int           { return len(p) }
func (p UInt8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p UInt8Slice) Sort() { sort.Sort(p) }
