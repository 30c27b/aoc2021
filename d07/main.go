package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func FuelCost(start int, target int) int64 {
	dist := int(math.Abs(float64(target - start)))
	fuel := int64(0)
	for i := 1; i <= dist; i++ {
		fuel += int64(i)
	}
	return fuel
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Program requires 1 argument")
	}

	fileName := os.Args[1]

	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("File reading error", err)
	}

	positionsStr := strings.Split(string(f), ",")

	positions := make([]int, len(positionsStr))

	for idx, ps := range positionsStr {
		n, _ := strconv.Atoi(ps)
		positions[idx] = n
	}

	maxPos := 0

	for _, p := range positions {
		if p > maxPos {
			maxPos = p
		}
	}

	bestPos := 0
	bestFuelCost := int64(maxPos * maxPos * 100000)

	for target := 0; target <= maxPos; target++ {

		fuel := int64(0)

		for _, pos := range positions {
			fuel += FuelCost(pos, target)
		}

		if fuel < bestFuelCost {
			bestPos = target
			bestFuelCost = fuel
		}

	}

	fmt.Printf("bestPos: %v\n", bestPos)
	fmt.Printf("bestFuelCost: %v\n", bestFuelCost)
}
