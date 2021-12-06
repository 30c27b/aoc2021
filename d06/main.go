package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxDays = 256

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Program requires 1 argument")
	}

	fileName := os.Args[1]

	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("File reading error", err)
	}

	elstr := strings.Split(string(f), ",")

	var ages [9]int64

	for _, el := range elstr {
		n, _ := strconv.Atoi(el)
		ages[n] += 1
	}

	for day := 0; day < maxDays; day++ {
		var newAges [9]int64
		for idx := 0; idx < 9; idx++ {
			nf := ages[idx]
			if idx == 0 {
				newAges[6] += nf
				newAges[8] += nf
			} else {
				newAges[idx-1] += nf
			}
		}
		ages = newAges
	}

	var sum int64 = 0
	for _, v := range ages {
		sum += v
	}

	fmt.Printf("nbr fish: %v\n", sum)

}
