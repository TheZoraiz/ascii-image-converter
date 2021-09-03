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
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

//go:embed Hack-Regular.ttf
var embeddedHackRegularFont []byte

//go:embed DejaVuSans-Oblique.ttf
var embeddedDejaVuObliqueFont []byte

var tempFont *truetype.Font

// Load embedded font
func init() {
	tempFont, _ = truetype.Parse(embeddedHackRegularFont)
}

/*
Unlike createGifFrameToSave(), this function is altered to ignore execution time and has a fixed font size.
This creates maximum quality ascii art, although the resulting image will not have the same dimensions
as the original image, but the ascii art quality will be maintained. This is required, since smaller provided
images will considerably decrease ascii art quality because of smaller font size.

Size of resulting image may also be considerably larger than original image.
*/
func createImageToSave(asciiArt [][]imgManip.AsciiChar, colored bool, saveImagePath, imagePath, urlImgName string) error {

	constant := 14.0

	x := len(asciiArt[0])
	y := len(asciiArt)

	// Multipying resulting image dimensions with respect to constant
	x = int(constant * float64(x))

	y = int(constant * float64(y))
	y = y * 2

	// 10 extra pixels on both x and y-axis to have 5 pixels of padding on each side
	y += 10
	x += 10

	tempImg := image.NewRGBA(image.Rect(0, 0, x, y))

	imgWidth := tempImg.Bounds().Dx()
	imgHeight := tempImg.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)

	// Set image background
	dc.SetRGB(
		float64(saveBgColor[0])/255,
		float64(saveBgColor[1])/255,
		float64(saveBgColor[2])/255,
	)
	dc.Clear()

	dc.DrawImage(tempImg, 0, 0)

	fontFace := truetype.NewFace(tempFont, &truetype.Options{Size: constant * 1.5})
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
			xImgPointer += float64(constant)
		}

		dc.DrawStringWrapped("\n", xImgPointer, yImgPointer, 0, 0, float64(x), 1.8, gg.AlignLeft)

		// Incremet pointer for y axis after every line printed, so
		// new line can start at below the previous one on next iteration
		yImgPointer += float64(constant * 2)
	}

	imageName, err := createSaveFileName(imagePath, urlImgName, "-ascii-art.png")
	if err != nil {
		return err
	}

	fullPathName, err := getFullSavePath(imageName, saveImagePath)
	if err != nil {
		return err
	}

	return dc.SavePNG(fullPathName)
}
