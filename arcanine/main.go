package main

import (
	"fmt"
	"os"
	"github.com/palomino513/arcanine/aTools"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run script.go [up|down|ps|restart]")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "up":
		aTools.RunCommand("docker-compose", "up", "-d")

	case "down":
		aTools.RunCommand("docker-compose", "down")

	case "ps":
		aTools.SuccessMessage("Mostrando contenedores ...")
		aTools.RunCommand("docker-compose", "ps")

	case "restart":
		aTools.RunCommand("docker-compose", "down")
		aTools.RunCommand("docker-compose", "up", "-d")

	default:
		fmt.Println("Comando no reconocido:", arg)
	}
}
