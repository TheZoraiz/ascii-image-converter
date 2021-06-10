// +build unix, !windows

package winsize

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/nathan-fiscaletti/consolesize-go"
)

// By default, this functions calculates terminal dimensions from stdout but in case
// stdout isn't a the terminal, it'll calculate terminal dimensions from stdin. This
// functionality isn't supported for windows yet
func GetTerminalSize() (int, int, error) {

	// Check if stdout is terminal
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return 0, 0, err
	}

	var stdoutIsTerminal bool

	if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		stdoutIsTerminal = true
	} else {
		stdoutIsTerminal = false
	}

	if stdoutIsTerminal {
		x, y := consolesize.GetConsoleSize()
		return x, y, nil

	} else {
		// Get size from stdin if stdout is not terminal

		var sz struct {
			rows    uint16
			cols    uint16
			xpixels uint16
			ypixels uint16
		}
		_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
			uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))

		return int(sz.cols), int(sz.rows), nil
	}
}
