package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadLine takes anything of type io.Reader and returns a trimmed string (initial
// and trailing white space) or an empty string and error if any error
// is encountered.
func ReadLine(in io.Reader) (string, error) {
	out, err := bufio.NewReader(in).ReadString('\n')
	out = strings.TrimSpace(out)
	return out, err
}

func GetInput(prompt string) string {
	// Create a new bufio reader to read input from standard input (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter their name
	fmt.Print(prompt)

	// Read a line of input from the user
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(input)
}
