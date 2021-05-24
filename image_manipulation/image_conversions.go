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
func ConvertToAsciiPixels(img image.Image, dimensions []int) ([][]AsciiPixel, error) {

	var terminalWidth, terminalHeight int

	var smallImg image.Image

	// Get dimensions of current terminal
	if len(dimensions) == 0 {

		terminalWidth, _ = consolesize.GetConsoleSize()

		// Sometimes full length outputs print empty lines between ascii art
		terminalWidth -= 1

		// Passing 0 in place of height keeps the original image's aspect ratio
		smallImg = resize.Resize(uint(terminalWidth), 0, img, resize.Lanczos3)
		terminalHeight = smallImg.Bounds().Max.Y - smallImg.Bounds().Min.Y

		// To fix height ratio in eventual ascii art
		terminalHeight = int(0.5 * float32(terminalHeight))

		smallImg = resize.Resize(uint(terminalWidth), uint(terminalHeight), img, resize.Lanczos3)

	} else {
		terminalWidth = dimensions[0]
		terminalHeight = dimensions[1]
		smallImg = resize.Resize(uint(terminalWidth), uint(terminalHeight), img, resize.Lanczos3)
	}

	// If there are passed dimensions, check whether the width exceeds terminal width
	if len(dimensions) > 0 {
		defaultTermWidth, _ := consolesize.GetConsoleSize()
		defaultTermWidth -= 1
		if dimensions[0] > defaultTermWidth {
			return nil, fmt.Errorf("Set width is larger than terminal width")
		}
	}

	// Initialize imgSet 2D slice
	imgSet := make([][]AsciiPixel, terminalHeight)
	for i := range imgSet {
		imgSet[i] = make([]AsciiPixel, terminalWidth)
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

	return imgSet, nil
}
