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
	"fmt"
	"image"
	"image/color"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nfnt/resize"
	"golang.org/x/crypto/ssh/terminal"
)

// This function shrinks the passed image according to terminal size and
// turns it into grayscale pixel by pixel to make ASCII character matching easier
//
// The returned 2D uint32 slice contains each corresponding pixel's value to be
// compared to an ASCII character
func ConvertToTerminalSizedSlices(img image.Image, dimensions []int) [][]uint32 {

	var terminalWidth, terminalHeight int

	// Get dimensions of current terminal
	if len(dimensions) == 0 {
		var err error
		terminalWidth, terminalHeight, err = terminal.GetSize(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		// Fix height
		var ratio float32
		imgHeight := img.Bounds().Max.Y
		imgWidth := img.Bounds().Max.X

		if imgHeight > imgWidth {
			ratio = float32(imgHeight) / float32(imgWidth)

			terminalHeight = int(ratio * float32(terminalWidth))
			terminalHeight -= terminalHeight / 2
		}

	} else {
		terminalWidth = dimensions[0]
		terminalHeight = dimensions[1]
	}

	if len(dimensions) > 0 {
		defaultTermWidth, _, _ := terminal.GetSize(int(os.Stdin.Fd()))
		if dimensions[0] > defaultTermWidth {
			fmt.Println("Error: Set width is larger than terminal width")
			os.Exit(1)
		}
	}

	// initialize imgSet 2D slice
	imgSet := make([][]uint32, terminalHeight)
	for i := range imgSet {
		imgSet[i] = make([]uint32, terminalWidth)
	}

	smallImg := resize.Resize(uint(terminalWidth), uint(terminalHeight), img, resize.Lanczos3)
	b := smallImg.Bounds()

	newImg := image.NewGray(b)

	for y := b.Min.Y; y < b.Max.Y; y++ {

		var temp []uint32
		for x := b.Min.X; x < b.Max.X; x++ {

			oldPixel := smallImg.At(x, y)
			pixel := color.GrayModel.Convert(oldPixel)
			newImg.Set(x, y, pixel)

			// We only need Red from Red, Green, Blue since they have the same value for grayscale images
			r, _, _, _ := pixel.RGBA()
			temp = append(temp, r)

		}
		imgSet[y] = temp
	}

	return imgSet
}
