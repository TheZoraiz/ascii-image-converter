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
	"path"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/TheZoraiz/ascii-image-converter/aic_package/winsize"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Flags
	cfgFile       string
	complex       bool
	dimensions    []int
	width         int
	height        int
	saveTxtPath   string
	saveImagePath string
	saveGifPath   string
	negative      bool
	formatsTrue   bool
	colored       bool
	grayscale     bool
	customMap     string
	flipX         bool
	flipY         bool
	full          bool
	fontFile      string

	// Root commands
	rootCmd = &cobra.Command{
		Use:     "ascii-image-converter [image paths/urls]",
		Short:   "Converts images and gifs into ascii art",
		Version: "1.5.0",
		Long:    "This tool converts images into ascii art and prints them on the terminal.\nFurther configuration can be managed with flags.",

		// Not RunE since help text is getting larger and seeing it for every error impacts user experience
		Run: func(cmd *cobra.Command, args []string) {

			if checkInputAndFlags(args) {
				return
			}

			flags := aic_package.Flags{
				Complex:       complex,
				Dimensions:    dimensions,
				Width:         width,
				Height:        height,
				SaveTxtPath:   saveTxtPath,
				SaveImagePath: saveImagePath,
				SaveGifPath:   saveGifPath,
				Negative:      negative,
				Colored:       colored,
				Grayscale:     grayscale,
				CustomMap:     customMap,
				FlipX:         flipX,
				FlipY:         flipY,
				Full:          full,
				FontFilePath:  fontFile,
			}

			for _, imagePath := range args {

				if asciiArt, err := aic_package.Convert(imagePath, flags); err == nil {
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

// Check input and flags for any errors or interruptons
func checkInputAndFlags(args []string) bool {

	gifCount := 0
	gifPresent := false
	nonGifPresent := false
	for _, arg := range args {
		extension := path.Ext(arg)

		if extension == ".gif" {
			gifPresent = true
			gifCount++
		} else {
			nonGifPresent = true
		}
	}

	if gifPresent && nonGifPresent {
		fmt.Printf("Error: There are other inputs along with GIFs\nDue to the potential looping nature of GIFs, non-GIFs must not be supplied alongside\n\n")
		return true
	}

	if gifCount > 1 {
		fmt.Printf("Error: There are multiple GIFs supplied\nDue to the potential looping nature of GIFs, only one GIF per command is supported\n\n")
		return true
	}

	if formatsTrue {
		fmt.Printf("Supported input formats:\n\n" +
			"JPEG/JPG\n" +
			"PNG\n" +
			"WEBP\n" +
			"BMP\n" +
			"TIFF/TIF\n" +
			"GIF (Experimental)\n\n")
		return true
	}

	if len(args) < 1 {
		fmt.Printf("Error: Need at least 1 input path/url\nUse the -h flag for more info\n\n")
		return true
	}

	if customMap != "" && len(customMap) < 2 {
		fmt.Printf("Need at least 2 characters for --map flag\n\n")
		return true
	}

	if dimensions != nil {

		numberOfDimensions := len(dimensions)
		if numberOfDimensions != 2 {
			fmt.Printf("Error: requires 2 dimensions, got %v\n\n", numberOfDimensions)
			return true
		}

		if dimensions[0] < 1 || dimensions[1] < 1 {
			fmt.Printf("Error: invalid values for dimensions\n\n")
			return true
		}

		defaultTermWidth, _, err := winsize.GetTerminalSize()
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			return true
		}

		defaultTermWidth -= 1
		if dimensions[0] > defaultTermWidth {
			fmt.Printf("Error: set width must be lower than terminal width\n\n")
			return true
		}
	}

	if width != 0 || height != 0 {

		if width != 0 && height != 0 {
			fmt.Printf("Error: both --width and --height can't be set. Use --dimensions instead\n\n")
			return true
		} else {

			defaultTermWidth, _, err := winsize.GetTerminalSize()
			if err != nil {
				fmt.Printf("Error: %v\n\n", err)
				return true
			}

			// Check if set width exceeds terminal
			defaultTermWidth -= 1
			if width > defaultTermWidth {
				fmt.Printf("Error: set width must be lower than terminal width\n\n")
				return true
			}

			if width < 0 {
				fmt.Printf("Error: invalid value for width\n\n")
				return true
			}

			if height < 0 {
				fmt.Printf("Error: invalid value for height\n\n")
				return true
			}

		}

	}

	return false
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

	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ascii-image-converter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with original colors\n(Can work with the --negative flag)\n(Overrides --grayscale flag)\n")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length\ne.g. -d 60,30 (defaults to terminal height)\n(Overrides --width and --height flags)\n")
	rootCmd.PersistentFlags().IntVarP(&width, "width", "W", 0, "Set width for ascii art in CHARACTER length\nHeight is kept to aspect ratio\ne.g. -W 60\n")
	rootCmd.PersistentFlags().IntVarP(&height, "height", "H", 0, "Set height for ascii art in CHARACTER length\nWidth is kept to aspect ratio\ne.g. -H 60\n")
	rootCmd.PersistentFlags().StringVarP(&customMap, "map", "m", "", "Give custom ascii characters to map against\nOrdered from darkest to lightest\ne.g. -m \" .-+#@\" (Quotation marks excluded from map)\n(Overrides --complex flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&grayscale, "grayscale", "g", false, "Display grayscale ascii art\n(Can work with --negative flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&complex, "complex", "c", false, "Display ascii characters in a larger range\nMay result in higher quality\n")
	rootCmd.PersistentFlags().BoolVarP(&full, "full", "f", false, "Use largest dimensions for ascii art\nthat fill the terminal width\n(Overrides --dimensions, --width and --height flags)\n")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors\n")
	rootCmd.PersistentFlags().BoolVarP(&flipX, "flipX", "x", false, "Flip ascii art horizontally\n")
	rootCmd.PersistentFlags().BoolVarP(&flipY, "flipY", "y", false, "Flip ascii art vertically\n")
	rootCmd.PersistentFlags().StringVarP(&saveImagePath, "save-img", "s", "", "Save ascii art as a .png file\nFormat: <image-name>-ascii-art.png\nImage will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveTxtPath, "save-txt", "", "Save ascii art as a .txt file\nFormat: <image-name>-ascii-art.txt\nFile will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveGifPath, "save-gif", "", "(Experimental)\nIf input is a gif, save it as a .gif file\nFormat: <image-name>-ascii-art.gif\nGif will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&fontFile, "font", "", "Set font for --save-img and --save-gif flags\nPass file path to font .ttf file\ne.g. --font ./RobotoMono-Regular.ttf\n(Defaults to embedded Hack-Regular)\n")
	rootCmd.PersistentFlags().BoolVar(&formatsTrue, "formats", false, "Display supported input formats\n")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for "+rootCmd.Name()+"\n")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Version for "+rootCmd.Name())

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
