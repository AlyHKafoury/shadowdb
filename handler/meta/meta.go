package meta

import (
	"errors"
	"os"
	"strings"
)

//DoMetaCommand executes meta commands starting with "."
func DoMetaCommand(command string) error {
	if strings.Compare(command, ".exit") == 0 {
		os.Exit(0)
	} else {
		return errors.New("Unknown Meta Command")
	}
	return nil
}
