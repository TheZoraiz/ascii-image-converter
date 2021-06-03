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
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	// Image format initialization
	_ "image/jpeg"
	_ "image/png"

	// Image format initialization
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"

	"github.com/asaskevich/govalidator"
)

// Return default configuration for flags.
// Can be sent directly to ConvertImage() for default ascii art
func DefaultFlags() map[string]interface{} {
	return map[string]interface{}{
		"complex":       false,
		"dimensions":    nil,
		"saveTxtPath":   "",
		"saveImagePath": "",
		"negative":      false,
		"colored":       false,
		"customMap":     "",
		"flipX":         false,
		"flipY":         false,
	}
}

/*
ConvertImage takes an image path/url as its first argument
and a map of flags as the second argument, with which it alters
the returned ascii art string.

The "flags" argument should be declared as follows before passing:

 flags := map[string]interface{}{
 	"complex": bool, // Pass true for using complex character set
 	"dimensions": []int, // Pass 2 integer dimensions. Pass nil to ignore
	"saveTxtPath": string, // System path to save the ascii art string as a .txt  file. Pass "" to ignore
	"saveImagePath": string, // System path to save the ascii art string as a .png  file. Pass "" to ignore
 	"negative": bool, // Pass true for negative color-depth ascii art
 	"colored": bool, // Pass true for returning colored ascii string
 	"customMap": string, // Custom map of ascii chars e.g. " .-+#@" . Nullifies "complex" flag. Pass "" to ignore.
 	"flipX": bool, // Pass true to return horizontally flipped ascii art
 	"flipY": bool, // Pass true to return vertically flipped ascii art
 }
*/
func ConvertImage(imagePath string, flags map[string]interface{}) (string, error) {

	var dimensions []int
	if flags["dimensions"] == nil {
		dimensions = nil
	} else {
		dimensions = flags["dimensions"].([]int)
	}

	complex := flags["complex"].(bool)
	saveTxtPath := flags["saveTxtPath"].(string)
	saveImagePath := flags["saveImagePath"].(string)
	negative := flags["negative"].(bool)
	colored := flags["colored"].(bool)
	customMap := flags["customMap"].(string)
	flipX := flags["flipX"].(bool)
	flipY := flags["flipY"].(bool)

	// Declared at the start since some variables are initially used in conditional blocks
	var (
		pic         *os.File
		urlImgBytes []byte
		urlImgName  string = ""
		err         error
	)

	pathIsURl := govalidator.IsRequestURL(imagePath)

	// Different modes of reading data depending upon whether or not imagePath is a url
	if pathIsURl {
		fmt.Printf("Fetching image from url...\r")

		retrievedImage, err := http.Get(imagePath)
		if err != nil {
			return "", fmt.Errorf("can't fetch image: %v", err)
		}

		urlImgBytes, err = ioutil.ReadAll(retrievedImage.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read fetched content: %v", err)
		}
		defer retrievedImage.Body.Close()

		urlImgName = path.Base(imagePath)
		fmt.Printf("                          \r") // To erase "Fetching image from url..." text from console

	} else {

		pic, err = os.Open(imagePath)
		if err != nil {
			return "", fmt.Errorf("unable to open file: %v", err)
		}
		defer pic.Close()

	}

	var imData image.Image

	if pathIsURl {
		imData, _, err = image.Decode(bytes.NewReader(urlImgBytes))
	} else {
		imData, _, err = image.Decode(pic)
	}
	if err != nil {
		return "", fmt.Errorf("can't decode %v: %v", imagePath, err)
	}

	// Ascii art height and width are important for creating png image to save
	imgSet, imgWidth, imgHeight, err := imgManip.ConvertToAsciiPixels(imData, dimensions, flipX, flipY)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	asciiSet := imgManip.ConvertToAscii(imgSet, negative, colored, complex, customMap)

	// Save ascii art as .png image before printing it, if --save-img flag is passed
	if saveImagePath != "" {
		if err := createImageToSave(asciiSet, imgWidth, imgHeight, colored, saveImagePath, imagePath, urlImgName); err != nil {
			return "", fmt.Errorf("can't save file: %v", err)
		}
	}

	// Save ascii art as .txt file before printing it, if --save-txt flag is passed
	if saveTxtPath != "" {
		if err := saveAsciiArt(asciiSet, imagePath, saveTxtPath, urlImgName); err != nil {
			return "", fmt.Errorf("can't save file: %v", err)
		}
	}

	ascii := flattenAscii(asciiSet, colored)

	result := strings.Join(ascii, "\n")

	return result, nil
}

func checkOS() string {
	if string(os.PathSeparator) == "/" && string(os.PathListSeparator) == ":" {
		return "linux"
	} else {
		return "windows"
	}
}

// flattenAscii flattens a two-dimensional grid of ascii characters into a one dimension
// of lines of ascii
func flattenAscii(asciiSet [][]imgManip.AsciiChar, colored bool) []string {
	var ascii []string

	for _, line := range asciiSet {
		var tempAscii []string

		for i := 0; i < len(line); i++ {
			if colored {
				tempAscii = append(tempAscii, line[i].Colored)
			} else {
				tempAscii = append(tempAscii, line[i].Simple)
			}
		}

		ascii = append(ascii, strings.Join(tempAscii, ""))
	}

	return ascii
}

func saveAsciiArt(asciiSet [][]imgManip.AsciiChar, imagePath, savePath, urlImgName string) error {
	// To make sure uncolored ascii art is the one saved as .txt
	saveAscii := flattenAscii(asciiSet, false)

	saveFileName, err := createSaveFileName(imagePath, urlImgName, ".txt")
	if err != nil {
		return err
	}

	savePathLastChar := string(savePath[len(savePath)-1])

	// Check if path is closed with appropriate path separator (depending on OS)
	if savePathLastChar != string(os.PathSeparator) {
		if checkOS() == "linux" {
			savePath += "/"
		} else {
			savePath += "\\"
		}
	}

	// If path exists
	if _, err := os.Stat(savePath); !os.IsNotExist(err) {
		return ioutil.WriteFile(savePath+saveFileName, []byte(strings.Join(saveAscii, "\n")), 0666)
	} else {
		return fmt.Errorf("save path %v does exist", savePath)
	}
}

// Returns new image file name along with extension
func createSaveFileName(imagePath, urlImgName, newExtension string) (string, error) {
	if urlImgName != "" {
		currExt := path.Ext(urlImgName)
		newName := urlImgName[:len(urlImgName)-len(currExt)] // e.g. Grabs myImage from myImage.jpeg

		return newName + "-ascii-art" + newExtension, nil
	}

	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		return "", err
	}

	currName := fileInfo.Name()
	currExt := path.Ext(currName)
	newName := currName[:len(currName)-len(currExt)] // e.g. Grabs myImage from myImage.jpeg

	return newName + "-ascii-art" + newExtension, nil
}
