package cmd

import (
	"fmt"
	"strings"
)

func debugPrint(format string, args ...any) {
	if optDebug {
		fmt.Printf(format, args...)
	}
}

func isYesAnswer(answer string) bool {
	char := strings.TrimSpace(strings.ToLower(answer))[0]
	return char == 'y'
}
