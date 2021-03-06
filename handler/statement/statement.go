package statement

import (
	"bytes"
	"errors"
	"fmt"
	"shadowdb/memory/helpers"
	"shadowdb/memory/table"
)

//Type denotes type of statement
type Type int8

const (
	// StatementUnknown unknown statement tyep
	StatementUnknown = iota
	// StatementInsert statement of type insert
	StatementInsert
	// StatementSelect statement of type select
	StatementSelect
)

//Statement object type
type Statement struct {
	Type Type
	Row  table.Row
}

//Prepare make statement ready for execution
func (statement *Statement) Prepare(input string) error {
	if len(input) < 6 {
		return errors.New("Wrong Command")
	}
	command := string(input[0:6])
	switch command {
	case "insert":
		statement.Type = StatementInsert
		var tempUserName, tempEmail []byte
		r := bytes.NewReader([]byte(input))
		numberOfItems, err := fmt.Fscanf(r, "insert %d %s %s", &statement.Row.ID, &tempUserName, &tempEmail)
		if len(tempUserName) > len(statement.Row.Username) || len(tempEmail) > len(statement.Row.Email) {
			fmt.Println("Input too long")
			return errors.New("Input too long")
		}
		if numberOfItems != 3 || err != nil {
			return errors.New("Syntax Error insert (number) (string) (string)")
		}
		copy(statement.Row.Username[:], tempUserName)
		copy(statement.Row.Email[:], tempEmail)
	case "select":
		statement.Type = StatementSelect
	default:
		statement.Type = StatementUnknown
	}
	return nil
}

//Execute runs the command
func (statement *Statement) Execute(currentTable *table.Table) error {
	switch statement.Type {
	case StatementUnknown:
		return errors.New("Unknown command")
	case StatementInsert:
		rowbytes, err := helpers.RowToBytes(statement.Row)
		if err != nil {
			return err
		}
		if err = currentTable.Insert(rowbytes); err != nil {
			return err
		}
		fmt.Println("Added Row to table")
	case StatementSelect:
		for i := int8(0); i <= currentTable.CurrentPage; i++ {
			var lastRow uint32
			if i == currentTable.CurrentPage {
				lastRow = currentTable.RowInPage
			} else {
				lastRow = uint32(table.RowsPerPage)
			}
			for j := 0; uint32(j) < lastRow; j++ {
				rowBytes := currentTable.ReadRow(i, uint32(j))
				rowString, err := helpers.BytesToRow(rowBytes)
				if err != nil {
					return err
				}
				fmt.Println(rowString.ID, string(bytes.Trim(rowString.Username[:], "\x00")), string(bytes.Trim(rowString.Email[:], "\x00")))
			}
		}
	}
	return nil
}
