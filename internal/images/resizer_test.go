package images

import (
	"context"
	"os"
	"testing"
)

func BenchmarkResizeImage(b *testing.B) {

	f, err := os.ReadFile("../../jp2137.jpg")

	if err != nil {
		panic(err)
	}

	sizes := []Size{{Width: 71, Height: 55}, {Width: 190, Height: 338}, {Width: 350, Height: 360}, {Width: 720, Height: 1280}}
	defer clean(sizes)
	for n := 0; n < b.N; n++ {
		ResizeImage(context.Background(), f, sizes)
	}
}
