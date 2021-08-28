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

type Flags struct {
	// Set dimensions of ascii art. Accepts a slice of 2 integers
	// e.g. []int{60,30}.
	// This overrides Flags.Width and Flags.Height
	Dimensions []int

	// Set width of ascii art while calculating height from aspect ratio.
	// Setting this along with Flags.Height will throw an error
	Width int

	// Set height of ascii art while calculating width from aspect ratio.
	// Setting this along with Flags.Width will throw an error
	Height int

	// Use set of 69 characters instead of the default 10
	Complex bool

	// Path to save ascii art .txt file
	SaveTxtPath string

	// Path to save ascii art .png file
	SaveImagePath string

	// Path to save ascii art .gif file, if gif is passed
	SaveGifPath string

	// Invert ascii art character mapping as well as colors
	Negative bool

	// Keep colors from the original image. This uses the True color codes for
	// the terminal and will work on saved .png and .gif files as well.
	// This overrides Flags.Grayscale and Flags.FontColor
	Colored bool

	// Keep grayscale colors from the original image. This uses the True color
	// codes for the terminal and will work on saved .png and .gif files as well
	// This overrides Flags.FontColor
	Grayscale bool

	// Pass custom ascii art characters as a string.
	// e.g. " .-=+#@".
	// This overrides Flags.Complex
	CustomMap string

	// Flip ascii art horizontally
	FlipX bool

	// Flip ascii art vertically
	FlipY bool

	// Use terminal width to calculate ascii art size while keeping aspect ratio.
	// This overrides Flags.Dimensions, Flags.Width and Flags.Height
	Full bool

	// File path to a font .ttf file to use when saving ascii art gif or png file.
	// This will be ignored if Flags.SaveImagePath or Flags.SaveGifPath are not set
	FontFilePath string

	// Font RGB color in saved png or gif files.
	// This will be ignored if Flags.SaveImagePath or Flags.SaveGifPath are not set
	FontColor [3]int

	// Background RGB color in saved png or gif files.
	// This will be ignored if Flags.SaveImagePath or Flags.SaveGifPath are not set
	SaveBackgroundColor [3]int
}

var (
	dimensions    []int
	width         int
	height        int
	complex       bool
	saveTxtPath   string
	saveImagePath string
	saveGifPath   string
	grayscale     bool
	negative      bool
	colored       bool
	customMap     string
	flipX         bool
	flipY         bool
	full          bool
	fontPath      string
	fontColor     [3]int
	saveBgColor   [3]int
)
