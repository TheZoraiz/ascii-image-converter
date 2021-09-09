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
	"image"
	"image/color"
)

type AsciiPixel struct {
	charDepth      uint32
	grayscaleValue [3]uint32
	rgbValue       [3]uint32
}

/*
This function shrinks the passed image according to specified or default dimensions.
Stores each pixel's grayscale and RGB values in an AsciiPixel instance to simplify
getting numeric data for ASCII character comparison.

The returned 2D AsciiPixel slice contains each corresponding pixel's values
*/
func ConvertToAsciiPixels(img image.Image, dimensions []int, width, height int, flipX, flipY, full, isBraille, dither, noTermSizeComparison bool) ([][]AsciiPixel, error) {

	var smallImg image.Image
	var err error

	if noTermSizeComparison {
		smallImg, err = resizeImageNoTerm(img, isBraille, dimensions, width, height)
	} else {
		smallImg, err = resizeImage(img, full, isBraille, dimensions, width, height)
	}
	if err != nil {
		return nil, err
	}

	// We mainatin a dithered image literal along with original image
	// The colors are kept from original image
	var ditheredImage image.Image

	if isBraille && dither {
		ditheredImage = ditherImage(smallImg)
	}

	var imgSet [][]AsciiPixel

	b := smallImg.Bounds()

	// These nested loops iterate through each pixel of resized image and get an AsciiPixel instance
	for y := b.Min.Y; y < b.Max.Y; y++ {

		var temp []AsciiPixel
		for x := b.Min.X; x < b.Max.X; x++ {

			oldPixel := smallImg.At(x, y)
			grayPixel := color.GrayModel.Convert(oldPixel)

			r1, g1, b1, _ := grayPixel.RGBA()
			charDepth := r1 / 257 // Only Red is needed from RGB for charDepth in AsciiPixel since they have the same value for grayscale images
			r1 = uint32(r1 / 257)
			g1 = uint32(g1 / 257)
			b1 = uint32(b1 / 257)

			if isBraille && dither {

				// Change charDepth if image dithering is applied
				// 		Note that neither grayscale nor original color values are changed.
				// 		Only charDepth is kept from dithered image. This is because a
				// 		dithered image loses its colors so it's only used to check braille
				// 		dots' visibility

				ditheredGrayPixel := color.GrayModel.Convert(ditheredImage.At(x, y))
				charDepth, _, _, _ = ditheredGrayPixel.RGBA()
				charDepth = charDepth / 257
			}

			// Get co1ored RGB values of original pixel for rgbValue in AsciiPixel
			r2, g2, b2, _ := oldPixel.RGBA()
			r2 = uint32(r2 / 257)
			g2 = uint32(g2 / 257)
			b2 = uint32(b2 / 257)

			temp = append(temp, AsciiPixel{
				charDepth:      charDepth,
				grayscaleValue: [3]uint32{r1, g1, b1},
				rgbValue:       [3]uint32{r2, g2, b2},
			})

		}
		imgSet = append(imgSet, temp)
	}

	// This rarely affects performance since the ascii art 2D slice size isn't that large
	if flipX || flipY {
		imgSet = reverse(imgSet, flipX, flipY)
	}

	return imgSet, nil
}
