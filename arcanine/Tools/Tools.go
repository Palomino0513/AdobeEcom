package Tools

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
	"strings"
)

const (
	ColorReset  = "\033[0m"  // Resetear color
	ColorRed    = "\033[31m" // Error
	ColorGreen  = "\033[32m" // Éxito
	ColorYellow = "\033[33m" // Advertencia
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func Info() {
	fmt.Println(string(ColorWhite), "Este commando se hizo con el objetivo de ejecutar comandos communes de Magento de"+
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
	fmt.Println(string(ColorCyan), message, string(ColorReset))
}

func SuccessMessage(message string) {
	fmt.Println(string(ColorGreen), message, string(ColorReset))
}

func ErrorMessage(message string) {
	fmt.Println(string(ColorRed), message, string(ColorReset))
}

func WarningMessage(message string) {
	fmt.Println(string(ColorYellow), message, string(ColorReset))
}

func GetContainerIDByName(name string) (string, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", fmt.Errorf("error creating Docker client: %w", err)
	}

	// Lista todos los contenedores en ejecución
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return "", fmt.Errorf("error listing containers: %w", err)
	}

	for _, container := range containers {
		for _, containerName := range container.Names {
			cleanName := strings.TrimPrefix(containerName, "/")
			if cleanName == name {
				return container.ID, nil
			}
		}
	}

	return "", fmt.Errorf("container with name '%s' not found", name)
}
