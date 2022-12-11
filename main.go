package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {

		coloredLogo := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, LOGO)
		fmt.Println(coloredLogo)
		fmt.Println("𝑯𝒂𝒃𝒂𝒓𝒊, 𝒌𝒂𝒓𝒊𝒃𝒖 𝒖𝒕𝒖𝒎𝒊𝒆 𝒍𝒖𝒈𝒉𝒂 𝒚𝒂 𝑵𝒖𝒓𝒖 ✨")

		repl.Start(os.Stdin, os.Stdout)
	} else if len(args) == 1 {
		file := args[0]
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(fmt.Sprintf("\x1b[%dm%s%s\x1b[0m", 31, "Error: Nimeshindwa kusoma file: ", args[0]))
			os.Exit(0)
		}

		repl.Read(string(contents))
	} else {
		fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "Error: Opereshen imeshindikana boss."))
		fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "\nTumia Command: 'nuru' kutmia program AU\nTumia Command: 'nuru' ikifuatwa na program file.\n\n\tMfano:\tnuru fileYangu.nr\n"))
		os.Exit(0)
	}
}
