package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/evaluator"
	"github.com/AvicennaJr/Nuru/lexer"
	"github.com/AvicennaJr/Nuru/object"
	"github.com/AvicennaJr/Nuru/parser"
)

const PROMPT = ">>> "
const ERROR_FACE = `
	███████████████████████████
	███████▀▀▀░░░░░░░▀▀▀███████
	████▀░░░░░░░░░░░░░░░░░▀████
	███│░░░░░░░░░░░░░░░░░░░│███
	██▌│░░░░░░░░░░░░░░░░░░░│▐██
	██░└┐░░░░░░░░░░░░░░░░░┌┘░██
	██░░└┐░░░░░░░░░░░░░░░┌┘░░██
	██░░┌┘▄▄▄▄▄░░░░░▄▄▄▄▄└┐░░██
	██▌░│██████▌░░░▐██████│░▐██
	███░│▐███▀▀░░▄░░▀▀███▌│░███
	██▀─┘░░░░░░░▐█▌░░░░░░░└─▀██
	██▄░░░▄▄▄▓░░▀█▀░░▓▄▄▄░░░▄██
	████▄─┘██▌░░░░░░░▐██└─▄████
	█████░░▐█─┬┬┬┬┬┬┬─█▌░░█████
	████▌░░░▀┬┼┼┼┼┼┼┼┬▀░░░▐████
	█████▄░░░└┴┴┴┴┴┴┴┘░░░▄█████
	███████▄░░░░░░░░░░░▄███████
	██████████▄▄▄▄▄▄▄██████████
	███████████████████████████

  █▄▀ █░█ █▄░█ ▄▀█   █▀ █░█ █ █▀▄ ▄▀█
  █░█ █▄█ █░▀█ █▀█   ▄█ █▀█ █ █▄▀ █▀█

`

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit()" {
			fmt.Println("✨🅺🅰🆁🅸🅱🆄 🆃🅴🅽🅰✨")
			os.Exit(0)
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, colorfy(evaluated.Inspect(), 32))
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, colorfy(ERROR_FACE, 31))
	io.WriteString(out, "Oi! Umeleta shida gani??\n\n")
	io.WriteString(out, "Parser Errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+colorfy(msg, 31)+"\n")
	}
}

func colorfy(str string, colorCode int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, str)
}
