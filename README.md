# ascii-image-converter

[![release-version](https://img.shields.io/github/v/release/TheZoraiz/ascii-image-converter?label=Latest%20Version)](https://github.com/TheZoraiz/ascii-image-converter/releases/latest)
[![license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)
[![language](https://img.shields.io/badge/Language-Go-blue)](https://golang.org/)
![release-downloads](https://img.shields.io/github/downloads/TheZoraiz/ascii-image-converter/total?color=1d872d&label=Release%20Downloads)
[![ascii-image-converter-snap](https://snapcraft.io/ascii-image-converter/badge.svg)](https://snapcraft.io/ascii-image-converter)

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. Available on Windows, Linux and macOS.

Now supports braille art!

Input formats currently supported:
* JPEG/JPG
* PNG
* BMP
* WEBP
* TIFF/TIF
* GIF

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/all.gif">
</p>

## Table of Contents

-  [Installation](#installation)
	*  [Debian / Ubuntu-based](#debian-or-ubuntu-based-distros)
	*  [Homebrew](#homebrew)
	*  [AUR](#aur)
	*  [Scoop](#scoop)
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

### Debian or Ubuntu-based Distros

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
<br>

To remove the package source (which means you won't be getting any further updates), execute this command:

```
sudo rm -v /etc/apt/sources.list.d/ascii-image-converter.list
```

<hr>

### Homebrew

Installation with homebrew is available for both Linux and macOS.
```
brew install TheZoraiz/ascii-image-converter/ascii-image-converter
```
[Link to homebrew repository](https://github.com/TheZoraiz/homebrew-ascii-image-converter)

<hr>

### AUR

The AUR repo is maintained by [magnus-tesshu](https://aur.archlinux.org/account/magnus-tesshu)

Standard way:
```
git clone https://aur.archlinux.org/ascii-image-converter-git.git
```
```
cd ascii-image-converter-git/
```
```
makepkg -si
```
AUR helper:
```
<aur-helper> -S ascii-image-converter-git
```
<hr>

### Scoop

The scoop manifest is maintained by [brian6932](https://github.com/brian6932)

```
scoop install ascii-image-converter
```

<hr>

### Snap


> **Note:** The snap will not have access to hidden files and files outside the $HOME directory. This includes write access for saving ascii art as well.

```
sudo snap install ascii-image-converter
```
Visit [the app's snap store listing](https://snapcraft.io/ascii-image-converter) for instructions regarding enabling snapd on your distribution.


[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/ascii-image-converter)

<hr>

### Go

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
Now you can use ascii-image-converter in the terminal. Execute `ascii-image-converter -h` for more details.

### Windows

You will need to set an Environment Variable to the folder the ascii-image-converter.exe executable is placed in to be able to use it in the command prompt. Follow the instructions in case of confusion:

Download the archive for your Windows architecture [here](https://github.com/TheZoraiz/ascii-image-converter/releases/latest), extract it, and open the extracted folder. Now, copy the folder path from the top of the file explorer and follow these instructions:
* In Search, search for and then select: Advanced System Settings
* Click Environment Variables. In the section User Variables find the Path environment variable and select it. Click "Edit".
* In the Edit Environment Variable window, click "New" and then paste the path of the folder that you copied initially.
* Click "Ok" on all open windows.

Now, restart any open command prompt and execute `ascii-image-converter -h` for more details.

<br>

## CLI Usage

> **Note:** Decrease font size or increase terminal width (like zooming out) for maximum quality ascii art

The basic usage for converting an image into ascii art is as follows. You can also supply multiple image paths and urls as well as a GIF.

```
ascii-image-converter [image paths/urls]
```
Example:
```
ascii-image-converter myImage.jpeg
```

> **Note:** Piped binary input is also supported
> ```
> cat myImage.png | ascii-image-converter -
> ```


### Flags

#### --color OR -C

> **Note:** Your terminal must support 24-bit or 8-bit colors for appropriate results. If 24-bit colors aren't supported, 8-bit color escape codes will be used

Display ascii art with the colors from original image.

```
ascii-image-converter [image paths/urls] -C
# Or
ascii-image-converter [image paths/urls] --color
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/color.gif">
</p>

#### --braille OR -b

> **Note:** Braille pattern display heavily depends on which terminal or font you're using. In windows, try changing the font from command prompt properties if braille characters don't display

Use braille characters instead of ascii. For this flag, your terminal must support braille patters (UTF-8) properly. Otherwise, you may encounter problems with colored or even uncolored braille art.
```
ascii-image-converter [image paths/urls] -b
# Or
ascii-image-converter [image paths/urls] --braille
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/braille.gif">
</p>

#### --threshold

Set threshold value to compare for braille art when converting each pixel into a dot. Value must be between 0 and 255.

Example:
```
ascii-image-converter [image paths/urls] -b --threshold 170
```

#### --dither

Apply dithering on image to make braille art more visible. Since braille dots can only be on or off, dithering images makes them more visible in braille art.

Example:
```
ascii-image-converter [image paths/urls] -b --dither
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/dither.gif">
</p>

#### --color-bg

If any of the coloring flags is passed, this flag will transfer its color to each character's background. instead of foreground. However, this option isn't available for `--save-img` and `--save-gif`
```
ascii-image-converter [image paths/urls] -C --color-bg
```

#### --dimensions OR -d

> **Note:** Don't immediately append another flag with -d

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

#### --width OR -W

> **Note:** Don't immediately append another flag with -W

Set width of ascii art. Height is calculated according to aspect ratio.
```
ascii-image-converter [image paths/urls] -W <width>
# Or
ascii-image-converter [image paths/urls] --width <width>
```
Example:
```
ascii-image-converter [image paths/urls] -W 60
```

#### --height OR -H

> **Note:** Don't immediately append another flag with -H

Set height of ascii art. Width is calculated according to aspect ratio.
```
ascii-image-converter [image paths/urls] -H <height>
# Or
ascii-image-converter [image paths/urls] --height <height>
```
Example:
```
ascii-image-converter [image paths/urls] -H 60
```

#### --map OR -m

> **Note:** Don't immediately append another flag with -m

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

#### --grayscale OR -g

Display ascii art in grayscale colors. This is the same as --color flag, except each character will be encoded with a grayscale RGB value.

```
ascii-image-converter [image paths/urls] -g
# Or
ascii-image-converter [image paths/urls] --grayscale
```

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

> **Note:** Don't immediately append another flag with -s

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

> **Note:** This is an experimental feature and may not result in the finest quality GIFs, because all GIFs still aren't supported by ascii-image-converter.

Saves the passed GIF as an ascii art GIF with the name `<image-name>-ascii-art.gif` in the directory path passed to the flag.

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/save.gif">
</p>

#### --save-html

Similar to --save-img but it creates a HTML file with the name `<image-name>-ascii-art.html` in the directory path passed to the flag. Can save colored text, that can be copied from open html file.

Example for current directory:

```
ascii-image-converter [image paths/urls] --save-html .
```

#### --save-bg

> **Note:** This flag will be ignored if `--save-img` or `--save-gif` flags are not set

This flag takes an RGBA value that sets the background color in saved png and gif files. The fourth value (alpha value) is the measure of background opacity ranging between 0 and 100.

```
ascii-image-converter [image paths/urls] -s . --save-bg 255,255,255,100 # For white background
```

#### --font

> **Note:** This flag will be ignored if `--save-img` or `--save-gif` flags are not set

This flag takes path to a font .ttf file that will be used to set font in saved png or gif files.

```
ascii-image-converter [image paths/urls] -s . --font /path/to/font-file.ttf
```

#### --font-color

This flag takes an RGB value that sets the font color in saved png and gif files as well as displayed ascii art in terminal.

```
ascii-image-converter [image paths/urls] -s . --font-color 0,0,0 # For black font color
```

#### --only-save

Don't print ascii art on the terminal if some saving flag is passed.

```
ascii-image-converter [image paths/urls] -s . --only-save
```

#### --formats

Display supported input formats.

```
ascii-image-converter --formats
```

<br>

## Library Usage

> **Note:** The library may throw errors during Go tests due to some unresolved bugs with the [consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go) package (Only during tests, not main program execution).

First, install the library with:
```
go get -u github.com/TheZoraiz/ascii-image-converter/aic_package
```

For an image:

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
	// You can directly pass default flags variable to aic_package.Convert() if you wish.
	// There are more flags, but these are the ones shown for demonstration
	flags.Dimensions = []int{50, 25}
	flags.Colored = true
	flags.SaveTxtPath = "."
	flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	flags.FontFilePath = "./RobotoMono-Regular.ttf" // If file is in current directory
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}

	// Note: For environments where a terminal isn't available (such as web servers), you MUST
	// specify atleast one of flags.Width, flags.Height or flags.Dimensions

	// Conversion for an image
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}
```
<br>

> **Note:** GIF conversion is not advised as the function may run infinitely, depending on the GIF. More work needs to be done on this to make it more library-compatible.

For a GIF:

```go
package main

import (
	"fmt"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func main() {
	filePath = "myGif.gif"

	flags := aic_package.DefaultFlags()

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

[github.com/disintegration/imaging](https://github.com/disintegration/imaging)

[github.com/gookit/color](https://github.com/gookit/color)

[github.com/makeworld-the-better-one/dither](https://github.com/makeworld-the-better-one/dither)

## License

[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)