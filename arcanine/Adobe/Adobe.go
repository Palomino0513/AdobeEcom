package Adobe

import (
	"github.com/palomino513/arcanine/Tools"
	"os"
)

func IsAdobeCommand() bool {
	command := os.Args[1]
	shortcut := []rune(command)
	return command == "adobe" || shortcut[0] == 'A'
}

func getAdobeCommand() string {
	if os.Args[1] == "adobe" {
		return os.Args[2]
	}
	return os.Args[1]
}

func RunAdobeCommand(command string) bool {
	return Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", command)
}

func AdobeBash() bool {
	return Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "bash")
}

func AdobeComposerInstall() bool {
	return Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "composer", "install")
}

func AdobeCacheEnable() bool {
	Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "cache:enable", "collections")
	Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "cache:enable", "eav full_page")
	Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "cache:enable", "full_page")
	Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "cache:enable", "layout")
	return true
}

func AdobeCreateAdmin() bool {
	Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "admin:user:create",
		"--admin-user='palomino'",
		"--admin-password='Law200110513'",
		"--admin-email='palomino@wolfsellers.com'",
		"--admin-firstname='palomino'",
		"--admin-lastname='palomino'")
	return true
}

func ExecuteAdobeCommand() bool {
	command := getAdobeCommand()

	if command == "install" || command == "Ai" {
		return AdobeComposerInstall()
	} else if command == "upgrade" || command == "Au" {
		return RunAdobeCommand("setup:upgrade")
	} else if command == "reindex" || command == "Ar" {
		return RunAdobeCommand("indexer:reindex")
	} else if command == "reset" || command == "Az" {
		Tools.RunCommand("docker-compose", "down")
		Tools.RunCommand("rm", "-rf", "generated/*")
		Tools.RunCommand("rm", "-rf", "pub/static/*")
		Tools.RunCommand("docker-compose", "up", "-d")
		RunAdobeCommand("setup:upgrade")
		RunAdobeCommand("cache:clean")
	} else {
		return AdobeBash()
	}

	return false
}
