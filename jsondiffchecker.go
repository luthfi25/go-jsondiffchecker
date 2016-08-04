package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

//for testing purpose
func test() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":[{"Name":"Gomez", "Age":24},{"Name":"Morticia", "Age":33}]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	for key, val := range m {
		fmt.Println("key", key)
		fmt.Println("val", val)
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			arr := val.([]interface{})
			for _, aaa := range arr {
				if reflect.TypeOf(aaa).Kind() == reflect.Map {
					mp := aaa.(map[string]interface{})
					for _, bbb := range mp {
						fmt.Println("val3", bbb)
					}
				}
				fmt.Println("vall", aaa)
			}
		}
	}
}
