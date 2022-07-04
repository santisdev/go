package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < len(image); i++ {
		image[i] = make([]uint8, dx)
		for j := 0; j < len(image[i]); j++ {
			image[i][j] = uint8(i * j)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}

// why you have to create the inner slice with make
// https://stackoverflow.com/questions/7703251/slice-of-slices-types
