package pihash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameHash(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm HashAlgorithm
		img1      string
		img2      string
	}{
		{"Average", NewAverage(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
		{"Difference", NewDifference(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
		{"Perceptual", NewPerceptual(), "testdata/sample1.jpg", "testdata/sample1.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := DecodeImageByPath(tt.img1)
		assert.NoError(err, target)
		img2, err := DecodeImageByPath(tt.img2)
		assert.NoError(err, target)

		assert.Equal(0, hash.Compare(img1, img2), target)
	}
}

func TestUnSameHash(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		algorithm HashAlgorithm
		img1      string
		img2      string
	}{
		{"Average", NewAverage(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
		{"Difference", NewDifference(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
		{"Perceptual", NewPerceptual(), "testdata/sample1.jpg", "testdata/sample2.jpg"},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)

		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := DecodeImageByPath(tt.img1)
		assert.NoError(err, target)
		img2, err := DecodeImageByPath(tt.img2)
		assert.NoError(err, target)

		assert.NotEqual(0, hash.Compare(img1, img2), target)
	}
}
func TestSimilarHash(t *testing.T) {

	assert := assert.New(t)
	tests := []struct {
		name        string
		algorithm   HashAlgorithm
		img1        string
		img2        string
		maxDistance int
	}{
		{"Average", NewAverage(), "testdata/sample1.jpg", "testdata/sample1-small.jpg", 10},
		{"Difference", NewDifference(), "testdata/sample1.jpg", "testdata/sample1-small.jpg", 10},
		{"Perceptual", NewPerceptual(), "testdata/sample1.jpg", "testdata/sample1-small.jpg", 10},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)
		hash := NewHash()
		hash.Algorithm = tt.algorithm

		img1, err := DecodeImageByPath(tt.img1)
		assert.NoError(err, target)
		img2, err := DecodeImageByPath(tt.img2)
		assert.NoError(err, target)

		distance := hash.Compare(img1, img2)
		if tt.maxDistance < distance {
			t.Fatal("Not similar.", target)
		} else {
			t.Log(distance, target)
		}
	}
}
