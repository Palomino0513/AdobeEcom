package main

import (
	"fmt"
	"github.com/palomino513/arcanine/Tools"
	"os"
)

const (
	DBName         = "limpio"
	Hostname       = "http://limpio.local/"
	SecureHostname = "https://limpio.local/"
)

func main() {
	if len(os.Args) < 2 {
		Tools.Info()
		return
	}

	arg := os.Args[1]

	// validamos y ejecutamos en caso de que sea un comando de docker.
	if arg == "up" {
		Tools.RunCommand("docker-compose", "up", "-d")
		return
	} else if arg == "down" {
		Tools.RunCommand("docker-compose", "down")
		return
	} else if arg == "ps" {
		Tools.RunCommand("docker-compose", "ps")
		return
	} else if arg == "restart" {
		Tools.RunCommand("docker-compose", "down")
		Tools.RunCommand("docker-compose", "up", "-d")
		return
	}

	shortcut := []rune(arg)

	// validamos si el comando es de base de datos.
	if arg == "db" || shortcut[0] == 'D' {
		containerId, err := Tools.GetContainerIDByName("mysql")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("ID del contenedor:", containerId)
	}

	Tools.Info()
}
