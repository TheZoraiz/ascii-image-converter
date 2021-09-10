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
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	imgManip "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
)

type GifFrame struct {
	asciiCharSet [][]imgManip.AsciiChar
	delay        int
}

/*
This function grabs each image frame from passed gif and turns it into ascii art. If SaveGifPath flag is passed,
it'll turn each ascii art into an image instance of the same dimensions as the original gif and save them
as an ascii art gif.

Multi-threading has been implemented in multiple places due to long execution time
*/
func pathIsGif(gifPath, urlImgName string, pathIsURl bool, urlImgBytes []byte, localGif *os.File) error {

	var (
		originalGif *gif.GIF
		err         error
	)

	if pathIsURl {
		originalGif, err = gif.DecodeAll(bytes.NewReader(urlImgBytes))
	} else {
		originalGif, err = gif.DecodeAll(localGif)
	}
	if err != nil {
		return fmt.Errorf("can't decode %v: %v", gifPath, err)
	}

	var (
		asciiArtSet    = make([]string, len(originalGif.Image))
		gifFramesSlice = make([]GifFrame, len(originalGif.Image))

		counter             = 0
		concurrentProcesses = 0
		wg                  sync.WaitGroup
		hostCpuCount        = runtime.NumCPU()
	)

	fmt.Printf("Generating ascii art... 0%%\r")

	// Get first frame of gif and its dimensions
	firstGifFrame := originalGif.Image[0].SubImage(originalGif.Image[0].Rect)
	firstGifFrameWidth := firstGifFrame.Bounds().Dx()
	firstGifFrameHeight := firstGifFrame.Bounds().Dy()

	// Multi-threaded loop to decrease execution time
	for i, frame := range originalGif.Image {

		wg.Add(1)
		concurrentProcesses++

		go func(i int, frame *image.Paletted) {

			frameImage := frame.SubImage(frame.Rect)

			// If a frame is found that is smaller than the first frame, then this gif contains smaller subimages that are
			// positioned inside the original gif. This behavior isn't supported by this app
			if firstGifFrameWidth != frameImage.Bounds().Dx() || firstGifFrameHeight != frameImage.Bounds().Dy() {
				fmt.Printf("Error: GIF contains subimages smaller than default width and height\nProcess aborted because ascii-image-converter doesn't support subimage placement and transparency in GIFs\n\n")
				os.Exit(0)
			}

			var imgSet [][]imgManip.AsciiPixel

			imgSet, err = imgManip.ConvertToAsciiPixels(frameImage, dimensions, width, height, flipX, flipY, full, braille, dither, noTermSizeComparison)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(0)
			}

			var asciiCharSet [][]imgManip.AsciiChar
			if braille {
				asciiCharSet, err = imgManip.ConvertToBrailleChars(imgSet, negative, colored, grayscale, colorBg, fontColor, threshold)
			} else {
				asciiCharSet, err = imgManip.ConvertToAsciiChars(imgSet, negative, colored, grayscale, complex, colorBg, customMap, fontColor)
			}
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(0)
			}

			gifFramesSlice[i].asciiCharSet = asciiCharSet
			gifFramesSlice[i].delay = originalGif.Delay[i]

			ascii := flattenAscii(asciiCharSet, colored || grayscale, false)

			asciiArtSet[i] = strings.Join(ascii, "\n")

			counter++
			percentage := int((float64(counter) / float64(len(originalGif.Image))) * 100)
			fmt.Printf("Generating ascii art... " + strconv.Itoa(percentage) + "%%\r")

			wg.Done()

		}(i, frame)

		// Limit concurrent processes according to host's CPU count to avoid overwhelming memory
		if concurrentProcesses == hostCpuCount {
			wg.Wait()
			concurrentProcesses = 0
		}
	}

	wg.Wait()
	fmt.Printf("                              \r")

	// Save ascii art as .gif file before displaying it, if --save-gif flag is passed
	if saveGifPath != "" {

		// Storing save path string before executing ascii art to gif conversion
		// This is done to avoid wasting time for invalid path errors

		saveFileName, err := createSaveFileName(gifPath, urlImgName, "-ascii-art.gif")
		if err != nil {
			return err
		}

		fullPathName, err := getFullSavePath(saveFileName, saveGifPath)
		if err != nil {
			return fmt.Errorf("can't save file: %v", err)
		}

		// Initializing some constants for gif. Done outside loop to save execution
		outGif := &gif.GIF{
			LoopCount: originalGif.LoopCount,
		}
		opts := gif.Options{
			NumColors: 256,
			Drawer:    draw.FloydSteinberg,
		}

		// Initializing slices for each ascii art image as well as delay
		var (
			palettedImageSlice = make([]*image.Paletted, len(gifFramesSlice))
			delaySlice         = make([]int, len(gifFramesSlice))
		)

		// For the purpose of displaying counter and limiting concurrent processes
		counter = 0
		concurrentProcesses = 0

		fmt.Printf("Saving gif... 0%%\r")

		// Multi-threaded loop to decrease execution time
		for i, gifFrame := range gifFramesSlice {

			wg.Add(1)
			concurrentProcesses++

			go func(i int, gifFrame GifFrame) {

				img := originalGif.Image[i].SubImage(originalGif.Image[i].Rect)

				tempImg, err := createGifFrameToSave(
					gifFrame.asciiCharSet,
					img,
					colored || grayscale,
				)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(0)
				}

				// Following code takes tempImg as image.Image instance and converts it into *image.Paletted instance
				b := tempImg.Bounds()

				palettedImg := image.NewPaletted(b, palette.Plan9[:opts.NumColors])

				opts.Drawer.Draw(palettedImg, b, tempImg, image.Point{})

				palettedImageSlice[i] = palettedImg
				delaySlice[i] = gifFrame.delay

				counter++
				percentage := int((float64(counter) / float64(len(gifFramesSlice))) * 100)
				fmt.Printf("Saving gif... " + strconv.Itoa(percentage) + "%%\r")

				wg.Done()

			}(i, gifFrame)

			// Limit concurrent processes according to host's CPU count to avoid overwhelming memory
			if concurrentProcesses == hostCpuCount {
				wg.Wait()
				concurrentProcesses = 0
			}

		}

		wg.Wait()

		outGif.Image = palettedImageSlice
		outGif.Delay = delaySlice

		gifFile, err := os.OpenFile(fullPathName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("can't save file: %v", err)
		}
		defer gifFile.Close()

		gif.EncodeAll(gifFile, outGif)

		fmt.Printf("                     \r")
	}

	// Display the gif
	loopCount := 0
	for {
		for i, asciiFrame := range asciiArtSet {
			clearScreen()
			fmt.Println(asciiFrame)
			time.Sleep(time.Duration((time.Second * time.Duration(originalGif.Delay[i])) / 100))
		}

		// If gif is infinite loop
		if originalGif.LoopCount == 0 {
			continue
		}

		loopCount++
		if loopCount == originalGif.LoopCount {
			break
		}
	}

	return nil
}
