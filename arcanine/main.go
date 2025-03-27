package main

import (
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

	// validamos y ejecutamos en caso de que sea un comando de docker.
	if Tools.IsGeneralCommand() {
		Tools.ExecuteGeneralCommand()
	}

	// validamos si el comando es de base de datos.
	if Maria.IsDBCommand() {
		Maria.ExecuteMariaCommand(DBName)
	}

	// validamos y ejecutamos en caso de que sea un comando de Adobe.
	if Adobe.IsAdobeCommand() {
		Adobe.ExecuteAdobeCommand()
	}

	Tools.Info()
}
