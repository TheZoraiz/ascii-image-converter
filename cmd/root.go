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
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"

	// Image format initialization
	_ "image/jpeg"
	_ "image/png"

	// Image format initialization
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	imgMani "github.com/TheZoraiz/ascii-image-converter/image_manipulation"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
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

	// Root commands
	rootCmd = &cobra.Command{
		Use:   "ascii-image-converter [image path]",
		Short: "Converts images into ascii format",
		Long:  `This tool converts images into ascii format and prints them onto the terminal window. Further configuration can be managed with flags`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if formatsTrue {
				fmt.Println("Supported image formats: JPEG/JPG, PNG, WEBP, BMP, TIFF/TIF")
				return nil
			}

			if len(args) != 1 {
				return fmt.Errorf("Requires 1 image path, got %v", len(args))
			}

			numberOfDimensions := len(dimensions)
			if dimensions != nil && numberOfDimensions != 2 {
				return fmt.Errorf("-d requires 2 dimensions, got %v", numberOfDimensions)
			}

			imagePath := args[0]

			return convertImage(imagePath)
		},
	}
)

func convertImage(imagePath string) error {

	pic, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("Unable to open file: %w", err)
	}
	defer pic.Close()

	imData, _, err := image.Decode(pic)
	if err != nil {
		return fmt.Errorf("Unable to decode file: %w", err)
	}

	imgSet, err := imgMani.ConvertToAsciiPixels(imData, dimensions)
	if err != nil {
		return err
	}

	var asciiSet [][]imgMani.AsciiChar
	asciiSet = imgMani.ConvertToAscii(imgSet, negative, colored, compl)

	var ascii []string
	ascii = flattenAscii(asciiSet, colored)

	for _, line := range ascii {
		fmt.Println(line)
	}

	if savePath != "" {
		// To make sure uncolored ascii art is the one saved
		saveAscii := flattenAscii(asciiSet, false)
		if savePath == "." {
			savePath = "./"
		}
		return ioutil.WriteFile(savePath+"ascii-image.txt", []byte(strings.Join(saveAscii, "\n")), 0777)
	}
	return nil
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

// Cobra configuration from here on

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with the colors from original image (Can work with the -n flag)")
	rootCmd.PersistentFlags().BoolVarP(&compl, "complex", "c", false, "Display ascii characters in a larger range, may result in higher quality")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length e.g. 100,30 (defaults to terminal size)")
	rootCmd.PersistentFlags().BoolVarP(&formatsTrue, "formats", "f", false, "Display supported image formats")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors (Can work with the -C flag)")
	rootCmd.PersistentFlags().StringVarP(&savePath, "save", "s", "", "Save ascii art in an ascii-image.txt file in a given path (pass ./ for current directory)")
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
