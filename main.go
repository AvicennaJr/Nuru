package main

import (
	"fmt"
	"os"

	"github.com/AvicennaJr/Nuru/repl"
)

const (
	LOGO = `

█░░ █░█ █▀▀ █░█ ▄▀█   █▄█ ▄▀█   █▄░█ █░█ █▀█ █░█
█▄▄ █▄█ █▄█ █▀█ █▀█   ░█░ █▀█   █░▀█ █▄█ █▀▄ █▄█                                        

            | Authored by Avicenna |                    
`
)

func main() {

	coloredLogo := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, LOGO)
	fmt.Println(coloredLogo)
	fmt.Println("𝑯𝒂𝒃𝒂𝒓𝒊, 𝒌𝒂𝒓𝒊𝒃𝒖 𝒖𝒕𝒖𝒎𝒊𝒆 𝒍𝒖𝒈𝒉𝒂 𝒚𝒂 𝑵𝒖𝒓𝒖 ✨")

	repl.Start(os.Stdin, os.Stdout)
}
