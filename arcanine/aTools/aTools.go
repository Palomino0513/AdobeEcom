package aTools

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	ColorCyan  = "\033[36m"
	ColorReset = "\033[0m"
)

func RunCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		ErrorMessage("Error al ejecutar el comando:")
		fmt.Println(err)
	}
}

func SuccessMessage(message string) {
	fmt.Println(string(ColorCyan), message, string(ColorReset))
}

func ErrorMessage(message string) {
	fmt.Println(string(ColorCyan), message, string(ColorReset))
}
