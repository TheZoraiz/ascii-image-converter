## Note

These files are just wrappers around [consolesize-go](https://github.com/nathan-fiscaletti/consolesize-go) package. For unix, they resort to terminal size calculation from stdin if stdout is not directed to terminal. For windows, they currently throw an error.