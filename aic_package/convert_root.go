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

	"github.com/golang/freetype/truetype"
)

var pipedInputTypes = []string{
	"image/png",
	"image/jpeg",
	"image/webp",
	"image/tiff",
	"image/bmp",
}

// Return default configuration for flags.
// Can be sent directly to ConvertImage() for default ascii art
func DefaultFlags() Flags {
	return Flags{
		Complex:             false,
		Dimensions:          nil,
		Width:               0,
		Height:              0,
		SaveTxtPath:         "",
		SaveImagePath:       "",
		SaveGifPath:         "",
		SaveHtmlPath:         "",
		Negative:            false,
		Colored:             false,
		CharBackgroundColor: false,
		Grayscale:           false,
		CustomMap:           "",
		FlipX:               false,
		FlipY:               false,
		Full:                false,
		FontFilePath:        "",
		FontColor:           [3]int{255, 255, 255},
		SaveBackgroundColor: [4]int{0, 0, 0, 100},
		Braille:             false,
		Threshold:           128,
		Dither:              false,
		OnlySave:            false,
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
	saveHtmlPath = flags.SaveHtmlPath
	negative = flags.Negative
	colored = flags.Colored
	colorBg = flags.CharBackgroundColor
	grayscale = flags.Grayscale
	customMap = flags.CustomMap
	flipX = flags.FlipX
	flipY = flags.FlipY
	full = flags.Full
	fontPath = flags.FontFilePath
	fontColor = flags.FontColor
	saveBgColor = flags.SaveBackgroundColor
	braille = flags.Braille
	threshold = flags.Threshold
	dither = flags.Dither
	onlySave = flags.OnlySave

	inputIsGif = path.Ext(filePath) == ".gif"

	// Declared at the start since some variables are initially used in conditional blocks
	var (
		localFile       *os.File
		urlImgBytes     []byte
		urlImgName      string = ""
		pipedInputBytes []byte
		err             error
	)

	pathIsURl := isURL(filePath)

	// Different modes of reading data depending upon whether or not filePath is a url

	if filePath != "-" {
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

	} else {
		// Check file/data type of piped input

		if !isInputFromPipe() {
			return "", fmt.Errorf("there is no input being piped to stdin")
		}

		pipedInputBytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("unable to read piped input: %v", err)
		}

		fileType := http.DetectContentType(pipedInputBytes)
		invalidInput := true

		if fileType == "image/gif" {
			inputIsGif = true
			invalidInput = false

		} else {
			for _, inputType := range pipedInputTypes {
				if fileType == inputType {
					invalidInput = false
					break
				}
			}
		}

		// Not sure if I should uncomment this.
		// The output may be piped to another program and a warning would contaminate that
		if invalidInput {
			// fmt.Println("Warning: file type of piped input could not be determined, treating it as an image")
		}
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
	} else if braille {
		tempFont, _ = truetype.Parse(embeddedDejaVuObliqueFont)
	}

	if inputIsGif {
		return "", pathIsGif(filePath, urlImgName, pathIsURl, urlImgBytes, pipedInputBytes, localFile)
	} else {
		return pathIsImage(filePath, urlImgName, pathIsURl, urlImgBytes, pipedInputBytes, localFile)
	}
}
