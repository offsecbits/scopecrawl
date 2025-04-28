package aesthetics

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

func PrintError(msg string) {
	fmt.Println(Red + "[ERROR] " + msg + Reset)
}

func PrintSuccess(msg string) {
	fmt.Println(Green + "[SUCCESS] " + msg + Reset)
}

func PrintWarning(msg string) {
	fmt.Println(Yellow + "" + msg + Reset)
}

func PrintInfo(msg string) {
	fmt.Println(Blue + "" + msg + Reset)
}
