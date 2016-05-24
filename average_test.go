package pihash

import "testing"

func TestAverageHash(t *testing.T) {
	average := NewAverage()
	img, err := DecodeImageByPath("testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("img == nil.")
	}
	res := average.Hash(img)
	if res == 0 {
		t.Fatal(res)
	} else {
		t.Log(res)
	}
}
