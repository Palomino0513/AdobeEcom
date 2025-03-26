package main

import (
	"os"
	"github.com/palomino513/arcanine/Tools"
)

func main() {
	if len(os.Args) < 2 {
		Tools.Info()
		return
	}

	arg := os.Args[1]

	switch arg {
	case "up":
		Tools.RunCommand("docker-compose", "up", "-d")

	case "down":
		Tools.RunCommand("docker-compose", "down")

	case "ps":
		Tools.SuccessMessage("Mostrando contenedores ...")
		Tools.RunCommand("docker-compose", "ps")

	case "restart":
		Tools.RunCommand("docker-compose", "down")
		Tools.RunCommand("docker-compose", "up", "-d")

	default:
		Tools.Info()
	}
}
