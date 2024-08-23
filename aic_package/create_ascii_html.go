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
	"os"
	"strings"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
)

/*
Creates and saves an HTML file containing the passed ascii art. The HTML file will contain the ascii art as spans with inline color style if [colored] is true
*/
func createHtmlToSave(asciiArt [][]imgManip.AsciiChar, colored bool, saveHtmlPath, imagePath, urlImgName string, saveBgColor [4]int, onlySave bool) error {
	var builder strings.Builder

	backgroundColor := fmt.Sprintf("#%02x%02x%02x", saveBgColor[0], saveBgColor[1], saveBgColor[2])

	// Start the HTML file with the necessary headers and styles
	builder.WriteString(fmt.Sprintf(`<!DOCTYPE html><html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>ASCII Art</title>
		<style>
			body {
				background-color: %s;
			}
			span { 
				display: inline-block; 
				white-space: pre; 
				font-family: monospace; 
			}
		</style>
	</head><body><pre>`, backgroundColor))

	for _, line := range asciiArt {
		for _, asciiChar := range line {
			// Extract the RGB values from the RgbValue field
			rgb := asciiChar.RgbValue
			// Convert the character and color into an HTML span
			if colored {
				builder.WriteString(fmt.Sprintf(
					"<span style=\"color:rgb(%d,%d,%d)\">%s</span>",
					rgb[0], rgb[1], rgb[2], asciiChar.Simple,
				))
			} else {
				builder.WriteString(fmt.Sprintf(
					"<span>%s</span>",
					asciiChar.Simple,
				))
			}
		}
		// Add a line break after each line
		builder.WriteString("<br>")
	}

	// End the HTML document
	builder.WriteString(`</pre></body></html>`)


	saveFileName, err := createSaveFileName(imagePath, urlImgName, "-ascii-art.html")
	if err != nil {
		return err
	}

	savePathLastChar := string(saveHtmlPath[len(saveHtmlPath)-1])

	// Check if path is closed with appropriate path separator (depending on OS)
	if savePathLastChar != string(os.PathSeparator) {
		saveHtmlPath += string(os.PathSeparator)
	}

	// If path exists
	if _, err := os.Stat(saveHtmlPath); !os.IsNotExist(err) {
		err := os.WriteFile(saveHtmlPath+saveFileName, []byte(builder.String()), 0666)
		if err != nil {
			return err
		} else if onlySave {
			fmt.Println("Saved " + saveHtmlPath + saveFileName)
		}
		return nil
	} else {
		return fmt.Errorf("save path %v does not exist", saveHtmlPath)
	}
}
