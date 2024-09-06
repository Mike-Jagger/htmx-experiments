// Just rewritting my own makefile executor since freaking makefile doesn't work on windows
package main

import (
	"bufio"
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

func ParseMakeFile(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	targets := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	var currentTarget string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			currentTarget = strings.TrimSpace(parts[0])
			trailingTargets := strings.TrimSpace(parts[1]) 

			if  trailingTargets != "" && len(trailingTargets) != 0 {
				if !strings.HasPrefix(trailingTargets, "@"){ 
					for _, target := range strings.Fields(trailingTargets) {
						targets[target] = []string{}
						targets[currentTarget] = append(targets[currentTarget], target)

					}
				} else {
					targets[currentTarget] = append(targets[currentTarget], trailingTargets)
				}

			}

			if _, exists := targets[currentTarget]; !exists {
				targets[currentTarget] = []string{}
			}


		} else if currentTarget != "" {
			targets[currentTarget] = append(targets[currentTarget], strings.TrimSpace(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return targets, nil

}	

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Running default commands: build run")

		fmt.Println(ParseMakeFile("makefile"))

		return
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