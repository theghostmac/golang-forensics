package main

import (
	"io"
	"log"
	"os"
)

func main() {
	firstFile, err := os.Open("go-generated-image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	secondFile, err := os.Open("zipper.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	steganImage, err := os.Create("stegan-image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer steganImage.Close()

	_, err = io.Copy(steganImage, firstFile)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(steganImage, secondFile)
	if err != nil {
		log.Fatal(err)
	}
}
