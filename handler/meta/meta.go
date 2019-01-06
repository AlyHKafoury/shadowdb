package meta

import (
	"errors"
	"os"
	"strings"
)

//DoMetaCommand executes meta commands starting with "."
func DoMetaCommand(command []rune) error {
	if strings.Compare(string(command), ".exit") == 0 {
		os.Exit(0)
	} else {
		return errors.New("Unknown Meta Command")
	}
	return nil
}
