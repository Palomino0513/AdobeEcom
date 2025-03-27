package Adobe

import (
	"github.com/palomino513/arcanine/Tools"
)

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
