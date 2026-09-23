package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		fmt.Printf("Your command was: %s\n", commandName)
	}

	if err := reader.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
