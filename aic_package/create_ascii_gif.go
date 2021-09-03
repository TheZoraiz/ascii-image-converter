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

package aic_package

import (
	"image"
	"image/color"

	_ "embed"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
	"github.com/golang/freetype/truetype"

	"github.com/fogleman/gg"
)

/*
Unlike createImageToSave(), this function is optimized to maintain original image dimensions and shrink ascii
art font size to match it. This allows for greater execution speed, which is necessary since a gif contains
multiple images that need to be converted to ascii art, and the potential loss of ascii art quality (since
large ascii art instances will shrink the font too much).

Furthermore, maintaining original gif's width and height also allows for gifs of smaller size.
*/
func createGifFrameToSave(asciiArt [][]imgManip.AsciiChar, img image.Image, colored bool) (image.Image, error) {

	// Original image dimensions
	x := img.Bounds().Dx()
	y := img.Bounds().Dy()

	// Ascii art dimensions
	asciiWidth := len(asciiArt[0])
	asciiHeight := len(asciiArt)

	// Iterators to move pointer on the image to be made
	var xIter float64
	var yIter float64

	var fontSize float64

	// Conditions to alter resulting ascii gif dimensions according to ascii art dimensions
	if asciiWidth > asciiHeight*2 {
		yIter = float64(y) / float64(asciiHeight)

		xIter = yIter / 2
		x = int(xIter * float64(asciiWidth))

		fontSize = xIter

	} else {
		xIter = float64(x) / float64(asciiWidth)

		yIter = xIter * 2
		y = int(yIter * float64(asciiHeight))

		fontSize = xIter
	}

	// 10 extra pixels on both x and y-axis to have 5 pixels of padding on each side
	x += 10
	y += 10

	tempImg := image.NewRGBA(image.Rect(0, 0, x, y))

	dc := gg.NewContext(x, y)

	// Set image background
	dc.SetRGB(
		float64(saveBgColor[0])/255,
		float64(saveBgColor[1])/255,
		float64(saveBgColor[2])/255,
	)
	dc.Clear()

	dc.DrawImage(tempImg, 0, 0)

	// Font size increased during assignment to become more visible. This will not affect image drawing
	fontFace := truetype.NewFace(tempFont, &truetype.Options{Size: fontSize * 1.5})

	dc.SetFontFace(fontFace)

	// Font color of text on picture is white by default
	dc.SetColor(color.White)

	// Pointer to track y-axis on the image frame
	yImgPointer := 5.0

	// These nested loops print each character in asciArt 2D slice separately
	// so that their RGB colors can be maintained in the resulting image
	for _, line := range asciiArt {

		// Pointer to track x-axis on the image frame
		xImgPointer := 5.0

		for _, char := range line {

			if colored {
				// dc.SetColor() sets color for EACH character before printing it
				r := uint8(char.RgbValue[0])
				g := uint8(char.RgbValue[1])
				b := uint8(char.RgbValue[2])
				dc.SetColor(color.RGBA{r, g, b, 255})

			} else {
				r := uint8(fontColor[0])
				g := uint8(fontColor[1])
				b := uint8(fontColor[2])
				dc.SetColor(color.RGBA{r, g, b, 255})
			}

			dc.DrawStringWrapped(char.Simple, xImgPointer, yImgPointer, 0, 0, float64(x), 1.8, gg.AlignLeft)

			// Incremet x-axis pointer character so new one can be printed after it
			// Set to the same constant as in line
			xImgPointer += xIter
		}

		dc.DrawStringWrapped("\n", xImgPointer, yImgPointer, 0, 0, float64(x), 1.8, gg.AlignLeft)

		// Incremet pointer for y axis after every line printed, so
		// new line can start at below the previous one on next iteration
		yImgPointer += yIter
	}

	return dc.Image(), nil
}
