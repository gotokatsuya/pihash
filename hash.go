package pihash

import (
	"image"
)

type Hash struct {
	algorithm HashAlgorithm
}

func NewHash() *Hash {
	return &Hash{
		algorithm: NewDifference(),
	}
}

func (h *Hash) Do(src image.Image) uint8 {
	return h.algorithm.Hash(src)
}

func (h *Hash) Compare(src1, src2 image.Image) int {
	hash1 := h.algorithm.Hash(src1)
	hash2 := h.algorithm.Hash(src2)
	return h.Distance(hash1, hash2)
}

func (h *Hash) Distance(hash1, hash2 uint8) int {
	distance := 0
	var i, k uint8
	for i = 0; i < 64; i++ {
		k = (1 << i)
		if (hash1 & k) != (hash2 & k) {
			distance++
		}
	}
	return distance
}
