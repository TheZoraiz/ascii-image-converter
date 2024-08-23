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
	width         int
	height        int
	saveTxtPath   string
	saveImagePath string
	saveHtmlPath  string
	saveGifPath   string
	negative      bool
	formatsTrue   bool
	colored       bool
	colorBg       bool
	grayscale     bool
	customMap     string
	flipX         bool
	flipY         bool
	full          bool
	fontFile      string
	fontColor     []int
	saveBgColor   []int
	braille       bool
	threshold     int
	dither        bool
	onlySave      bool

	// Root commands
	rootCmd = &cobra.Command{
		Use:     "ascii-image-converter [image paths/urls or piped stdin]",
		Short:   "Converts images and gifs into ascii art",
		Version: "1.13.1",
		Long:    "This tool converts images into ascii art and prints them on the terminal.\nFurther configuration can be managed with flags.",

		// Not RunE since help text is getting larger and seeing it for every error impacts user experience
		Run: func(cmd *cobra.Command, args []string) {

			if checkInputAndFlags(args) {
				return
			}

			flags := aic_package.Flags{
				Complex:             complex,
				Dimensions:          dimensions,
				Width:               width,
				Height:              height,
				SaveTxtPath:         saveTxtPath,
				SaveImagePath:       saveImagePath,
				SaveGifPath:         saveGifPath,
				SaveHtmlPath:        saveHtmlPath,
				Negative:            negative,
				Colored:             colored,
				CharBackgroundColor: colorBg,
				Grayscale:           grayscale,
				CustomMap:           customMap,
				FlipX:               flipX,
				FlipY:               flipY,
				Full:                full,
				FontFilePath:        fontFile,
				FontColor:           [3]int{fontColor[0], fontColor[1], fontColor[2]},
				SaveBackgroundColor: [4]int{saveBgColor[0], saveBgColor[1], saveBgColor[2], saveBgColor[3]},
				Braille:             braille,
				Threshold:           threshold,
				Dither:              dither,
				OnlySave:            onlySave,
			}

			if args[0] == "-" {
				printAscii(args[0], flags)
				return
			}

			for _, imagePath := range args {
				if err := printAscii(imagePath, flags); err != nil {
					return
				}
			}
		},
	}
)

func printAscii(imagePath string, flags aic_package.Flags) error {

	if asciiArt, err := aic_package.Convert(imagePath, flags); err == nil {
		fmt.Printf("%s", asciiArt)
	} else {
		fmt.Printf("Error: %v\n", err)

		// Because this error will then be thrown for every image path/url passed
		// if save path is invalid
		if err.Error()[:15] == "can't save file" {
			fmt.Println()
			return err
		}
	}
	if !onlySave {
		fmt.Println()
	}
	return nil
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
	rootCmd.PersistentFlags().BoolVarP(&colored, "color", "C", false, "Display ascii art with original colors\nIf 24-bit colors aren't supported, uses 8-bit\n(Inverts with --negative flag)\n(Overrides --grayscale and --font-color flags)\n")
	rootCmd.PersistentFlags().BoolVar(&colorBg, "color-bg", false, "If some color flag is passed, use that color\non character background instead of foreground\n(Inverts with --negative flag)\n(Only applicable for terminal display)\n")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length\ne.g. -d 60,30 (defaults to terminal height)\n(Overrides --width and --height flags)\n")
	rootCmd.PersistentFlags().IntVarP(&width, "width", "W", 0, "Set width for ascii art in CHARACTER length\nHeight is kept to aspect ratio\ne.g. -W 60\n")
	rootCmd.PersistentFlags().IntVarP(&height, "height", "H", 0, "Set height for ascii art in CHARACTER length\nWidth is kept to aspect ratio\ne.g. -H 60\n")
	rootCmd.PersistentFlags().StringVarP(&customMap, "map", "m", "", "Give custom ascii characters to map against\nOrdered from darkest to lightest\ne.g. -m \" .-+#@\" (Quotation marks excluded from map)\n(Overrides --complex flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&braille, "braille", "b", false, "Use braille characters instead of ascii\nTerminal must support braille patterns properly\n(Overrides --complex and --map flags)\n")
	rootCmd.PersistentFlags().IntVar(&threshold, "threshold", 0, "Threshold for braille art\nValue between 0-255 is accepted\ne.g. --threshold 170\n(Defaults to 128)\n")
	rootCmd.PersistentFlags().BoolVar(&dither, "dither", false, "Apply dithering on image for braille\nart conversion\n(Only applicable with --braille flag)\n(Negates --threshold flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&grayscale, "grayscale", "g", false, "Display grayscale ascii art\n(Inverts with --negative flag)\n(Overrides --font-color flag)\n")
	rootCmd.PersistentFlags().BoolVarP(&complex, "complex", "c", false, "Display ascii characters in a larger range\nMay result in higher quality\n")
	rootCmd.PersistentFlags().BoolVarP(&full, "full", "f", false, "Use largest dimensions for ascii art\nthat fill the terminal width\n(Overrides --dimensions, --width and --height flags)\n")
	rootCmd.PersistentFlags().BoolVarP(&negative, "negative", "n", false, "Display ascii art in negative colors\n")
	rootCmd.PersistentFlags().BoolVarP(&flipX, "flipX", "x", false, "Flip ascii art horizontally\n")
	rootCmd.PersistentFlags().BoolVarP(&flipY, "flipY", "y", false, "Flip ascii art vertically\n")
	rootCmd.PersistentFlags().StringVarP(&saveImagePath, "save-img", "s", "", "Save ascii art as a .png file\nFormat: <image-name>-ascii-art.png\nImage will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveTxtPath, "save-txt", "", "Save ascii art as a .txt file\nFormat: <image-name>-ascii-art.txt\nFile will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveGifPath, "save-gif", "", "If input is a gif, save it as a .gif file\nFormat: <gif-name>-ascii-art.gif\nGif will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().StringVar(&saveHtmlPath, "save-html", "", "Save ascii art as a .html file\nFormat: <image-name>-ascii-art.html\nFile will be saved in passed path\n(pass . for current directory)\n")
	rootCmd.PersistentFlags().IntSliceVar(&saveBgColor, "save-bg", nil, "Set background color for --save-img\nand --save-gif \nand --save-html flags\nPass an RGBA value\ne.g. --save-bg 255,255,255,100\n(Defaults to 0,0,0,100)\n")
	rootCmd.PersistentFlags().StringVar(&fontFile, "font", "", "Set font for --save-img and --save-gif flags\nPass file path to font .ttf file\ne.g. --font ./RobotoMono-Regular.ttf\n(Defaults to Hack-Regular for ascii and\n DejaVuSans-Oblique for braille)\n")
	rootCmd.PersistentFlags().IntSliceVar(&fontColor, "font-color", nil, "Set font color for terminal as well as\n--save-img and --save-gif flags\nPass an RGB value\ne.g. --font-color 0,0,0\n(Defaults to 255,255,255)\n")
	rootCmd.PersistentFlags().BoolVar(&onlySave, "only-save", false, "Don't print ascii art on terminal\nif some saving flag is passed\n")
	rootCmd.PersistentFlags().BoolVar(&formatsTrue, "formats", false, "Display supported input formats\n")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for "+rootCmd.Name()+"\n")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Version for "+rootCmd.Name())

	rootCmd.SetVersionTemplate("{{printf \"v%s\" .Version}}\n")

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
