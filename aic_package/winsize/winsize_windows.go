// +build windows, !unix

package winsize

import (
	"fmt"

	"github.com/nathan-fiscaletti/consolesize-go"
)

// By default, this functions calculates terminal dimensions from stdout but in case
// stdout isn't a the terminal, it'll throw an error instead of panicking.
func GetTerminalSize() (int, int, error) {
	x, y := consolesize.GetConsoleSize()

	if x < 1 && y < 1 {
		return x, y, fmt.Errorf("altering stdout isn't currently supported on windows")
	} else {
		return x, y, nil
	}
}
