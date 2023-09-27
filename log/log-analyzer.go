package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Recover Panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	// Open File 'log.txt'
	file, err := os.Open("log.txt")
	if err != nil {
		panic("Tidak dapat membuka file 'log.txt'")
	}
	defer file.Close() // 'log.txt' file akan di close setelah proses selesai

	scanner := bufio.NewScanner(file)

	// Scan 'log.txt'
	var infoCount, errCount, debugCount int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[ERROR]") {
			errCount++
			continue
		}
		if strings.Contains(line, "[DEBUG]") {
			debugCount++
			continue
		}
		if strings.Contains(line, "[INFO]") {
			infoCount++
			continue
		}
	}

	// Display Output
	fmt.Printf("[INFO] Level : %d occurrences\n", infoCount)
	fmt.Printf("[ERROR] Level : %d occurrences\n", errCount)
	fmt.Printf("[DEBUG] Level : %d occurrences\n", debugCount)

}
