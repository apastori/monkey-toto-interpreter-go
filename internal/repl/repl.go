package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
	"github.com/apastori/monkey-toto-interpreter-go/internal/engine/tokenizer"
)

const PROMPT string = ">>> "

func StartREPL(currentUser string) {
	var inputLayer io.Reader = os.Stdin
	var outputLayer io.Writer = os.Stdout
	var now time.Time = time.Now()
	var nowFormattted string = now.Format("2006-01-02 15:04:05")
	fmt.Println("Welcome to the Monkey Toto programming language REPL")
	fmt.Printf("Feel free to type in commands\n")
	fmt.Printf("Current user: %s\n", currentUser)
	fmt.Println("Current time:", nowFormattted)
	fmt.Println("Type 'exit' to quit")
	var scanner *bufio.Scanner = bufio.NewScanner(inputLayer)
	for {
		fmt.Fprint(outputLayer, PROMPT)
		var scanned bool = scanner.Scan()
		if !scanned {
			return
		}
		var line string = scanner.Text()
		line = strings.TrimSpace(line)
		if line == "exit" {
			fmt.Println("Exiting Monnkey Toto REPL. Goodbye!")
			return
		}
		var tokenizer *tokenizer.Tokenizer = tokenizer.NewTokenizer(line)
		for currentToken := tokenizer.NextToken(); currentToken.Type != token.EOF; currentToken = tokenizer.NextToken() {
			fmt.Printf("%+v\n", currentToken)
		}
	}
}
