# pihash

Perceptual Image Hash for Go

## usage

```go
hash := NewHash()
hash.Algorithm = NewPerceptual()
// hash.Algorithm = NewAveragel()
// hash.Algorithm = NewDifference()
img1, _ := DecodeImageByPath(imgPath1)
img2, _ := DecodeImageByPath(imgPath2)

const threshold = 5
if distance := hash.Compare(img1, img2); distance <= threshold {
    // same image
} else {
    // not same image
}
```

## test

```bash
go test github.com/gotokatsuya/pihash -v
```
