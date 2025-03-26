package Maria

import (
	"github.com/palomino513/arcanine/Tools"
)

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
	Tools.RunCommand("docker", "exec", "-it ", containerId, "mysql -u root -p -e", query)
	Tools.SuccessMessage("Hecho")
	return true
}

func FixDatabase(dbname string) bool {
	return true
}
