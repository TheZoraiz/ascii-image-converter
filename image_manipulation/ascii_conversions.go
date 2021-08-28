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
	"strconv"

	"github.com/gookit/color"
)

// Reference taken from http://paulbourke.net/dataformats/asciiart/
var asciiTableSimple = " .:-=+*#%@"
var asciiTableDetailed = " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

// For each individual element of imgSet in ConvertToASCIISlice()
const MAX_VAL float64 = 65535

type AsciiChar struct {
	OriginalColor string
	SetColor      string
	Simple        string
	RgbValue      [3]uint32
}

/*
Converts the 2D image_conversions.AsciiPixel slice of image data (each instance representing each compressed pixel of original image)
to a 2D image_conversions.AsciiChar slice

If complex parameter is true, values are compared to 70 levels of color density in ASCII characters.
Otherwise, values are compared to 10 levels of color density in ASCII characters.
*/
func ConvertToAsciiChars(imgSet [][]AsciiPixel, negative, colored, complex bool, customMap string, fontColor [3]int) [][]AsciiChar {

	height := len(imgSet)
	width := len(imgSet[0])

	chosenTable := map[int]string{}

	// Turn ascii character-set string into map[int]string{} literal
	if customMap == "" {
		var charSet string

		if complex {
			charSet = asciiTableDetailed
		} else {
			charSet = asciiTableSimple
		}

		for index, char := range charSet {
			chosenTable[index] = string(char)
		}

	} else {
		chosenTable = map[int]string{}

		for index, char := range customMap {
			chosenTable[index] = string(char)
		}
	}

	result := make([][]AsciiChar, height)
	for i := range result {
		result[i] = make([]AsciiChar, width)
	}

	for i := 0; i < height; i++ {

		var tempSlice []AsciiChar

		for j := 0; j < width; j++ {
			value := float64(imgSet[i][j].charDepth)

			// Gets appropriate string index from chosenTable by percentage comparisons with its length
			tempFloat := (value / MAX_VAL) * float64(len(chosenTable))
			if value == MAX_VAL {
				tempFloat = float64(len(chosenTable) - 1)
			}
			tempInt := int(tempFloat)

			var r, g, b int

			if colored {
				r = int(imgSet[i][j].rgbValue[0])
				g = int(imgSet[i][j].rgbValue[1])
				b = int(imgSet[i][j].rgbValue[2])
			} else {
				r = int(imgSet[i][j].grayscaleValue[0])
				g = int(imgSet[i][j].grayscaleValue[1])
				b = int(imgSet[i][j].grayscaleValue[2])
			}

			if negative {
				// Select character from opposite side of table as well as turn pixels negative
				r = 255 - r
				g = 255 - g
				b = 255 - b

				// To preserve negative rgb values for saving png image later down the line, since it uses imgSet
				if colored {
					imgSet[i][j].rgbValue = [3]uint32{uint32(r), uint32(g), uint32(b)}
				} else {
					imgSet[i][j].grayscaleValue = [3]uint32{uint32(r), uint32(g), uint32(b)}
				}

				tempInt = (len(chosenTable) - 1) - tempInt
			}

			rStr := strconv.Itoa(r)
			gStr := strconv.Itoa(g)
			bStr := strconv.Itoa(b)

			var char AsciiChar

			char.OriginalColor = color.Sprintf("<fg="+rStr+","+gStr+","+bStr+">%v</>", chosenTable[tempInt])

			// If font color is not set, use a simple string. Otherwise, use True color
			if fontColor != [3]int{255, 255, 255} {
				fcR := strconv.Itoa(fontColor[0])
				fcG := strconv.Itoa(fontColor[1])
				fcB := strconv.Itoa(fontColor[2])

				char.SetColor = color.Sprintf("<fg="+fcR+","+fcG+","+fcB+">%v</>", chosenTable[tempInt])
			}

			char.Simple = chosenTable[tempInt]

			if colored {
				char.RgbValue = imgSet[i][j].rgbValue
			} else {
				char.RgbValue = imgSet[i][j].grayscaleValue
			}

			tempSlice = append(tempSlice, char)
		}
		result[i] = tempSlice
	}

	return result
}
