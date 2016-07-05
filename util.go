package pihash

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"sort"

	"github.com/disintegration/gift"
)

func GetResizedImage(src image.Image, width, height int) image.Image {
	gi := gift.New(gift.Resize(width, height, gift.LinearResampling))
	dst := image.NewRGBA(gi.Bounds(src.Bounds()))
	gi.Draw(dst, src)
	return dst
}

func GetResizedGrayscaledImage(src image.Image, width, height int) image.Image {
	gi := gift.New(gift.Resize(width, height, gift.LinearResampling), gift.Grayscale())
	dst := image.NewRGBA(gi.Bounds(src.Bounds()))
	gi.Draw(dst, src)
	return dst
}

func DecodeImageByPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func DecodeImageByFile(file io.Reader) (image.Image, error) {
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, nil
}

type UInt64Slice []uint64

func (p UInt64Slice) Len() int           { return len(p) }
func (p UInt64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p UInt64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p UInt64Slice) Sort()              { sort.Sort(p) }
