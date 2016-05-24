package pihash

import "testing"

func TestDifferenceHash(t *testing.T) {
	difference := NewDifference()
	img, err := DecodeImageByPath("testdata/sample1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("img == nil.")
	}
	res := difference.Hash(img)
	if res == 0 {
		t.Fatal(res)
	} else {
		t.Log(res)
	}
}
