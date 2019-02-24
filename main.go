package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shadowdb/handler/meta"
	"shadowdb/handler/statement"
	"shadowdb/memory/table"
)

func printPrompt() {
	fmt.Print("Shadow-DB *>> ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	currentTable := table.New()
	for {
		printPrompt()
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		command := scanner.Text()
		if string(command[0]) == "." {
			if err := meta.DoMetaCommand(command); err != nil {
				log.Println(err)
			}
			continue
		}
		statement := new(statement.Statement)
		if err := statement.Prepare(command); err != nil {
			log.Println(err)
			continue
		}
		if err := statement.Execute(&currentTable); err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Executed command")
	}
}
