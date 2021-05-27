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
var asciiTableSimple = map[int]string{
	0: " ",
	1: ".",
	2: ":",
	3: "-",
	4: "=",
	5: "+",
	6: "*",
	7: "#",
	8: "%",
	9: "@",
}

// Reference taken from http://paulbourke.net/dataformats/asciiart/
var asciiTableDetailed = map[int]string{
	0:  " ",
	1:  ".",
	2:  "'",
	3:  "`",
	4:  "^",
	5:  "\"",
	6:  ",",
	7:  ":",
	8:  ";",
	9:  "I",
	10: "l",
	11: "!",
	12: "i",
	13: ">",
	14: "<",
	15: "~",
	16: "+",
	17: "_",
	18: "-",
	19: "?",
	20: "]",
	21: "[",
	22: "}",
	23: "{",
	24: "1",
	25: ")",
	26: "(",
	27: "|",
	28: "/",
	29: "t",
	30: "f",
	31: "j",
	32: "r",
	33: "x",
	34: "n",
	35: "u",
	36: "v",
	37: "c",
	38: "z",
	39: "X",
	40: "Y",
	41: "U",
	42: "J",
	43: "C",
	44: "L",
	45: "Q",
	46: "0",
	47: "O",
	48: "Z",
	49: "m",
	50: "w",
	51: "q",
	52: "p",
	53: "d",
	54: "b",
	55: "k",
	56: "h",
	57: "a",
	58: "o",
	59: "*",
	60: "#",
	61: "M",
	62: "W",
	63: "&",
	64: "8",
	65: "%",
	66: "B",
	67: "@",
	68: "$",
}

// For each individual element of imgSet in ConvertToASCIISlice()
const MAX_VAL float32 = 65535

type AsciiChar struct {
	Colored string
	Simple  string
}

// Converts the 2D AsciiPixel slice of image data (each instance representing each pixel of original image)
// to a 2D AsciiChar slice with each colored and simple string having an ASCII character corresponding to
// the original grayscale and RGB values in AsciiPixel.
//
// If complex parameter is true, values are compared to 69 levels of color density in ASCII characters.
// Otherwise, values are compared to 10 levels of color density in ASCII characters.
func ConvertToAscii(imgSet [][]AsciiPixel, negative bool, colored bool, complex bool, customMap string) [][]AsciiChar {

	height := len(imgSet)
	width := len(imgSet[0])

	var chosenTable map[int]string

	if customMap == "" {
		if complex {
			chosenTable = asciiTableDetailed
		} else {
			chosenTable = asciiTableSimple
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
			value := float32(imgSet[i][j].grayscaleValue)

			// Gets appropriate string index from asciiTableSimple by percentage comparisons with its length
			tempFloat := (value / MAX_VAL) * float32(len(chosenTable))
			if value == MAX_VAL {
				tempFloat = float32(len(chosenTable) - 1)
			}
			tempInt := int(tempFloat)

			r := int(imgSet[i][j].rgbValue[0])
			g := int(imgSet[i][j].rgbValue[1])
			b := int(imgSet[i][j].rgbValue[2])

			if negative {
				// Select character from opposite side of table as well as turn pixels negative
				r = 255 - r
				g = 255 - g
				b = 255 - b

				tempInt = (len(chosenTable) - 1) - tempInt
			}

			rStr := strconv.Itoa(r)
			gStr := strconv.Itoa(g)
			bStr := strconv.Itoa(b)

			var char AsciiChar

			char.Colored = color.Sprintf("<fg="+rStr+","+gStr+","+bStr+">%v</>", chosenTable[tempInt])
			char.Simple = chosenTable[tempInt]

			tempSlice = append(tempSlice, char)
		}
		result[i] = tempSlice
	}

	return result
}
