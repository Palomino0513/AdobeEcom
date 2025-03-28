package Adobe

import (
	"github.com/palomino513/arcanine/Tools"
	"os"
)

func IsAdobeCommand() bool {
	command := os.Args[1]
	shortcut := []rune(command)
	return command == "adobe" || shortcut[0] == 'a'
}

func getAdobeCommand() string {
	if os.Args[1] == "adobe" {
		return os.Args[2]
	}
	return os.Args[1]
}

func isAdobeCacheCommand() bool {
	if !IsAdobeCommand() {
		return false
	}

	if os.Args[1] == "adobe" && os.Args[2] == "cache" {
		return true
	}

	shortCommands := []string{"acs", "ace", "acd", "acc", "acf"}
	if Tools.StringInSlice(os.Args[1], shortCommands) {
		return true
	}

	commands := []string{"status", "enable", "disable", "clear", "flush"}
	if os.Args[1] == "adobe" && os.Args[2] == "cache" && len(os.Args) == 3 && Tools.StringInSlice(os.Args[3], commands) {
		return true
	}

	return false
}

func getAdobeCacheCommand() string {
	if len(os.Args) == 3 {
		return os.Args[3]
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
	Tools.InfoMessage("Commando a ejecutar " + command)

	if command == "install" || command == "ai" {
		return AdobeComposerInstall()
	} else if command == "upgrade" || command == "au" {
		return RunAdobeCommand("setup:upgrade")
	} else if command == "reindex" || command == "ar" {
		return RunAdobeCommand("indexer:reindex")
	} else if command == "reset" || command == "az" {
		Tools.RunCommand("docker-compose", "down")
		Tools.RunCommand("rm", "-rf", "generated/*")
		Tools.RunCommand("rm", "-rf", "pub/static/*")
		Tools.RunCommand("docker-compose", "up", "-d")
		RunAdobeCommand("setup:upgrade")
		RunAdobeCommand("cache:clean")
		return true
	} else if command == "static-content" || command == "asc" {
		return Tools.RunCommand("docker-compose", "exec", "-u", "magento2", "web", "php", "bin/magento", "setup:static-content:deploy", "es_MX", "-f")
	} else if command == "create-admin" || command == "aca" {
		return AdobeCreateAdmin()
	} else if isAdobeCacheCommand() {
		subcommand := getAdobeCacheCommand()
		Tools.InfoMessage("Subcomando a ejecutar " + subcommand)

		if subcommand == "status" || subcommand == "acs" {
			return RunAdobeCommand("cache:status")
		} else if subcommand == "enable" || subcommand == "ace" {
			return AdobeCacheEnable()
		} else if subcommand == "disable" || subcommand == "acd" {
			return RunAdobeCommand("cache:disable")
		} else if subcommand == "clean" || subcommand == "acc" {
			return RunAdobeCommand("cache:clean")
		} else if subcommand == "flush" || subcommand == "acf" {
			return RunAdobeCommand("cache:flush")
		} else {
			return RunAdobeCommand("cache:clean")
		}
	}

	return AdobeBash()
}
