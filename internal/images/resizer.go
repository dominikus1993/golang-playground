package images

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"sync"

	"github.com/disintegration/imaging"
)

type Size struct {
	Width  int
	Height int
}

func resizeAndSave(ctx context.Context, file *[]byte, size Size, wg *sync.WaitGroup) {
	defer wg.Done()
	output, _ := os.Create(fmt.Sprintf("jp2137_resized_%d_%d.png", size.Height, size.Width))
	defer output.Close()
	reader := bytes.NewReader(*file)

	image, err := jpeg.Decode(reader)
	// check err

	if err != nil {
		panic(err)
	}
	newImage := imaging.Resize(image, size.Width, size.Height, imaging.Lanczos)

	// Encode uses a Writer, use a Buffer if you need the raw []byte
	err = png.Encode(output, newImage)
	// check err
	if err != nil {
		panic(err)
	}
}

func clean(sizes []Size) {
	for _, size := range sizes {
		os.Remove(fmt.Sprintf("jp2137_resized_%d_%d.png", size.Height, size.Width))
	}
}

func ResizeImage(ctx context.Context, file []byte, sizes []Size) {
	var wg sync.WaitGroup
	for _, v := range sizes {
		wg.Add(1)
		go resizeAndSave(ctx, &file, v, &wg)
	}
	wg.Wait()
}
