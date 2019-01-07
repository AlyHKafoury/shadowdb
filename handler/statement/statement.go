package statement

import (
	"errors"
	"fmt"
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
}

//Prepare make statement ready for execution
func (statement *Statement) Prepare(command []rune) {
	switch string(command[0:6]) {
	case "insert":
		statement.Type = StatementInsert
	case "select":
		statement.Type = StatementSelect
	default:
		statement.Type = StatementUnknown
	}
}

//Execute runs the command
func (statement *Statement) Execute() error {
	switch statement.Type {
	case StatementUnknown:
		return errors.New("Unknown command")
	case StatementInsert:
		fmt.Println("Do insert")
	case StatementSelect:
		fmt.Println("Do select")
	}
	return nil
}
