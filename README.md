# ascii-image-converter

ascii-image-converter is a command-line tool that converts images into ascii art and prints them out onto the console. It is cross-platform so both Windows and Linux distributions are supported

Image formats currently supported:
* JPEG/JPG
* PNG
* WEBP
* BMP
* TIFF/TIF

<br>

### Example ([Source](https://medium.com/@sean.glancy/practical-applications-of-binary-trees-3097cf663062)):
![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/tree.png)

### ASCII Art:
![Example](https://raw.githubusercontent.com/TheZoraiz/ascii-image-converter/master/example_images/ascii_tree.png)

<br>


## Installation

### Go
For custom Go installs
```
go install github.com/TheZoraiz/ascii-image-converter@latest
```

For physically installing the binaries, download the executables from [here](https://github.com/TheZoraiz/ascii-image-converter/releases/tag/v1.2.0), and follow the steps with respect to your OS.

### Linux
Extract Executables.zip and open the "Linux" directory.

Now, open a terminal in the same directory and execute this command:

```
sudo cp ascii-image-converter /usr/local/bin/
```
Now you can use ascii-image-converter in the terminal. Execute "ascii-image-converter -h" for more details.

### Windows

You will need to set an Environment Variable to the folder the ascii-image-converter.exe executable is placed in to be able to use it in the command prompt. Follow the instructions in case of confusion:

Extract Executables.zip and open the "Windows" folder. Copy the path to folder from the top of the file explorer.
* In Search, search for and then select: System (Control Panel)
* Click the Advanced System settings link.
* Click Environment Variables. In the section User Variables find the Path environment variable and select it. Click "Edit".
* In the Edit Environment Variable window, click "New" and then paste the path of the folder that you copied initially.
* Afterwards, you can use it anywhere by typing "ascii-image-converter" in command prompt. Note: Make sure you restart the command prompt.

<br>

## Usage

Note: Decrease font size or zoom out of terminal for maximum quality ascii art

To convert an image into ascii format, the usage is as follows:
```
ascii-image-converter [path to image]
```
Example
```
ascii-image-converter myImage.jpeg
```
<br>

### Flags

#### --complex OR -c
Print the image with a wider array of ascii characters. Sometimes improves accuracy.
```
ascii-image-converter [path to image] -c
# Or
ascii-image-converter [path to image] --complex
```


#### --dimensions OR -d
Set the width and height for ascii art in CHARACTER lengths. (Don't immediately append another flag with -d)
```
ascii-image-converter [path to image] -d <width>,<height>
# Or
ascii-image-converter [path to image] --dimensions <width>,<height>
```
Example:
```
ascii-image-converter [path to image] -d 100,30
```

#### --color OR -C
Display ascii art with the colors from original image. Works with the -n flag as well.
```
ascii-image-converter [path to image] -C
# Or
ascii-image-converter [path to image] --color
```

#### --negative OR -n
Display ascii art in negative colors. Works with both uncolored and colored text from -C flag.
```
ascii-image-converter [path to image] -n
# Or
ascii-image-converter [path to image] --negative
```

#### --save OR -s
Save the printed ascii art in a file ascii-image.txt in the directory passed alongside. (Don't immediately append another flag with -s)

Example for current directory:
```
ascii-image-converter [path to image] --save ./
# Or
ascii-image-converter [path to image] -s ./
```

#### --formats OR -f
Display supported image formats.
```
ascii-image-converter [path to image] --formats
# Or
ascii-image-converter [path to image] -f
```

<br>

You can combine flags as well. Following command outputs colored and negative ascii art, with complex characters, fixed 100 by 30 character dimensions and saves the output in current directory as well.
```
ascii-image-converter [path to image] -Ccnd 100,30 -s ./
```

<br>

## Contributing

You can fork the project and implement any changes you want for a pull request. However, for major changes, please open an issue first to discuss what you would like to implement.

## Packges used

[github.com/spf13/viper](https://github.com/spf13/viper)

[github.com/spf13/cobra](https://github.com/spf13/cobra)

[github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)

[github.com/nathan-fiscaletti/consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go)

[github.com/nfnt/resize](https://github.com/nfnt/resize)

[github.com/gookit/color](https://github.com/gookit/color)


## License
[Apache-2.0](https://github.com/TheZoraiz/ascii-image-converter/blob/master/LICENSE)
