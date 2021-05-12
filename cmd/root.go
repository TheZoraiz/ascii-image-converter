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

	imgMani "github.com/TheZoraiz/ascii-image-converter/image_manipulation"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	simple     bool
	dimensions []int
	save       bool

	rootCmd = &cobra.Command{
		Use:   "ascii-image-converter [image path]",
		Short: "Converts images into ascii format",
		Long:  `ascii-image-converter converts images into ascii format and prints them onto the terminal window. Further configuration can be managed with flags`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) > 1 {
				cmd.Help()
				os.Exit(1)
			} else if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}

			var imagePath string

			isComplex, err := cmd.Flags().GetBool("complex")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			save, err := cmd.Flags().GetBool("save")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			dims, err := cmd.Flags().GetIntSlice("dimensions")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if len(dims) > 0 && len(dims) != 2 {
				fmt.Println("Error: Need two dimensions\n")
				cmd.Help()
				os.Exit(1)
			}

			imagePath = args[0]

			convertPicture(imagePath, isComplex, dims, save)
		},
	}
)

func convertPicture(imagePath string, isComplex bool, dimensions []int, save bool) {

	pic, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer pic.Close()

	imData, _, err := image.Decode(pic)
	if err != nil {
		panic(err)
	}

	imgSet := imgMani.ConvertToTerminalSizedSlices(imData, dimensions)

	if isComplex {
		printDetailedAscii(imgSet, save)
	} else {
		printSimpleAscii(imgSet, save)
	}
}

func printDetailedAscii(imgSet [][]uint32, save bool) {
	final := imgMani.ConvertToAsciiDetailed(imgSet)

	var temp string
	for i := 0; i < len(final); i++ {
		for j := 0; j < len(final[i]); j++ {
			temp += fmt.Sprintf("%s", string(final[i][j]))
			fmt.Printf("%s", string(final[i][j]))
		}
		temp += "\n"
		fmt.Println()
	}
	if save {
		ioutil.WriteFile("ascii-image.txt", []byte(temp), 0777)
	}
}

func printSimpleAscii(imgSet [][]uint32, save bool) {
	final := imgMani.ConvertToAsciiSimple(imgSet)

	var temp string
	for i := 0; i < len(final); i++ {
		for j := 0; j < len(final[i]); j++ {
			temp += fmt.Sprintf("%s", string(final[i][j]))
			fmt.Printf("%s", string(final[i][j]))
		}
		temp += "\n"
		fmt.Println()
	}
	if save {
		ioutil.WriteFile("ascii-image.txt", []byte(temp), 0777)
	}
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
	rootCmd.PersistentFlags().BoolVarP(&simple, "complex", "c", false, "Prints ascii characters in a larger range, may result in higher quality")
	rootCmd.PersistentFlags().IntSliceVarP(&dimensions, "dimensions", "d", nil, "Set width and height for ascii art in CHARACTER length e.g. 100,30 (defaults to terminal size)")
	rootCmd.PersistentFlags().BoolVarP(&save, "save", "S", false, "Save ascii text in current directory in an ascii-image.txt file")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
