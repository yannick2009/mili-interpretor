package repl

import (
	"bufio"
	"fmt"
	"io"
	"mili/lexer"
	"mili/token"
)

const PROMPT = ">> "

// start the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
