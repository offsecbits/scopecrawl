package aesthetics

import "fmt"



func PrintBanner() {
	fmt.Println(Cyan + `
███████  ██████  ██████  ██████  ███████      ██████ ██████   █████  ██     ██ ██      
██      ██      ██    ██ ██   ██ ██          ██      ██   ██ ██   ██ ██     ██ ██      
███████ ██      ██    ██ ██████  █████       ██      ██████  ███████ ██  █  ██ ██      
     ██ ██      ██    ██ ██      ██          ██      ██   ██ ██   ██ ██ ███ ██ ██      
███████  ██████  ██████  ██      ███████      ██████ ██   ██ ██   ██  ███ ███  ███████ 
` + Reset + Red + `
𝓫𝔂
█▓▒▒░░░ 𝓞𝓯𝓯𝓢𝓮𝓬𝓑𝓲𝓽𝓼 ░░░▒▒▓█
` + Reset)
}
