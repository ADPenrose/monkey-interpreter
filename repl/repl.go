package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// Scanners are great to read input line by line.
	scanner := bufio.NewScanner(in)

	for {
		// Printing the prompt to the output stream.
		fmt.Fprintf(out, PROMPT)
		// Scanning for the next line. If there is no next line, we return.
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// If a line is successfully scanned, we need the actual content of the line so that
		// we can pass it to the lexer.
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
