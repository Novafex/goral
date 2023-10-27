package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func debugPrint(format string, args ...any) {
	if optDebug {
		fmt.Printf(format, args...)
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func checkForConfigs(dir string, exts ...string) (bool, string) {
	for _, ext := range exts {
		if pathExists(filepath.Join(dir, CONFIG_NAME+"."+ext)) {
			return true, ext
		}
	}
	return false, "yaml"
}

func checkForStandardConfigs(dir string) (bool, string) {
	return checkForConfigs(dir, configExtensions...)
}

func isYesAnswer(answer string) bool {
	char := strings.TrimSpace(strings.ToLower(answer))[0]
	return char == 'y'
}
