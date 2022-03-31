package main

import (
	"image"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
)

func main() {
	JPEGImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	for pix := 0; pix < 100*200; pix++ {
		pixelOffset := 4 * pix
		// produce Red color
		JPEGImage.Pix[0+pixelOffset] = uint8(rand.Intn(256))
		// produce Green color
		JPEGImage.Pix[1+pixelOffset] = uint8(rand.Intn(256))
		// produce Blue color
		JPEGImage.Pix[2+pixelOffset] = uint8(rand.Intn(256))
		// produce Alpha color
		JPEGImage.Pix[3+pixelOffset] = 255
	}

	outputImage, err := os.Create("go-generated-image.jpg")
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(outputImage, JPEGImage, nil)
	if err != nil {
		return
	}

	err = outputImage.Close()
	if err != nil {
		log.Fatal(err)
	}
}
