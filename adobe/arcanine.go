package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	colorCyan  = "\033[36m"
	colorReset = "\033[0m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run script.go [up|down|ps|restart]")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "up":
		runCommand("docker-compose", "up", "-d")

	case "down":
		runCommand("docker-compose", "down")

	case "ps":
		fmt.Println(string(colorCyan), "Mostrando contenedores ...", string(colorReset))
		runCommand("docker-compose", "ps")

	case "restart":
		runCommand("docker-compose", "down")
		runCommand("docker-compose", "up", "-d")

	default:
		fmt.Println("Comando no reconocido:", arg)
	}
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
	}
}
