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
	"os"

	_ "embed"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
	"github.com/golang/freetype/truetype"

	"github.com/fogleman/gg"
)

// This file mostly has spaghetti code lol.
// Will work on organizing it into a readable format soon. Apologies...

// To embed font directly into the binary, instead of packaging it as a separate file
//go:embed RobotoMono-Bold.ttf
var embeddedFontFile []byte

func createImageToSave(asciiArt [][]imgManip.AsciiChar, x, y int, colored bool, saveImagePath, imagePath, urlImgName string) error {

	// IMPORTANT NOTE:
	//		The raw numbers from here on are strictly experimental
	//		and are used because they just happened to work with
	//		the tallest and widest images I could find, which cover
	//		the extreme cases of possible images encountered.

	// Multipying with respect to font size. I altered the numbers
	// until extreme cases' output images becaume right
	x = int(12.6 * float32(x))

	y = int(13.5 * float32(y))
	y = y * 2

	// To give small extra margins near the edges
	y += 5
	x += 10

	tempImg := image.NewRGBA(image.Rect(0, 0, x, y))

	imgWidth := tempImg.Bounds().Dx()
	imgHeight := tempImg.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)

	// Set image background as black
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	dc.DrawImage(tempImg, 0, 0)

	// Load embedded font
	tempFont, err := truetype.Parse(embeddedFontFile)
	if err != nil {
		return err
	}
	robotoBoldFontFace := truetype.NewFace(tempFont, &truetype.Options{Size: 21.0})

	dc.SetFontFace(robotoBoldFontFace)

	// Font color of text on picture is white by default
	dc.SetColor(color.White)

	// Pointer to track y-axis on the image frame
	yImgPointer := 2.5

	// These nested loops print each character in asciArt 2D slice separately
	// so that their RGB colors can be maintained in the resulting image
	for _, line := range asciiArt {

		// Pointer to track x-axis on the image frame
		xImgPointer := 5.0

		for _, char := range line {

			r := uint8(char.RgbValue[0])
			g := uint8(char.RgbValue[1])
			b := uint8(char.RgbValue[2])

			if colored {
				// Simple put, dc.SetColor() sets color for EACH character before printing it
				dc.SetColor(color.RGBA{r, g, b, 255})
			}

			dc.DrawStringWrapped(char.Simple, xImgPointer, yImgPointer, 0, 0, float64(x), 1.8, gg.AlignLeft)

			// Incremet x-axis pointer character so new one can be printed after it
			xImgPointer += 12.6
		}

		dc.DrawStringWrapped("\n", xImgPointer, yImgPointer, 0, 0, float64(x), 1.8, gg.AlignLeft)

		// Incremet pointer for y axis after every line printed, so
		// new line can start at below the previous one on next iteration
		yImgPointer += 27
	}

	imageName, err := createSaveFileName(imagePath, urlImgName, ".png")
	if err != nil {
		return err
	}

	fullPathName, err := getFullSavePath(imageName, saveImagePath)
	if err != nil {
		return err
	}

	return dc.SavePNG(fullPathName)
}

// Returns path with the file name concatenated to it
func getFullSavePath(imageName, saveImagePath string) (string, error) {
	savePathLastChar := string(saveImagePath[len(saveImagePath)-1])

	// Check if path is closed with appropriate path separator (depending on OS)
	if savePathLastChar != string(os.PathSeparator) {
		if checkOS() == "linux" {
			saveImagePath += "/"
		} else {
			saveImagePath += "\\"
		}
	}

	// If path exists
	if _, err := os.Stat(saveImagePath); !os.IsNotExist(err) {
		return saveImagePath + imageName, nil
	} else {
		return "", err
	}
}
