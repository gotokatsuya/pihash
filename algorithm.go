package pihash

import (
	"image"
)

type HashAlgorithm interface {
	Hash(image.Image) uint64
}
