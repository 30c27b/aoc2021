package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const inputSize int = 5

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Program requires 1 argument")
	}

	fileName := os.Args[1]

	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("File reading error", err)
	}

	lines := strings.Split(string(f), "\n")

	var valid []bool = make([]bool, inputSize)
	var validCount int = inputSize

	for idx := range valid {
		valid[idx] = true
	}

	for _, line := range lines {
		var frequency [2]int
		var primary bool = false
		for idx := 0; idx < inputSize; idx++ {
			switch line[idx] {
			case '0':
				frequency[0] += 1
			case '1':
				frequency[1] += 1
			}
		}
		if frequency[1] >= frequency[0] {
			primary = true
		}
		for idx := 0; idx < inputSize; idx++ {
			if !valid[idx] {
				break
			}
			switch line[idx] {
			case '0':
				valid[idx] = !primary
				if primary {
					validCount -= 1
				}
			case '1':
				valid[idx] = primary
				if !primary {
					validCount -= 1
				}
			}
		}
		if validCount < 2 {
			break
		}
	}
	fmt.Printf("validCount: %v\n", validCount)
	fmt.Printf("valid: %v\n", valid)
}
