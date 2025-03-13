package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice of size dy x dx
	image := make([][]uint8, dy)

	// Loop through each row
	for y := range image {
		// Allocate a slice of uint8 for each row
		image[y] = make([]uint8, dx)

		// Loop through each column in the row
		for x := range image[y] {
			// Calculate the pixel value using an interesting function
			// Here, we use (x + y) / 2 as an example
			image[y][x] = uint8((x + y) / 2)

			// You can experiment with other functions like:
			// image[y][x] = uint8(x * y)
			// image[y][x] = uint8(x ^ y)
		}
	}

	return image
}

func main() {
	// Display the image using the pic.Show function
	pic.Show(Pic)
}
