package pihash

import "testing"

func TestPerceptualHash(t *testing.T) {
	perceptual := NewPerceptual()
	img, err := DecodeImageByPath("testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("img == nil.")
	}
	res := perceptual.Hash(img)
	if res == 0 {
		t.Fatal(res)
	} else {
		t.Log(res)
	}
}
