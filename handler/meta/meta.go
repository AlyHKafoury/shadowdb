package meta

import (
	"errors"
	"os"
	memory "shadowdb/memory/table"
	"strings"
)

//DoMetaCommand executes meta commands starting with "."
func DoMetaCommand(command string, table *memory.Table) error {
	if strings.Compare(command, ".exit") == 0 {
		if err := table.WriteAll(); err != nil {
			return err
		}
		os.Exit(0)
	} else {
		return errors.New("Unknown Meta Command")
	}
	return nil
}
