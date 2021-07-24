

# ascii-image-converter

[![release version](https://img.shields.io/github/v/release/TheZoraiz/ascii-image-converter?label=Latest%20Version)](https://github.com/TheZoraiz/ascii-image-converter/releases/latest)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)
[![ascii-image-converter-lang](https://img.shields.io/badge/Language-Go-blue)](https://golang.org/) 
![Github All Releases](https://img.shields.io/github/downloads/TheZoraiz/ascii-image-converter/total?color=brightgreen&label=Release%20Downloads)
[![ascii-image-converter](https://snapcraft.io/ascii-image-converter/badge.svg)](https://snapcraft.io/ascii-image-converter)  

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. It is cross-platform so both Windows and Linux distributions are supported. GIFs are now experimentally supported as well.

It's also available as a package to be used in Go applications.

Input formats currently supported:
* JPEG/JPG
* PNG
* BMP
* WEBP
* TIFF/TIF
* GIF (Experimental)

## Table of Contents

-  [Installation](#installation)
	*  [Ubuntu / Ubuntu-based](#ubuntu-or-ubuntu-based-distros)
	*  [Snap](#snap)
	*  [Go](#go)
	*  [Linux (binaries)](#linux)
	*  [Windows (binaries)](#windows)
-  [CLI Usage](#cli-usage)
	*  [Flags](#flags)
-  [Library Usage](#library-usage)
-  [Contributing](#contributing)
-  [Packages Used](#packages-used)
-  [License](#license)

## Installation

###  Ubuntu or Ubuntu-based Distros

Execute the following commands in order:

```
echo 'deb [trusted=yes] https://apt.fury.io/ascii-image-converter/ /' | sudo tee /etc/apt/sources.list.d/ascii-image-converter.list
```
```
sudo apt update
```
```
sudo apt install -y ascii-image-converter
```

### Snap 

You can download through snap.

Note: The snap will not have access to hidden images and images outside the $HOME directory. This includes write access for saving ascii images and text files as well.

```
sudo snap install ascii-image-converter
```
Visit [the app's snap store listing](https://snapcraft.io/ascii-image-converter) for instructions regarding enabling snapd on your distribution.


[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/ascii-image-converter)

<hr>

### Go

For installing through Go
```
go install github.com/TheZoraiz/ascii-image-converter@latest
```
<hr>
For physically installing the binaries, follow the steps with respect to your OS.

### Linux

Download the archive for your distribution's architecture [here](https://github.com/TheZoraiz/ascii-image-converter/releases/latest), extract it, and open the extracted directory.

Now, open a terminal in the same directory and execute this command:

```
sudo cp ascii-image-converter /usr/local/bin/
```
Now you can use ascii-image-converter in the terminal. Execute "ascii-image-converter -h" for more details.

### Windows

You will need to set an Environment Variable to the folder the ascii-image-converter.exe executable is placed in to be able to use it in the command prompt. Follow the instructions in case of confusion:

Download the archive for your Windows architecture [here](https://github.com/TheZoraiz/ascii-image-converter/releases/latest), extract it, and open the extracted folder. Now, copy the folder path from the top of the file explorer and follow these instructions:
* In Search, search for and then select: System (Control Panel)
* Click the Advanced System settings link.
* Click Environment Variables. In the section User Variables find the Path environment variable and select it. Click "Edit".
* In the Edit Environment Variable window, click "New" and then paste the path of the folder that you copied initially.
* Click "Ok" on all open windows.

Now, restart any open command prompt and execute "ascii-image-converter -h" for more details.

<br>

## CLI Usage

Note: Decrease font size or increase terminal width (like zooming out) for maximum quality ascii art

The basic usage for converting an image into ascii art is as follows. You can also supply multiple image paths and urls as well as a GIF.

```
ascii-image-converter [image paths/urls]
```
Example:
```
ascii-image-converter myImage.jpeg
```
<br>
Single image:

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/base.gif">
</p>

Multiple images:

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/all.gif">
</p>

GIF:

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/gif-example.gif">
</p>

### Flags

#### --color OR -C

Display ascii art with the colors from original image. Works with the --negative flag as well.

```
ascii-image-converter [image paths/urls] -C
# Or
ascii-image-converter [image paths/urls] --color
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/color.gif">
</p>

#### --dimensions OR -d

Note: Don't immediately append another flag with -d

Set the width and height for ascii art in CHARACTER lengths.
```
ascii-image-converter [image paths/urls] -d <width>,<height>
# Or
ascii-image-converter [image paths/urls] --dimensions <width>,<height>
```
Example:
```
ascii-image-converter [image paths/urls] -d 60,30
```
<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/dimensions.gif">
</p>

#### --map OR -m

Note: Don't immediately append another flag with -m

Pass a string of your own ascii characters to map against. Passed characters must start from darkest character and end with lightest. There is no limit to number of characters.

Empty spaces can be passed if string is passed inside quotation marks. You can use both single or double quote for quotation marks. For repeating quotation mark inside string, append it with \ (such as  \\").
  
```
ascii-image-converter [image paths/urls] -m "<string-of-characters>"
# Or
ascii-image-converter [image paths/urls] --map "<string-of-characters>"
```
Following example contains 7 depths of lighting.
```
ascii-image-converter [image paths/urls] -m " .-=+#@"
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/map.gif">
</p>

#### --negative OR -n

Display ascii art in negative colors. Works with both uncolored and colored text from --color flag.

```
ascii-image-converter [image paths/urls] -n
# Or
ascii-image-converter [image paths/urls] --negative
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/negative.gif">
</p>

#### --complex OR -c

Print the image with a wider array of ascii characters for more detailed lighting density. Sometimes improves accuracy.
```
ascii-image-converter [image paths/urls] -c
# Or
ascii-image-converter [image paths/urls] --complex
```

#### --full OR -f

Print ascii art that fits the terminal width while maintaining aspect ratio.
```
ascii-image-converter [image paths/urls] -f
# Or
ascii-image-converter [image paths/urls] --full
```

#### --flipX OR -x

Flip the ascii art horizontally on the terminal.

```
ascii-image-converter [image paths/urls] --flipX
# Or
ascii-image-converter [image paths/urls] -x
```

#### --flipY OR -y
Flip the ascii art vertically on the terminal.

```
ascii-image-converter [image paths/urls] --flipY
# Or
ascii-image-converter [image paths/urls] -y
```



#### --save-img OR -s

Note: Don't immediately append another flag with -s

Saves the ascii as a PNG image with the name `<image-name>-ascii-art.png` in the directory path passed to the flag. Can work with both --color and --negative flag.

Example for current directory:

```
ascii-image-converter [image paths/urls] --save-img .
# Or
ascii-image-converter [image paths/urls] -s .
```

#### --save-txt

Similar to --save-img but it creates a TXT file with the name `<image-name>-ascii-art.txt` in the directory path passed to the flag. Only saves uncolored text.

Example for current directory:

```
ascii-image-converter [image paths/urls] --save-txt .
```

#### --save-gif

Note: This is an experimental feature and may not result in the finest quality GIFs, because all GIFs still aren't supported by ascii-image-converter.

Saves the passed GIF as an ascii art GIF with the name `<image-name>-ascii-art.gif` in the directory path passed to the flag.

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/save.gif">
</p>

#### --formats

Display supported input formats.

```
ascii-image-converter --formats
```

<br>

You can combine flags as well. Following command outputs colored and negative ascii art, flips ascii art horizontally and vertically, with fixed 60 by 30 character dimensions, custom defined ascii characters " .-=+#@" and saves a generated image and .txt file in current directory as well.

```
ascii-image-converter [image paths/urls] -Cnxyd 60,30 -m " .-=+#@" -s . --save-txt .
```
<br>

## Library Usage

Note: The library may throw errors during Go tests due to some unresolved bugs with the [consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go) package (Only during tests, not main program execution). Furthermore, GIF conversion is not advised as it isn't fully library-compatible yet.

First, install the library with:
```
go get github.com/TheZoraiz/ascii-image-converter/aic_package
```

The library is to be used as follows:

```go
package main

import (
	"fmt"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func main() {
	// If file is in current directory. This can also be a URL to an image or gif.
	filePath := "myImage.jpeg"

	flags := aic_package.DefaultFlags()

	// This part is optional.
	// You can directly pass default flags variable to Convert() if you wish.
	// For clarity, all flags are covered in this example, but you can use specific ones.
	flags.Complex = true  // Use complex character set
	flags.Dimensions = []int{50, 25} // 50 by 25 ascii art size
	flags.SaveTxtPath = "."  // Save generated text in same directory
	flags.SaveImagePath = "."  // Save generated PNG image in same directory
	flags.SaveGifPath = "." // If gif was provided, save ascii art gif in same directory
	flags.Negative = true  // Ascii art will have negative color-depth
	flags.Colored = true  // Keep colors from original image
	flags.CustomMap = " .-=+#@"  // Starting from darkest to brightest shades. This overrites "complex" flag
	flags.FlipX = true  // Flips ascii art horizontally
	flags.FlipY = true  // Flips ascii art vertically
	flags.Full = true  // Display ascii art that fills the terminal width
	
	// For an image
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)

	// GIF CONVERSION IS AN EXPERIMENTAL FEATURE
	// For a gif. This function may run infinitely, depending on the gif
	// Work needs to be done on gif conversion to be more library-compatible
	_, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}
}
```

<br>

## Contributing

You can fork the project and implement any changes you want for a pull request. However, for major changes, please open an issue first to discuss what you would like to implement.

## Packages Used

[github.com/spf13/cobra](https://github.com/spf13/cobra)

[github.com/fogleman/gg](https://github.com/fogleman/gg)

[github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)

[github.com/nathan-fiscaletti/consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go)

[github.com/nfnt/resize](https://github.com/nfnt/resize)

[github.com/gookit/color](https://github.com/gookit/color)

[github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator)

## License

[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)
