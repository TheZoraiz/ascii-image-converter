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

	"github.com/TheZoraiz/ascii-image-converter/aic_package/winsize"
	"github.com/disintegration/imaging"
	gookitColor "github.com/gookit/color"
	"github.com/makeworld-the-better-one/dither/v2"
)

func ditherImage(img image.Image) image.Image {

	palette := []color.Color{
		color.Black,
		color.White,
	}

	d := dither.NewDitherer(palette)
	d.Matrix = dither.FloydSteinberg

	return d.DitherCopy(img)
}

func resizeImage(img image.Image, full, isBraille bool, dimensions []int, width, height int) (image.Image, error) {

	var asciiWidth, asciiHeight int
	var smallImg image.Image

	terminalWidth, terminalHeight, err := winsize.GetTerminalSize()
	if err != nil {
		return nil, err
	}

	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())
	aspectRatio := imgWidth / imgHeight

	if full {
		asciiWidth = terminalWidth - 1
		asciiHeight = int(float64(asciiWidth) / aspectRatio)
		asciiHeight = int(0.5 * float64(asciiHeight))

	} else if (width != 0 || height != 0) && len(dimensions) == 0 {
		// If either width or height is set and dimensions aren't given

		if width > terminalWidth-1 {
			return nil, fmt.Errorf("set width must be lower than terminal width")
		}

		if width != 0 && height == 0 {
			// If width is set and height is not set, use width to calculate aspect ratio

			asciiWidth = width
			asciiHeight = int(float64(asciiWidth) / aspectRatio)
			asciiHeight = int(0.5 * float64(asciiHeight))

			if asciiHeight == 0 {
				asciiHeight = 1
			}

		} else if height != 0 && width == 0 {
			// If height is set and width is not set, use height to calculate aspect ratio

			asciiHeight = height
			asciiWidth = int(float64(asciiHeight) * aspectRatio)
			asciiWidth = int(2 * float64(asciiWidth))

			if asciiWidth == 0 {
				asciiWidth = 1
			}

			if asciiWidth > terminalWidth-1 {
				return nil, fmt.Errorf("width calculated with aspect ratio exceeds terminal width")
			}

		} else {
			return nil, fmt.Errorf("both width and height can't be set. Use dimensions instead")
		}

	} else if len(dimensions) == 0 {
		// This condition calculates aspect ratio according to terminal height

		asciiHeight = terminalHeight - 1
		asciiWidth = int(float64(asciiHeight) * aspectRatio)
		asciiWidth = int(2 * float64(asciiWidth))

		// If ascii width exceeds terminal width, change ratio with respect to terminal width
		if asciiWidth >= terminalWidth {
			asciiWidth = terminalWidth - 1
			asciiHeight = int(float64(asciiWidth) / aspectRatio)
			asciiHeight = int(0.5 * float64(asciiHeight))
		}

	} else {
		asciiWidth = dimensions[0]
		asciiHeight = dimensions[1]
	}

	// Repeated despite being in cmd/root.go to maintain support for library
	//
	// If there are passed dimensions, check whether the width exceeds terminal width
	if len(dimensions) > 0 && !full {
		if dimensions[0] > terminalWidth-1 {
			return nil, fmt.Errorf("set width must be lower than terminal width")
		}
	}

	if isBraille {
		asciiWidth *= 2
		asciiHeight *= 4
	}
	smallImg = imaging.Resize(img, asciiWidth, asciiHeight, imaging.Lanczos)

	return smallImg, nil
}

func resizeImageNoTerm(img image.Image, isBraille bool, dimensions []int, width, height int) (image.Image, error) {

	var asciiWidth, asciiHeight int
	var smallImg image.Image

	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())
	aspectRatio := imgWidth / imgHeight

	if (width != 0 || height != 0) && len(dimensions) == 0 {
		if width != 0 && height == 0 {

			asciiWidth = width
			asciiHeight = int(float64(asciiWidth) / aspectRatio)
			asciiHeight = int(0.5 * float64(asciiHeight))

			if asciiHeight == 0 {
				asciiHeight = 1
			}

		} else if height != 0 && width == 0 {

			asciiHeight = height
			asciiWidth = int(float64(asciiHeight) * aspectRatio)
			asciiWidth = int(2 * float64(asciiWidth))

			if asciiWidth == 0 {
				asciiWidth = 1
			}

		} else {
			return nil, fmt.Errorf("error: both width and height can't be set. Use dimensions instead")
		}

	} else if len(dimensions) != 0 {
		asciiWidth = dimensions[0]
		asciiHeight = dimensions[1]
	} else {
		return nil, fmt.Errorf("error: at least one of width, height or dimensions should be passed for NoTermSizeComparison")
	}

	if isBraille {
		asciiWidth *= 2
		asciiHeight *= 4
	}
	smallImg = imaging.Resize(img, asciiWidth, asciiHeight, imaging.Lanczos)

	return smallImg, nil
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

var termColorLevel string = gookitColor.TermColorLevel().String()

// This functions calculates terminal color level between rgb colors, 256-colors, 16-colors
// and returns the character with escape codes appropriately
func getColoredCharForTerm(r, g, b uint8, char string, background bool) (string, error) {
	var coloredChar string

	if termColorLevel == "millions" {
		colorRenderer := gookitColor.RGB(uint8(r), uint8(g), uint8(b), background)
		coloredChar = colorRenderer.Sprintf("%v", char)

	} else if termColorLevel == "hundreds" {
		colorRenderer := gookitColor.RGB(uint8(r), uint8(g), uint8(b), background).C256()
		coloredChar = colorRenderer.Sprintf("%v", char)

	} else {
		return "", fmt.Errorf("your terminal supports neither 24-bit nor 8-bit colors. Other coloring options aren't available")
	}

	return coloredChar, nil
}
