package log

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	// Open File
	file, err := os.Open("data.txt")
	if err != nil {
		panic("Tidak dapat membuka file 'log.txt'")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// data := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[ERROR]") {
		}
	}

}
