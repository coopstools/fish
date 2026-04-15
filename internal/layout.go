package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/coopstools/fish/internal/layout"
)

func Open(filename string) *layout.Layout {
	file, err := os.Open(filename)
	if err != nil {
		panic("could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic("failure in scanner: " + scanner.Err().Error())
	}

	return layout.New(lines)
}
