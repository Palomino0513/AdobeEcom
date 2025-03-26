package Tools

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	ColorInfo    = "\033[36m"
	ColorSuccess = "\033[36m"
	ColorError   = "\033[36m"
	ColorWarning = "\033[36m"
	ColorReset   = "\033[0m"
)

func Info() {
	fmt.Println(string(ColorInfo), "Este commando se hizo con el objetivo de ejecutar comandos communes de Magento de"+
		" la forma mas rapida y practica posible.", string(ColorReset))
}

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

func InfoMessage(message string) {
	fmt.Println(string(ColorInfo), message, string(ColorReset))
}

func SuccessMessage(message string) {
	fmt.Println(string(ColorSuccess), message, string(ColorReset))
}

func ErrorMessage(message string) {
	fmt.Println(string(ColorError), message, string(ColorReset))
}

func WarningMessage(message string) {
	fmt.Println(string(ColorWarning), message, string(ColorReset))
}
