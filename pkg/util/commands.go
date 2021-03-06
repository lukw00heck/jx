package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetCommandOutput evaluates the given command and returns the trimmed output
func GetCommandOutput(dir string, name string, args ...string) (string, error) {
	e := exec.Command(name, args...)
	if dir != "" {
		e.Dir = dir
	}
	data, err := e.CombinedOutput()
	text := string(data)
	text = strings.TrimSpace(text)
	return text, err
}

// RunCommand evaluates the given command and returns the trimmed output
func RunCommand(dir string, name string, args ...string) error {
	e := exec.Command(name, args...)
	if dir != "" {
		e.Dir = dir
	}
	e.Stdout = os.Stdout
	e.Stderr = os.Stdin
	err := e.Run()
	if err != nil {
		fmt.Printf("Error: Command failed  %s %s\n", name, strings.Join(args, " "))
	}
	return err

}
