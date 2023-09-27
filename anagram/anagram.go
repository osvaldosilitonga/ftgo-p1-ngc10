package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
	"strings"
	"unicode"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	// Profiling
	// # CPU #
	cpu, errCPU := os.Create("cpu.prof")
	if errCPU != nil {
		panic(errCPU)
	}
	defer cpu.Close()

	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	// # Memory #
	memory, errMemory := os.Create("memory.prof")
	if errMemory != nil {
		panic(errMemory)
	}
	defer memory.Close()

	errMemory = pprof.WriteHeapProfile(memory)
	if errMemory != nil {
		panic(errMemory)
	}

	// ----------------------------------

	var kata1, kata2 string

	// menerima input
	fmt.Print("Kata 1 : ")
	fmt.Scan(&kata1)

	fmt.Print("Kata 2 : ")
	fmt.Scan(&kata2)

	if !IsValidHandler(kata1) || !IsValidHandler(kata2) {
		panic("Kata tidak boleh lebih dari 10 karakter atau memiliki simbol")
	}

	anagram := IsAnagram(kata1, kata2)

	if anagram {
		fmt.Println("Anagram : True")
	} else {
		fmt.Println("Anagram : False")
	}
}

func IsAnagram(kata1, kata2 string) bool {

	// to lower
	kata1 = strings.ToLower(kata1)
	kata2 = strings.ToLower(kata2)

	// sort kata
	kataArray1 := []rune(kata1)
	sort.Slice(kataArray1, func(i, j int) bool {
		return kataArray1[i] < kataArray1[j]
	})
	kataArray2 := []rune(kata2)
	sort.Slice(kataArray2, func(i, j int) bool {
		return kataArray2[i] < kataArray2[j]
	})

	// Check
	for i := 0; i < len(kataArray1); i++ {
		if kataArray1[i] != kataArray2[i] {
			return false
		}
	}

	return true

}

func IsValidHandler(kata string) bool {
	if len(kata) > 10 {
		return false
	}

	for _, k := range kata {
		if !unicode.IsLetter(k) {
			return false
		}
	}

	return true
}
