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

	var depth int = 0
	var dist int = 0
	var aim int = 0

	for _, line := range lines {

		w := strings.Split(line, " ")
		n, err := strconv.Atoi(w[1])
		if err != nil {
			log.Fatalln("Input parsing error", err)
		}
		switch w[0] {
		case "up":
			aim -= n
		case "down":
			aim += n
		case "forward":
			dist += n
			depth += n * aim
		}
	}
	fmt.Printf("%v\n", depth*dist)
}
