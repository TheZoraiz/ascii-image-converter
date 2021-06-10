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
	"os"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Flags
	cfgFile       string
	complex       bool
	dimensions    []int
	saveTxtPath   string
	saveImagePath string
	negative      bool
	formatsTrue   bool
	colored       bool
	customMap     string
	flipX         bool
	flipY         bool

	// Root commands
	rootCmd = &cobra.Command{
		Use:     "ascii-image-converter [image paths/urls]",
		Short:   "Converts images into ascii art",
		Version: "1.3.4",
		Long:    "This tool converts images into ascii art and prints them on the terminal.\nFurther configuration can be managed with flags.",

		// Not RunE since help text is getting larger and seeing it for every error impacts user experience
		Run: func(cmd *cobra.Command, args []string) {

			if formatsTrue {
				fmt.Printf("Supported image formats: JPEG/JPG, PNG, WEBP, BMP, TIFF/TIF\n\n")
				return
			}

			if len(args) < 1 {
				fmt.Printf("Error: Need at least 1 image path/url\nUse the -h flag for more info\n\n")
				return
			}

			if customMap != "" && len(customMap) < 2 {
				fmt.Printf("Need at least 2 characters for --map flag\n\n")
				return
			}

			numberOfDimensions := len(dimensions)
			if dimensions != nil && numberOfDimensions != 2 {
				fmt.Printf("Error: requires 2 dimensions, got %v\n\n", numberOfDimensions)
				return
			}

			if dimensions != nil && (dimensions[0] < 1 || dimensions[1] < 1) {
				fmt.Printf("Error: invalid values for dimensions\n\n")
				return
			}

			flags := map[string]interface{}{
				"complex":       complex,
				"dimensions":    dimensions,
				"saveTxtPath":   saveTxtPath,
				"saveImagePath": saveImagePath,
				"negative":      negative,
				"colored":       colored,
				"customMap":     customMap,
				"flipX":         flipX,
				"flipY":         flipY,
			}

			for _, imagePath := range args {

				if asciiArt, err := aic_package.ConvertImage(imagePath, flags); err == nil {
					fmt.Printf("%s", asciiArt)
				} else {
					fmt.Printf("Error: %v\n", err)

					// Because this error will then be thrown for every image path/url passed
					// if save path is invalid
					if err.Error()[:15] == "can't save file" {
						fmt.Println()
						return
					}
				}
				fmt.Println()
			}
		},
	}
)

// Cobra configuration from here on

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().SortFlags = false

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with the colors from original image\n(Can work with the --negative flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&complex, "complex", "c", false, "Display ascii characters in a larger range\nMay result in higher quality\n")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length\ne.g. -d 100,30 (defaults to terminal height)\n")
	rootCmd.PersistentFlags().BoolVarP(&formatsTrue, "formats", "f", false, "Display supported image formats\n")
	rootCmd.PersistentFlags().StringVarP(&customMap, "map", "m", "", "Give custom ascii characters to map against\nOrdered from darkest to lightest\ne.g. -m \" .-+#@\" (Quotation marks excluded from map)\n(Cancels --complex flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors\n(Can work with the --color flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&flipX, "flipX", "x", false, "Flip ascii art horizontally\n")
	rootCmd.PersistentFlags().BoolVarP(&flipY, "flipY", "y", false, "Flip ascii art vertically\n")

	rootCmd.PersistentFlags().StringVarP(&saveImagePath, "save-img", "s", "", "Save ascii art in a .png file\nFormat: <image-name>-ascii-art.png\nImage will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveTxtPath, "save-txt", "", "Save ascii art in the a .txt file\nFormat: <image-name>-ascii-art.txt\nFile will be saved in passed path\n(pass . for current directory)\n")

	defaultUsageTemplate := rootCmd.UsageTemplate()
	rootCmd.SetUsageTemplate(defaultUsageTemplate + "\nCopyright © 2021 Zoraiz Hassan <hzoraiz8@gmail.com>\n" +
		"Distributed under the Apache License Version 2.0 (Apache-2.0)\n" +
		"For further details, visit https://github.com/TheZoraiz/ascii-image-converter\n")
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
