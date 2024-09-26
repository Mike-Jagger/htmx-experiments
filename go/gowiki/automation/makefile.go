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
	prompt = "/c " + prompt
	args := strings.Fields(prompt)
	cmd := exec.Command("cmd.exe", args...)

	fmt.Printf("\t%s:\t%s \n", message, cmd.Args[2:])

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running %s: \n\t%s \n", cmdName, err.Error())
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

func executeTarget(target string) {
	targets, err := ParseMakeFile("makefile")
	if err != nil {
		fmt.Println("Error reading makefile:", err)
		return
	}

	commands, exists := targets[target]
	if !exists {
		fmt.Printf("Target '%s' not found in makefile\n", target)
		return
	}

	fmt.Println("Target:", target)

	for _, cmd := range commands {
		cmd = strings.TrimPrefix(cmd, "@")
		_, exists := targets[cmd]
		if exists {
			fmt.Println("|")
			executeTarget(cmd)
		} else {
			ExecuteCmd("- Executing", cmd, target)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		executeTarget("all")
		return
	}

	target := os.Args[1]

	fmt.Println()

	switch target {
	case "--update":
		fmt.Println("Updating...")
		ExecuteCmd(
			"Updating make command", 
			"go build -o ./make.exe ./automation/makefile.go",
			target,
		)
		// doesn't work :((
		// ExecuteCmd(
		// 	"Updating make command", 
		// 	".\\make --purge",
		// 	target,
		// )

		fmt.Println()
	case "--purge":
		fmt.Println("Purging...")
		ExecuteCmd(
			"Removing dirty file okay",
			"del make.exe~",
			target,
		)

		fmt.Println()
	default:
		fmt.Println("Running makefile...")

		executeTarget(target)
		fmt.Println()
	}
}