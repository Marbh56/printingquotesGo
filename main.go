package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the quote? ")
	quote, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get quote")
		return
	}
	quote = strings.TrimSpace(quote)

	fmt.Print("Who said the quote? ")
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get author")
		return
	}
	author = strings.TrimSpace(author)

	fmt.Printf(`%s says "%s"`, author, quote)
}
