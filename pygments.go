// Pygments wrapper for golang. Pygments is a syntax highlighter

package pygments

import (
	"bytes"
	"fmt"
	"os/exec"
)

var (
	bin = "pygmentize"
)

func Binary(path string) {
	bin = path
}

func Which() string {
	return bin
}

func Highlight(code []byte, lexer string, format string, enc string) (string, error) {

	if _, err := exec.LookPath(bin); err != nil {
		fmt.Println("You do not have " + bin + " installed!")
		return "", err
	}

	// Guess the lexer based on content if a specific one is not provided
	lexerArg := "-g"
	if lexer != "" {
		lexerArg = "-l" + lexer
	}

	cmd := exec.Command(bin, lexerArg, "-f"+format, "-O encoding="+enc)
	cmd.Stdin = bytes.NewReader(code)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(stderr.String())
	}

	return out.String(), nil
}
