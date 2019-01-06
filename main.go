package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shadowdb/handler/meta"
)

func printPrompt() {
	fmt.Print("Shadow-DB >>> ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		command := []rune(scanner.Text())
		if command[0] == rune("."[0]) {
			if err := meta.DoMetaCommand(command); err != nil {
				log.Println(err)
			}
		}
	}
}
