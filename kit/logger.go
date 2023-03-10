package kit

import (
	"fmt"
	"os"
	"strings"
)

var outfile *os.File

func CreateLog(filename string) error {
	if outfile != nil {
		return fmt.Errorf("Logfile already created")
	}
	logfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	outfile = logfile					// i.e. store it in the private var above
	return nil
}

func Log(format_string string, args ...interface{}) {
	
	if outfile == nil {
		return
	}

	s := fmt.Sprintf(format_string, args...)

	var prefix string

	if kframe != nil {
		prefix = fmt.Sprintf("%3d | ", kframe.RealStep())
	} else {
		prefix = "??? | ";
	}

	lines := strings.Split(s, "\n")		// We want to display the current turn even for multi-line inputs...

	fmt.Fprintf(outfile, prefix)
	fmt.Fprintf(outfile, strings.Join(lines, "\n" + prefix))
	fmt.Fprintf(outfile, "\n")
}
