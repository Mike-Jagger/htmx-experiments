// Just rewritting my own makefile since freaking makefil doesn't work on windows
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecuteCmd(message string, prompt string, cmdName string) {
	fmt.Println(message)
	args := strings.Fields(prompt)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running %s: \n%s \n", cmdName, err.Error())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Running default commands: build run")
	}

	command := os.Args[1]

	switch command {
	case "--update-make":
		ExecuteCmd(
			"Updating make command", 
			"go build -o ./make.exe ./automation/makefile.go",
			command,
		)
	case "run":
	}
}