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
	"path/filepath"
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
	customMap   string

	// Root commands
	rootCmd = &cobra.Command{
		Use:     "ascii-image-converter [image-paths]",
		Short:   "Converts images into ascii format",
		Version: "1.2.4",
		Long:    "This tool converts images into ascii art and prints them on the terminal.\nFurther configuration can be managed with flags.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if formatsTrue {
				fmt.Println("Supported image formats: JPEG/JPG, PNG, WEBP, BMP, TIFF/TIF")
				return nil
			}

			numberOfDimensions := len(dimensions)
			if dimensions != nil && numberOfDimensions != 2 {
				return fmt.Errorf("-d requires 2 dimensions, got %v", numberOfDimensions)
			}

			if len(args) < 1 {
				return fmt.Errorf("Need at least 1 image path")
			}

			if len(customMap) < 2 && customMap != "" {
				fmt.Println("Need at least 2 characters")
				os.Exit(0)
			}

			for _, imagePath := range args {
				if err := convertImage(imagePath); err != nil {
					return err
				}
			}
			return nil
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

func convertImage(imagePath string) error {

	pic, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Unable to open file: %v\n", err)
		os.Exit(0)
	}
	defer pic.Close()

	imData, _, err := image.Decode(pic)
	if err != nil {
		fmt.Printf("Error decoding file: %v\n", err)
		os.Exit(0)
	}

	imgSet, err := imgMani.ConvertToAsciiPixels(imData, dimensions)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}

	var asciiSet [][]imgMani.AsciiChar
	asciiSet = imgMani.ConvertToAscii(imgSet, negative, colored, compl, customMap)

	var ascii []string
	ascii = flattenAscii(asciiSet, colored)

	// Save art before printing it, if flag is passed
	if savePath != "" {
		if err := saveAsciiArt(asciiSet, imagePath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(0)
		}
	}

	for _, line := range ascii {
		fmt.Println(line)
	}

	return nil
}

func saveAsciiArt(asciiSet [][]imgMani.AsciiChar, imagePath string) error {
	// To make sure uncolored ascii art is the one saved
	saveAscii := flattenAscii(asciiSet, false)

	saveFileName, err := createSaveFileName(imagePath)
	if err != nil {
		return err
	}

	savePathLastChar := string(savePath[len(savePath)-1])

	// Check if path is closed with appropriate path separator
	if savePathLastChar != string(os.PathSeparator) {
		if checkOS() == "linux" {
			savePath += "/"
		} else if checkOS() == "windows" {
			savePath += "\\"
		} else {
			return fmt.Errorf("Path not identified. OS isn't supported")
		}
	}

	// If path exists
	if _, err := os.Stat(savePath); !os.IsNotExist(err) {
		return ioutil.WriteFile(savePath+saveFileName, []byte(strings.Join(saveAscii, "\n")), 0777)
	} else {
		return fmt.Errorf("Save path does not exist.")
	}
}

func createSaveFileName(imagePath string) (string, error) {
	fileInfo, err := os.Stat(imagePath)
	if err != nil {
		return "", fmt.Errorf("Can't read file info for saving ascii art.")
	}
	currName := fileInfo.Name()
	extension := filepath.Ext(imagePath)
	newName := currName[:len(currName)-len(extension)]

	// Something like myImage-jpeg-ascii-art.txt
	return newName + "-" + extension[1:] + "-ascii-art.txt", nil
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with the colors from original image\n(Can work with the -n flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&compl, "complex", "c", false, "Display ascii characters in a larger range\nMay result in higher quality\n")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length\ne.g. -d 100,30 (defaults to terminal height)\n")
	rootCmd.PersistentFlags().BoolVarP(&formatsTrue, "formats", "f", false, "Display supported image formats\n")
	rootCmd.PersistentFlags().StringVarP(&customMap, "map", "m", "", "Give custom ascii characters to map against\nOrdered from darkest to lightest\ne.g. -m \" .-+#@\" (Quotation marks excluded from map)\n(Cancels --complex flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors\n(Can work with the --color flag)\n")
	rootCmd.PersistentFlags().StringVarP(&savePath, "save", "s", "", "Save ascii art in the format:\n<image-name>-<image-extension>-ascii-art.txt\nFile will be saved in passed path\n(pass . for current directory)\n")

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
