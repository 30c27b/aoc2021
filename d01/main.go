package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	var count int = 0
	var last int = 0

	first, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatalln("Input parsing error", err)
	}
	last = first

	for _, l := range lines[1:] {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalln("Input parsing error", err)
		}
		if n > last {
			count++
		}
		last = n
	}
	fmt.Printf("count: %v\n", count)
}
