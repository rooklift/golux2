package kit

import (
	"fmt"
	"os"
)

var logging_actions = false			// Used by send_actions()

var outfile *os.File

func CreateLog(filename string) error {
	if outfile != nil {
		return fmt.Errorf("Logfile already created")
	}
	logfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	outfile = logfile				// i.e. store it in the private var above
	return nil
}

func Log(format_string string, args ...interface{}) {
	if outfile == nil {
		return
	}
	fmt.Fprintf(outfile, format_string, args...)
	fmt.Fprintf(outfile, "\n")
}

func LogActions(b bool) {
	logging_actions = b
}
