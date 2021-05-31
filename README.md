# ascii-image-converter

[![ascii-image-converter](https://snapcraft.io/ascii-image-converter/badge.svg)](https://snapcraft.io/ascii-image-converter)

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. It is cross-platform so both Windows and Linux distributions are supported.

It's also available as a package to be used in Go applications.

Image formats currently supported:
* JPEG/JPG
* PNG
* BMP
* WEBP
* TIFF/TIF

## Table of Contents

-  [Installation](#installation)
	*  [Snap](#snap)
	*  [Go](#go)
	*  [Linux (binaries)](#linux)
	*  [Windows (binaries)](#windows)
-  [CLI Usage](#cli-usage)
	*  [Flags](#flags)
-  [Library Usage](#library-usage)
-  [Contributing](#contributing)
-  [Packages used](#packages-used)
-  [License](#license)

## Installation

### Snap

You can download through snap. However, the snap will not have access to hidden images and images outside the $HOME directory.

```
sudo snap install ascii-image-converter
```
Visit [the app's snap store listing](https://snapcraft.io/ascii-image-converter) for instructions regarding enabling snapd on your distribution.

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

The basic usage for converting an image into ascii art is as follows. You can also supply multiple image paths and urls.

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

Image from URL:

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/url.gif">
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

#### --complex OR -c

Print the image with a wider array of ascii characters for more detailed lighting density. Sometimes improves accuracy.
```
ascii-image-converter [image paths/urls] -c
# Or
ascii-image-converter [image paths/urls] --complex
```

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
ascii-image-converter [image paths/urls] -d 100,30
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

#### --save OR -s

Note: Don't immediately append another flag with -s

Save ascii art in the format `<image-name>.<image-extension>-ascii-art.txt` in the directory path passed to the flag.

Example for current directory:

```
ascii-image-converter [image paths/urls] --save .
# Or
ascii-image-converter [image paths/urls] -s .
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/save.gif">
</p>

#### --formats OR -f

Display supported image formats.

```
ascii-image-converter [image paths/urls] --formats
# Or
ascii-image-converter [image paths/urls] -f
```

#### --flipX OR -x

Flip the ascii art horizontally on the terminal.

```
ascii-image-converter [image paths/urls] --flipX
# Or
ascii-image-converter [image paths/urls] -x
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/flipx.gif">
</p>

#### --flipY OR -y
Flip the ascii art vertically on the terminal.

```
ascii-image-converter [image paths/urls] --flipY
# Or
ascii-image-converter [image paths/urls] -y
```

<p align="center">
  <img src="https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_gifs/flipy.gif">
</p>

<br>

You can combine flags as well. Following command outputs colored and negative ascii art, flips ascii art horizontally and vertically, with fixed 100 by 30 character dimensions, custom defined ascii characters " .-=+#@" and saves the output in current directory as well.

```
ascii-image-converter [image paths/urls] -Cnxyd 100,30 -m " .-=+#@" -s ./
```
<br>

## Library Usage

First import the library with:
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
	// If image is in current directory. This can also be a URL to an image.
	imagePath := "myImage.jpeg"

	flags := aic_package.DefaultFlags()

	// This part is optional. You can directly pass flags variable to ConvertImage() if you wish.
	// For clarity, all flags are covered in this example.
	flags["complex"] = true  // Use complex character set
	flags["dimensions"] = []int{50, 25} // 50 by 25 ascii art size
	flags["savePath"] = "."  // Saves to current directory
	flags["negative"] = true  // Ascii art will have negative color-depth
	flags["colored"] = true  // Keep colors from original image
	flags["customMap"] = " .-=+#@"  // Starting from darkest to brightest shades. This overrites "complex" flag
	flags["flipX"] = true  // Flips ascii art horizontally
	flags["flipY"] = true  // Flips ascii art vertically
	
	// Return ascii art as a single string
	asciiArt, err := aic_package.ConvertImage(imagePath, flags)
	if err != nil {
		fmt.Println(err)
	}
  
	fmt.Printf("%v\n", asciiArt)
}
```

<br>

## Contributing

You can fork the project and implement any changes you want for a pull request. However, for major changes, please open an issue first to discuss what you would like to implement.

## Packages used

[github.com/spf13/cobra](https://github.com/spf13/cobra)

[github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)

[github.com/nathan-fiscaletti/consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go)

[github.com/nfnt/resize](https://github.com/nfnt/resize)

[github.com/gookit/color](https://github.com/gookit/color)

[github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator)

## License

[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)
