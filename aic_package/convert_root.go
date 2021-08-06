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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	// Image format initialization
	_ "image/jpeg"
	_ "image/png"

	// Image format initialization
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/asaskevich/govalidator"
	"github.com/golang/freetype/truetype"
)

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
	// This overrides Flags.Grayscale
	Colored bool

	// Keep grayscale colors from the original image. This uses the True color
	// codes for the terminal and will work on saved .png and .gif files as well
	Grayscale bool

	// Pass custom ascii art characters as a string.
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
	// This will be ignored if Flags.SaveImagePath or Flags.SaveImagePath are not set
	FontFilePath string
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
)

// Return default configuration for flags.
// Can be sent directly to ConvertImage() for default ascii art
func DefaultFlags() Flags {
	return Flags{
		Complex:       false,
		Dimensions:    nil,
		Width:         0,
		Height:        0,
		SaveTxtPath:   "",
		SaveImagePath: "",
		SaveGifPath:   "",
		Negative:      false,
		Colored:       false,
		Grayscale:     false,
		CustomMap:     "",
		FlipX:         false,
		FlipY:         false,
		Full:          false,
		FontFilePath:  "",
	}
}

/*
Convert() takes an image or gif path/url as its first argument
and a aic_package.Flags literal as the second argument, with which it alters
the returned ascii art string.
*/
func Convert(filePath string, flags Flags) (string, error) {

	if flags.Dimensions == nil {
		dimensions = nil
	} else {
		dimensions = flags.Dimensions
	}
	width = flags.Width
	height = flags.Height
	complex = flags.Complex
	saveTxtPath = flags.SaveTxtPath
	saveImagePath = flags.SaveImagePath
	saveGifPath = flags.SaveGifPath
	negative = flags.Negative
	colored = flags.Colored
	grayscale = flags.Grayscale
	customMap = flags.CustomMap
	flipX = flags.FlipX
	flipY = flags.FlipY
	full = flags.Full
	fontPath = flags.FontFilePath

	// Declared at the start since some variables are initially used in conditional blocks
	var (
		localFile   *os.File
		urlImgBytes []byte
		urlImgName  string = ""
		err         error
	)

	pathIsURl := govalidator.IsRequestURL(filePath)

	// Different modes of reading data depending upon whether or not filePath is a url
	if pathIsURl {
		fmt.Printf("Fetching file from url...\r")

		retrievedImage, err := http.Get(filePath)
		if err != nil {
			return "", fmt.Errorf("can't fetch content: %v", err)
		}

		urlImgBytes, err = ioutil.ReadAll(retrievedImage.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read fetched content: %v", err)
		}
		defer retrievedImage.Body.Close()

		urlImgName = path.Base(filePath)
		fmt.Printf("                          \r") // To erase "Fetching image from url..." text from terminal

	} else {

		localFile, err = os.Open(filePath)
		if err != nil {
			return "", fmt.Errorf("unable to open file: %v", err)
		}
		defer localFile.Close()

	}

	// If path to font file is provided, use it
	if fontPath != "" {
		fontFile, err := ioutil.ReadFile(fontPath)
		if err != nil {
			return "", fmt.Errorf("unable to open font file: %v", err)
		}

		// tempFont is globally declared in aic_package/create_ascii_image.go
		if tempFont, err = truetype.Parse(fontFile); err != nil {
			return "", fmt.Errorf("unable to parse font file: %v", err)
		}
	}

	if path.Ext(filePath) == ".gif" {
		return pathIsGif(filePath, urlImgName, pathIsURl, urlImgBytes, localFile)
	} else {
		return pathIsImage(filePath, urlImgName, pathIsURl, urlImgBytes, localFile)
	}
}
