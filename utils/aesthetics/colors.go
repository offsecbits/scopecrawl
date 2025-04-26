package aesthetics

import "fmt"

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan = "\033[36m"
)

func PrintError(msg string) {
	fmt.Println(red + "[ERROR] " + msg + reset)
}

func PrintSuccess(msg string) {
	fmt.Println(green + "[SUCCESS] " + msg + reset)
}

func PrintWarning(msg string) {
	fmt.Println(yellow + "[WARNING] " + msg + reset)
}

func PrintInfo(msg string) {
	fmt.Println(blue + "" + msg + reset)
}
