package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingo struct {
	nbr    int
	marked bool
}

func Mark(table [][]bingo, nbr int) {
	for i := range table {
		for j := range table {
			if table[i][j].nbr == nbr {
				table[i][j].marked = true
			}
		}
	}
}

func HasWon(table [][]bingo) bool {

	for i := range table {
		okLines := 0
		for j := range table {
			if table[i][j].marked {
				okLines++
			}
		}
		if okLines > 4 {
			return true
		}
	}

	for i := range table {
		okRows := 0
		for j := range table {
			if table[j][i].marked {
				okRows++
			}
		}
		if okRows > 4 {
			return true
		}
	}

	return false
}

func MakeDraws(tables [][][]bingo, drawnNbr []int) ([][]bingo, int) {
	for _, drw := range drawnNbr {
		for idx := range tables {
			Mark(tables[idx], drw)
			if HasWon(tables[idx]) {
				return tables[idx], drw
			}
		}
	}
	return nil, 0
}

func GetDrawId(table [][]bingo, drawnNbr []int) int {
	for idx, drw := range drawnNbr {
		Mark(table, drw)
		if HasWon(table) {
			return idx
		}
	}
	return 0
}

func CountPts(table [][]bingo) int {
	pts := 0
	for i := range table {
		for j := range table {
			if !table[i][j].marked {
				pts += table[i][j].nbr
			}
		}
	}
	return pts
}

func MaxVal(a []int) int {
	m := 0
	mid := 0
	for i, e := range a {
		if i == 0 || e > m {
			m = e
			mid = i
		}
	}
	return mid
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

	lines := strings.Split(string(f), "\n")

	drawn := strings.Split(lines[0], ",")

	drawnNbr := make([]int, len(drawn))

	for idx := range drawn {

		n, err := strconv.Atoi(drawn[idx])
		if err != nil {
			log.Fatalln(err.Error())
		}
		drawnNbr[idx] = n
	}

	tables := make([][][]bingo, (len(lines)-1)/6)

	iter := 0

	for i := 1; i < len(lines); i += 6 {

		table := make([][]bingo, 5)

		for y := 0; y < 5; y++ {

			table[y] = make([]bingo, 5)

			tableLine := strings.Fields(lines[i+1+y])

			for idx := range table[y] {

				n, err := strconv.Atoi(tableLine[idx])
				if err != nil {
					log.Fatalln(err.Error())
				}
				table[y][idx].nbr = n
			}
		}
		tables[iter] = table
		iter++
	}

	tablesRes := make([]int, len(tables))

	for idx, t := range tables {
		tablesRes[idx] = GetDrawId(t, drawnNbr)
	}

	maxval := MaxVal(tablesRes)

	pts := CountPts(tables[maxval])

	lastPick := drawnNbr[tablesRes[maxval]]
	fmt.Printf("lastPick: %v\n", lastPick)
	fmt.Printf("pts: %v\n", pts*lastPick)
}
