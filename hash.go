package pihash

import (
	"image"
)

type Hash struct {
	Algorithm HashAlgorithm
}

func NewHash() *Hash {
	return &Hash{
		Algorithm: NewDifference(),
	}
}

func (h *Hash) Do(src image.Image) uint64 {
	return h.Algorithm.Hash(src)
}

func (h *Hash) Compare(src1, src2 image.Image) int {
	hash1 := h.Algorithm.Hash(src1)
	hash2 := h.Algorithm.Hash(src2)
	return h.Distance(hash1, hash2)
}

func (h *Hash) Distance(hash1, hash2 uint64) int {
	distance := 0
	var i, k uint64
	for i = 0; i < 64; i++ {
		k = (1 << i)
		if (hash1 & k) != (hash2 & k) {
			distance++
		}
	}
	return distance
}
