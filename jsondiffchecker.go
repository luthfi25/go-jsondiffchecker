package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

	js, err := readLines(goodPath[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	unmarshal(js)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	lines := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += " " + scanner.Text()
	}
	return lines, scanner.Err()
}

//unmarshal the json
func unmarshal(path string) {
	fmt.Println()

	b := []byte(path)
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})
	for key, val := range m {
		decode(key, val)
	}
}

//decode and print unmarshaled json
func decode(key string, val interface{}) {
	kind := reflect.ValueOf(val).Kind()

	if kind == reflect.Map {
		m := val.(map[string]interface{})

		for k, v := range m {
			decode(key+" -> "+k, v)
		}
	} else if kind == reflect.Slice {
		s := val.([]interface{})

		for k, v := range s {
			decode(key+" -> "+strconv.Itoa(k), v)
		}

	} else {
		fmt.Println(key, ":", val)
	}
}
