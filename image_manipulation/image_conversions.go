/*
Copyright © 2021 Zoraiz Hassan <hzoraiz8@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image_conversions

import (
	"fmt"
	"image"
	"image/color"

	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/nfnt/resize"
)

type AsciiPixel struct {
	grayscaleValue uint32
	rgbValue       []uint32
}

// This function shrinks the passed image according to passed dimensions or terminal
// size if none are passed. Stores each pixel's grayscale and RGB values in an AsciiPixel
// instance to simplify getting numeric data for ASCII character comparison.
//
// The returned 2D AsciiPixel slice contains each corresponding pixel's values. Grayscale value
// ranges from 0 to 65535, while RGB values are separate.
func ConvertToAsciiPixels(img image.Image, dimensions []int, flipX, flipY bool) ([][]AsciiPixel, int, int, error) {

	var asciiWidth, asciiHeight int
	var smallImg image.Image

	if len(dimensions) == 0 {

		// Following code in this condition calculates ratio according to terminal height

		terminalWidth, terminalHeight := consolesize.GetConsoleSize()
		asciiHeight = terminalHeight - 1

		// Passing 0 in place of width keeps the original image's aspect ratio
		smallImg = resize.Resize(0, uint(asciiHeight), img, resize.Lanczos3)
		asciiWidth = smallImg.Bounds().Max.X - smallImg.Bounds().Min.X

		// To fix aspect ratio in eventual ascii art
		asciiWidth = int(2 * float32(asciiWidth))

		// If ascii width exceeds terminal width, change ratio with respect to terminal width
		if asciiWidth > terminalWidth {
			smallImg = resize.Resize(uint(terminalWidth), 0, img, resize.Lanczos3)

			asciiWidth = terminalWidth - 1
			asciiHeight = smallImg.Bounds().Max.Y - smallImg.Bounds().Min.Y

			// To fix aspect ratio in eventual ascii art
			asciiHeight = int(0.5 * float32(asciiHeight))
		}

		smallImg = resize.Resize(uint(asciiWidth), uint(asciiHeight), img, resize.Lanczos3)

	} else {
		asciiWidth = dimensions[0]
		asciiHeight = dimensions[1]
		smallImg = resize.Resize(uint(asciiWidth), uint(asciiHeight), img, resize.Lanczos3)
	}

	// If there are passed dimensions, check whether the width exceeds terminal width
	if len(dimensions) > 0 {
		defaultTermWidth, _ := consolesize.GetConsoleSize()

		defaultTermWidth -= 1
		if dimensions[0] > defaultTermWidth {
			return nil, 0, 0, fmt.Errorf("set width is larger than terminal width")
		}
	}

	// Initialize imgSet 2D slice
	imgSet := make([][]AsciiPixel, asciiHeight)
	for i := range imgSet {
		imgSet[i] = make([]AsciiPixel, asciiWidth)
	}

	b := smallImg.Bounds()

	// These nested loops iterate through each pixel of resized image and get an AsciiPixel instance
	for y := b.Min.Y; y < b.Max.Y; y++ {

		var temp []AsciiPixel
		for x := b.Min.X; x < b.Max.X; x++ {

			oldPixel := smallImg.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)

			// We only need Red from Red, Green, Blue (RGB) for grayscaleValue in AsciiPixel since they have the same value for grayscale images
			r1, _, _, _ := pixel.RGBA()

			// Get colored RGB values of original pixel for rgbValue in AsciiPixel
			r2, g2, b2, _ := oldPixel.RGBA()
			r2 = uint32(r2 / 257)
			g2 = uint32(g2 / 257)
			b2 = uint32(b2 / 257)

			temp = append(temp, AsciiPixel{
				grayscaleValue: r1,
				rgbValue:       []uint32{r2, g2, b2},
			})

		}
		imgSet[y] = temp
	}

	// This rarely affects performance since the ascii art 2D slice size isn't that large
	if flipX || flipY {
		imgSet = reverse(imgSet, flipX, flipY)
	}

	return imgSet, asciiWidth, asciiHeight, nil
}

func reverse(imgSet [][]AsciiPixel, flipX, flipY bool) [][]AsciiPixel {

	if flipX {
		for _, row := range imgSet {
			for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
				row[i], row[j] = row[j], row[i]
			}
		}
	}

	if flipY {
		for i, j := 0, len(imgSet)-1; i < j; i, j = i+1, j-1 {
			imgSet[i], imgSet[j] = imgSet[j], imgSet[i]
		}
	}

	return imgSet
}
