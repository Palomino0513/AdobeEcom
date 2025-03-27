package main

import (
	"fmt"
	"github.com/palomino513/arcanine/Adobe"
	"github.com/palomino513/arcanine/Maria"
	"github.com/palomino513/arcanine/Tools"
	"os"
)

const (
	DBName         = "vitro"
	Hostname       = "http://vitro.local/"
	SecureHostname = "https://vitro.local/"
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
	} else if arg == "reset" {
		Tools.RunCommand("docker-compose", "down")
		Tools.RunCommand("rm", "-rf", "generated/*")
		Tools.RunCommand("rm", "-rf", "pub/static/*")
		Tools.RunCommand("docker-compose", "up", "-d")
		Adobe.RunAdobeCommand("setup:upgrade")
		Adobe.RunAdobeCommand("cache:clean")
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

		Tools.InfoMessage("ID del contenedor: " + containerId)

		command := func() string {
			if arg == "db" {
				return os.Args[2]
			}
			return arg
		}()

		if command == "create" || command == "Dc" {
			Maria.CreateDatabase(containerId, DBName)
		} else if command == "reset" || command == "Dr" {
			Maria.ResetDatabase(containerId, DBName)
		} else if command == "fix" || command == "Df" {
			Maria.FixDatabase(containerId, DBName)
		} else {
			Tools.InfoMessage("Command not recognized: " + arg)
		}

		return
	}

	Tools.Info()
}
