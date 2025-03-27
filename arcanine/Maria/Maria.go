package Maria

import (
	"fmt"
	"github.com/palomino513/arcanine/Tools"
	"os"
)

func IsDBCommand() bool {
	command := os.Args[1]
	shortcut := []rune(command)
	return command == "db" || shortcut[0] == 'D'
}

func getMariaCommand() string {
	if os.Args[1] == "db" {
		return os.Args[2]
	}
	return os.Args[1]
}

func CreateDatabase(containerId string, dbname string) bool {
	Tools.InfoMessage("Se usara el contedor " + containerId)
	query := "CREATE DATABASE IF NOT EXISTS " + dbname + ";"
	Tools.InfoMessage("Creando base de datos " + dbname)
	Tools.RunCommand("docker", "exec", "-it ", containerId, "mysql -u root -p -e", query)
	Tools.SuccessMessage("Hecho")
	return true
}

func ResetDatabase(containerId string, dbname string) bool {
	Tools.InfoMessage("Se usara el contedor " + containerId)
	query := "DROP DATABASE " + dbname + "; CREATE DATABASE " + dbname + ";"
	Tools.InfoMessage("Limpiando la base de datos " + dbname)
	Tools.RunCommand("docker", "exec", "-it", containerId, "mysql", "-u", "root", "-p", "-e", query)
	Tools.SuccessMessage("Hecho")
	return true
}

func FixDatabase(containerId string, dbname string) bool {
	return true
}

func ExecuteMariaCommand(DBName string) bool {
	command := getMariaCommand()
	containerId, err := Tools.GetContainerIDByName("mysql")

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	if command == "create" || command == "Dc" {
		return CreateDatabase(containerId, DBName)
	} else if command == "reset" || command == "Dr" {
		return ResetDatabase(containerId, DBName)
	} else if command == "fix" || command == "Df" {
		return FixDatabase(containerId, DBName)
	} else {
		Tools.InfoMessage("Command not recognized: " + command)
	}

	return false
}
