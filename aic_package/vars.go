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

	// Path to save ascii art .html file
	SaveHtmlPath string

	// Path to save ascii art .gif file, if gif is passed
	SaveGifPath string

	// Invert ascii art character mapping as well as colors
	Negative bool

	// Keep colors from the original image. This uses the True color codes for
	// the terminal and will work on saved .png and .gif files as well.
	// This overrides Flags.Grayscale and Flags.FontColor
	Colored bool

	// If Flags.Colored, Flags.Grayscale or Flags.FontColor is set, use that color
	// on each character's background in the terminal
	CharBackgroundColor bool

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

	// Font RGB color for terminal display and saved png or gif files.
	FontColor [3]int

	// Background RGB color in saved png or gif files.
	// This will be ignored if Flags.SaveImagePath or Flags.SaveGifPath are not set
	SaveBackgroundColor [4]int

	// Use braille characters instead of ascii. Terminal must support UTF-8 encoding.
	// Otherwise, problems may be encountered with colored or even uncolored braille art.
	// This overrides Flags.Complex and Flags.CustomMap
	Braille bool

	// Threshold for braille art if Flags.Braille is set to true. Value provided must
	// be between 0 and 255. Ideal value is 128.
	// This will be ignored if Flags.Braille is not set
	Threshold int

	// Apply FloydSteinberg dithering on an image before ascii conversion. This option
	// is meant for braille art. Therefore, it will be ignored if Flags.Braille is false
	Dither bool

	// If Flags.SaveImagePath, Flags.SaveTxtPath or Flags.SaveGifPath are set, then don't
	// print on terminal
	OnlySave bool
}

var (
	dimensions    []int
	width         int
	height        int
	complex       bool
	saveTxtPath   string
	saveImagePath string
	saveHtmlPath  string
	saveGifPath   string
	grayscale     bool
	negative      bool
	colored       bool
	colorBg       bool
	customMap     string
	flipX         bool
	flipY         bool
	full          bool
	fontPath      string
	fontColor     [3]int
	saveBgColor   [4]int
	braille       bool
	threshold     int
	dither        bool
	onlySave      bool
	inputIsGif    bool
)
