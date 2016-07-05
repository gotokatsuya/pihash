package pihash

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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

		distance := hash.Compare(img1, img2)
		assert.Equal(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
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

		distance := hash.Compare(img1, img2)
		assert.NotEqual(0, distance, target)
		t.Logf("Distance = %d. %s", distance, target)
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

func getImageData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// TestSimilarHashByHumanFace ...
func TestSimilarHashByHumanFace(t *testing.T) {
	t.Skip("Add tests data, please.")

	assert := assert.New(t)

	tests := []struct {
		name        string
		algorithm   HashAlgorithm
		url1        string
		url2        string
		similar     bool
		maxDistance int
	}{}

	for _, tt := range tests {
		target := fmt.Sprintf("%v", tt.name)
		hash := NewHash()
		hash.Algorithm = tt.algorithm

		data, err := getImageData(tt.url1)
		assert.NoError(err, target)
		img1, err := DecodeImageByFile(bytes.NewBuffer(data))
		assert.NoError(err, target)

		data, err = getImageData(tt.url2)
		assert.NoError(err, target)
		img2, err := DecodeImageByFile(bytes.NewBuffer(data))
		assert.NoError(err, target)

		distance := hash.Compare(img1, img2)

		switch tt.similar {
		case true:
			if tt.maxDistance < distance {
				t.Fatal("Should be similar....", distance, target, tt.similar)
			} else {
				t.Log(distance, target)
			}
		case false:
			if tt.maxDistance >= distance {
				t.Fatal("Should not be similar...", distance, target, tt.similar, tt.url1, tt.url2)
			} else {
				t.Log(distance, target)
			}
		}
	}
}
