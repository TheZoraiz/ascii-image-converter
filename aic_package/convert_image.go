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
	"os"
	"strings"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
)

// This function decodes the passed image and returns an ascii art string, optionaly saving it as a .txt and/or .png file
func pathIsImage(imagePath, urlImgName string, pathIsURl bool, urlImgBytes, pipedInputBytes []byte, localImg *os.File) (string, error) {

	var (
		imData image.Image
		err    error
	)

	if imagePath == "-" {
		imData, _, err = image.Decode(bytes.NewReader(pipedInputBytes))
	} else if pathIsURl {
		imData, _, err = image.Decode(bytes.NewReader(urlImgBytes))
	} else {
		imData, _, err = image.Decode(localImg)
	}
	if err != nil {
		if imagePath == "-" {
			return "", fmt.Errorf("can't decode piped input: %v", err)
		} else {
			return "", fmt.Errorf("can't decode %v: %v", imagePath, err)
		}
	}

	imgSet, err := imgManip.ConvertToAsciiPixels(imData, dimensions, width, height, flipX, flipY, full, braille, dither)
	if err != nil {
		return "", err
	}

	var asciiSet [][]imgManip.AsciiChar

	if braille {
		asciiSet, err = imgManip.ConvertToBrailleChars(imgSet, negative, colored, grayscale, colorBg, fontColor, threshold)
	} else {
		asciiSet, err = imgManip.ConvertToAsciiChars(imgSet, negative, colored, grayscale, complex, colorBg, customMap, fontColor)
	}
	if err != nil {
		return "", err
	}

	// Save ascii art as .png image before printing it, if --save-img flag is passed
	if saveImagePath != "" {
		if err := createImageToSave(
			asciiSet,
			colored || grayscale,
			saveImagePath,
			imagePath,
			urlImgName,
			onlySave,
		); err != nil {

			return "", fmt.Errorf("can't save file: %v", err)
		}
	}

	// Save ascii art as .txt file before printing it, if --save-txt flag is passed
	if saveTxtPath != "" {
		if err := saveAsciiArt(
			asciiSet,
			imagePath,
			saveTxtPath,
			urlImgName,
			onlySave,
		); err != nil {

			return "", fmt.Errorf("can't save file: %v", err)
		}
	}

	// Save ascii art as .html file before printing it, if --save-html flag is passed
	if saveHtmlPath != "" {
		if err := createHtmlToSave(
			asciiSet,
			colored,
			saveHtmlPath,
			imagePath,
			urlImgName,
			saveBgColor,
			onlySave,
		); err != nil {

			return "", fmt.Errorf("can't save file: %v", err)
		}
	}

	ascii := flattenAscii(asciiSet, colored || grayscale, false)
	result := strings.Join(ascii, "\n")

	if onlySave {
		return "", nil
	}
	return result, nil
}
