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
)

type Flags struct {
	Dimensions    []int
	Complex       bool
	SaveTxtPath   string
	SaveImagePath string
	SaveGifPath   string
	Negative      bool
	Colored       bool
	Grayscale     bool
	CustomMap     string
	FlipX         bool
	FlipY         bool
	Full          bool
}

var (
	dimensions    []int
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
)

// Return default configuration for flags.
// Can be sent directly to ConvertImage() for default ascii art
func DefaultFlags() Flags {
	return Flags{
		Complex:       false,
		Dimensions:    nil,
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
	}
}

/*
Convert takes an image or gif path/url as its first argument
and a Flags literal as the second argument, with which it alters
the returned ascii art string.

The "flags" argument should be declared as follows before passing:

 flags := aic_package.Flags{
 	Complex: bool, // Pass true for using complex character set
 	Dimensions: []int, // Pass 2 integer dimensions. Pass nil to ignore
	SaveTxtPath: string, // System path to save the ascii art string as a .txt  file. Pass "" to ignore
	SavefilePath: string, // System path to save the ascii art string as a .png  file. Pass "" to ignore
	SaveGifPath : string, // System path to save the ascii art gif as a .gif  file. Pass "" to ignore
 	Negative: bool, // Pass true for negative color-depth ascii art
 	Colored: bool, // Pass true for returning colored ascii string
	Grayscale: bool // Pass true for returning grayscale ascii string
 	CustomMap: string, // Custom map of ascii chars e.g. " .-+#@" . Nullifies "complex" flag. Pass "" to ignore.
 	FlipX: bool, // Pass true to return horizontally flipped ascii art
 	FlipY: bool, // Pass true to return vertically flipped ascii art
	Full: bool, // Pass true to use full terminal as ascii height
 }
*/
func Convert(filePath string, flags Flags) (string, error) {

	if flags.Dimensions == nil {
		dimensions = nil
	} else {
		dimensions = flags.Dimensions
	}
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
		fmt.Printf("                          \r") // To erase "Fetching image from url..." text from console

	} else {

		localFile, err = os.Open(filePath)
		if err != nil {
			return "", fmt.Errorf("unable to open file: %v", err)
		}
		defer localFile.Close()

	}

	if path.Ext(filePath) == ".gif" {
		return pathIsGif(filePath, urlImgName, pathIsURl, urlImgBytes, localFile)
	} else {
		return pathIsImage(filePath, urlImgName, pathIsURl, urlImgBytes, localFile)
	}
}
