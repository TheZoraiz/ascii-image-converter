# ascii-image-converter

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. It is cross-platform so both Windows and Linux distributions are supported.

Image formats currently supported:
* JPEG/JPG
* PNG
* BMP
* WEBP
* TIFF/TIF

## Table of Contents

-  [Example](#example-source)
-  [Installation](#installation)
	*  [Snap](#snap)
	*  [Go](#go)
	*  [Linux (binaries)](#linux)
	*  [Windows (binaries)](#windows)
-  [Usage](#usage)
	*  [Flags](#flags)
-  [Contributing](#contributing)
-  [Packages used](#packages-used)
-  [License](#license)

### Example ([Source](https://medium.com/@sean.glancy/practical-applications-of-binary-trees-3097cf663062)):

![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/tree.png)

### ASCII Art:

![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/ascii_tree.png)

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

Download the archive for your Windows architecture [here](https://github.com/TheZoraiz/ascii-image-converter/releases/latest), extract it, and open the extracted folder.
* In Search, search for and then select: System (Control Panel)
* Click the Advanced System settings link.
* Click Environment Variables. In the section User Variables find the Path environment variable and select it. Click "Edit".
* In the Edit Environment Variable window, click "New" and then paste the path of the folder that you copied initially.
* Click "Ok" on all open windows.

Now, restart any open command prompt and execute "ascii-image-converter -h" for more details.

<br>

## Usage

Note: Decrease font size or increase terminal width (like zooming out) for maximum quality ascii art

The basic usage for converting an image into ascii art is as follows. You can also supply paths to multiple images.

```
ascii-image-converter [image-paths]
```
Example:
```
ascii-image-converter myImage.jpeg
```

### Flags

#### --color OR -C

Display ascii art with the colors from original image. Works with the --negative flag as well.

```
ascii-image-converter [image-paths] -C
# Or
ascii-image-converter [image-paths] --color
```

#### --complex OR -c

Print the image with a wider array of ascii characters for more detailed lighting density. Sometimes improves accuracy.
```
ascii-image-converter [image-paths] -c
# Or
ascii-image-converter [image-paths] --complex
```

#### --dimensions OR -d

Note: Don't immediately append another flag with -d

Set the width and height for ascii art in CHARACTER lengths.
```
ascii-image-converter [image-paths] -d <width>,<height>
# Or
ascii-image-converter [image-paths] --dimensions <width>,<height>
```
Example:
```
ascii-image-converter [image-paths] -d 100,30
```

#### --map OR -m

Note: Don't immediately append another flag with -m

Pass a string of your own ascii characters to map against. Passed characters must start from darkest character and end with lightest. There is no limit to number of characters.

Notes: Empty spaces can be passed if string is passed inside quotation marks. You can use both single or double quote for quotation marks. For repeating quotation mark inside string, append it with \ (such as  \\").
  
```
ascii-image-converter [image-paths] -m "<string-of-characters>"
# Or
ascii-image-converter [image-paths] --map "<string-of-characters>"
```
Following example contains 7 depths of lighting.
```
ascii-image-converter [image-paths] -m " .-=+#@"
```

#### --negative OR -n

Display ascii art in negative colors. Works with both uncolored and colored text from --color flag.

```
ascii-image-converter [image-paths] -n
# Or
ascii-image-converter [image-paths] --negative
```

#### --save OR -s

Note: Don't immediately append another flag with -s

Save ascii art in the format `<image-name>-<image-extension>-ascii-art.txt` in the directory path passed to the flag. 

Example for current directory:

```
ascii-image-converter [image-paths] --save .
# Or
ascii-image-converter [image-paths] -s .
```

#### --formats OR -f

Display supported image formats.

```
ascii-image-converter [image-paths] --formats
# Or
ascii-image-converter [image-paths] -f
```

<br>
You can combine flags as well. Following command outputs colored and negative ascii art, with fixed 100 by 30 character dimensions, custom defined ascii characters " .-=+#@" and saves the output in current directory as well.

```
ascii-image-converter [image-paths] -Cnd 100,30 -m " .-=+#@" -s ./
```

<br>

## Contributing

You can fork the project and implement any changes you want for a pull request. However, for major changes, please open an issue first to discuss what you would like to implement.

## Packages used

[github.com/spf13/viper](https://github.com/spf13/viper)

[github.com/spf13/cobra](https://github.com/spf13/cobra)

[github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)

[github.com/nathan-fiscaletti/consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go)

[github.com/nfnt/resize](https://github.com/nfnt/resize)

[github.com/gookit/color](https://github.com/gookit/color)

## License

[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE.txt)