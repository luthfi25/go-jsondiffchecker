package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("path : ")
	badPath, err := reader.ReadString('\n')
	goodPath := strings.Split(badPath, "\n")

	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := readLines(goodPath[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, playground\n", strings.Join(text, "\n"))
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
