package main

import (
	"bufio"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	f, err := os.Open("./jp2137.jpg")

	if err != nil {
		panic(err)
	}
	output, _ := os.Create("jp2137_resized.png")
	defer output.Close()

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Bytes()
	image, err := jpeg.Decode(f)
	// check err

	if err != nil {
		panic(err)
	}
	newImage := resize.Resize(160, 0, image, resize.Lanczos3)

	// Encode uses a Writer, use a Buffer if you need the raw []byte
	err = png.Encode(output, newImage)
	// check err
	if err != nil {
		panic(err)
	}
}
