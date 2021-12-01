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

	var wins []int = make([]int, len(lines)-2)

	for idx, l := range lines {

		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalln("Input parsing error", err)
		}
		if idx-2 >= 0 && idx-2 < len(wins) {
			wins[idx-2] += n
		}
		if idx-1 >= 0 && idx-1 < len(wins) {
			wins[idx-1] += n
		}
		if idx >= 0 && idx < len(wins) {
			wins[idx] += n
		}
	}

	var count int = 0
	var last int = wins[0]

	for _, n := range wins[1:] {
		if n > last {
			count++
		}
		last = n
	}
	fmt.Printf("count: %v\n", count)
}
