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
package cmd

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

	imgMani "github.com/TheZoraiz/ascii-image-converter/image_manipulation"
	"github.com/asaskevich/govalidator"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Flags
	cfgFile     string
	compl       bool
	dimensions  []int
	savePath    string
	negative    bool
	formatsTrue bool
	colored     bool
	customMap   string

	// Root commands
	rootCmd = &cobra.Command{
		Use:     "ascii-image-converter [image paths/urls]",
		Short:   "Converts images into ascii art",
		Version: "1.2.5",
		Long:    "This tool converts images into ascii art and prints them on the terminal.\nFurther configuration can be managed with flags.",

		// Not RunE since help text is getting larger and seeing it for every error impacts user experience
		Run: func(cmd *cobra.Command, args []string) {

			if formatsTrue {
				fmt.Printf("Supported image formats: JPEG/JPG, PNG, WEBP, BMP, TIFF/TIF\n\n")
				return
			}

			numberOfDimensions := len(dimensions)
			if dimensions != nil && numberOfDimensions != 2 {
				fmt.Printf("-d requires 2 dimensions, got %v\n\n", numberOfDimensions)
				return
			}

			if len(args) < 1 {
				fmt.Printf("Error: Need at least 1 image path/url\n\n")
				cmd.Help()
				return
			}

			if customMap != "" && len(customMap) < 2 {
				fmt.Printf("Need at least 2 characters for --map flag\n\n")
				return
			}

			for _, imagePath := range args {
				convertImage(imagePath)
			}
		},
	}
)

func checkOS() string {
	if string(os.PathSeparator) == "/" && string(os.PathListSeparator) == ":" {
		return "linux"
	} else {
		return "windows"
	}
}

func convertImage(imagePath string) {

	// Declared at the start since some variables are initially used in conditional blocks
	var pic *os.File
	var urlImgBytes []byte
	var urlImgName string = ""
	var err error

	pathIsURl := govalidator.IsRequestURL(imagePath)

	// Different modes of reading data depending upon whether or not imagePath is a url
	if pathIsURl {
		fmt.Printf("Fetching image from url...\r")

		retrievedImage, err := http.Get(imagePath)
		if err != nil {
			fmt.Printf("Error fetching image: %v\n\n", err)
			return
		}

		urlImgBytes, err = ioutil.ReadAll(retrievedImage.Body)
		if err != nil {
			fmt.Printf("Failed to read fetched content: %v\n\n", err)
			return
		}
		defer retrievedImage.Body.Close()

		urlImgName = path.Base(imagePath)
		fmt.Printf("                          \r") // To erase "Fetching image from url..." text from console
	} else {
		pic, err = os.Open(imagePath)
		if err != nil {
			fmt.Printf("Unable to open file: %v\n\n", err)
			return
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
		fmt.Printf("Error decoding %v. %v\n\n", imagePath, err)
		return
	}

	imgSet, err := imgMani.ConvertToAsciiPixels(imData, dimensions)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
		return
	}

	asciiSet := imgMani.ConvertToAscii(imgSet, negative, colored, compl, customMap)

	ascii := flattenAscii(asciiSet, colored)

	// Save ascii art before printing it, if --save flag is passed
	if savePath != "" {
		if err := saveAsciiArt(asciiSet, imagePath, urlImgName); err != nil {
			fmt.Printf("Error: %v\n\n", err)
			os.Exit(0) // Because this error will be thrown for every image passed to this function if we use "return"
		}
	}

	for _, line := range ascii {
		fmt.Println(line)
	}
}

// flattenAscii flattens a two-dimensional grid of ascii characters into a one dimension
// of lines of ascii
func flattenAscii(asciiSet [][]imgMani.AsciiChar, color bool) []string {
	var ascii []string

	for _, line := range asciiSet {
		var tempAscii []string

		for i := 0; i < len(line); i++ {
			if color {
				tempAscii = append(tempAscii, line[i].Colored)
			} else {
				tempAscii = append(tempAscii, line[i].Simple)
			}
		}

		ascii = append(ascii, strings.Join(tempAscii, ""))
	}

	return ascii
}

func saveAsciiArt(asciiSet [][]imgMani.AsciiChar, imagePath string, urlImgName string) error {
	// To make sure uncolored ascii art is the one saved
	saveAscii := flattenAscii(asciiSet, false)

	saveFileName, err := createSaveFileName(imagePath, urlImgName)
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
		return fmt.Errorf("Save path does not exist.")
	}
}

func createSaveFileName(imagePath string, urlImgName string) (string, error) {
	if urlImgName != "" {
		return urlImgName + "-ascii-art.txt", nil
	}

	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		return "", err
	}

	currName := fileInfo.Name()
	currExt := path.Ext(currName)
	newName := currName[:len(currName)-len(currExt)] // e.g. Grabs myImage from myImage.jpeg

	// Something like myImage.jpeg-ascii-art.txt
	return newName + "." + currExt[1:] + "-ascii-art.txt", nil
}

// Cobra configuration from here on

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with the colors from original image\n(Can work with the -n flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&compl, "complex", "c", false, "Display ascii characters in a larger range\nMay result in higher quality\n")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length\ne.g. -d 100,30 (defaults to terminal height)\n")
	rootCmd.PersistentFlags().BoolVarP(&formatsTrue, "formats", "f", false, "Display supported image formats\n")
	rootCmd.PersistentFlags().StringVarP(&customMap, "map", "m", "", "Give custom ascii characters to map against\nOrdered from darkest to lightest\ne.g. -m \" .-+#@\" (Quotation marks excluded from map)\n(Cancels --complex flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors\n(Can work with the --color flag)\n")
	rootCmd.PersistentFlags().StringVarP(&savePath, "save", "s", "", "Save ascii art in the format:\n<image-name>.<image-extension>-ascii-art.txt\nFile will be saved in passed path\n(pass . for current directory)\n")

	defaultUsageTemplate := rootCmd.UsageTemplate()
	rootCmd.SetUsageTemplate("\nCopyright © 2021 Zoraiz Hassan <hzoraiz8@gmail.com>\n" +
		"Distributed under the Apache License Version 2.0 (Apache-2.0)\n" +
		"For further details, visit https://github.com/TheZoraiz/ascii-image-converter\n\n" +
		defaultUsageTemplate)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ascii-image-converter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ascii-image-converter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
