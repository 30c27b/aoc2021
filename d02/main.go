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

	var height int = 0
	var dist int = 0

	for _, line := range lines {

		w := strings.Split(line, " ")
		n, err := strconv.Atoi(w[1])
		if err != nil {
			log.Fatalln("Input parsing error", err)
		}
		switch w[0] {
		case "up":
			height -= n
		case "down":
			height += n
		case "forward":
			dist += n
		}
	}
	fmt.Printf("%v\n", height*dist)
}
