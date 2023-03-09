package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dominikus1993/golang-playground/internal/images"
)

func main() {
	f, err := os.ReadFile("./jp2137.jpg")

	if err != nil {
		panic(err)
	}
	sizes := []images.Size{{Width: 71, Height: 55}, {Width: 190, Height: 338}, {Width: 350, Height: 360}, {Width: 720, Height: 1280}}
	for i := 0; i < 100; i++ {
		fmt.Println(fmt.Sprintf("Start %d", i))
		images.ResizeImage(context.Background(), f, sizes)
		fmt.Println(fmt.Sprintf("End %d", i))
	}
}
