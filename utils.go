package bunker

import (
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func IsEmptyString(str string) bool {
	return strings.Compare(strings.TrimSpace(str), "") == 0
}

func PrintErr(msg string) {
	println(color.Ize(color.Red, fmt.Sprintf("[DEBUG] %s", msg)))
}

func FatalPrintErr(msg string) {
	println(color.Ize(color.Red, fmt.Sprintf("[DEBUG] %s", msg)))
	os.Exit(1)
}

func LogInfo(msg string) {
	println(color.Ize(color.Gray, fmt.Sprintf("[DEBUG] %s", msg)))
}

func LogWarn(msg string) {
	println(color.Ize(color.Yellow, fmt.Sprintf("[DEBUG] %s", msg)))
}
