package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func collinear(a Point, b Point, c Point) bool {
	return (b.x-a.x)*(c.y-a.y) == (c.x-a.x)*(b.y-a.y)
}

func within(p int, q int, r int) bool {
	return (p <= q && q <= r) || (r <= q && q <= p)
}

func is_on(a Point, b Point, c Point) bool {
	if a.x != b.x {
		return collinear(a, b, c) && within(a.x, c.x, b.x)
	} else {
		return collinear(a, b, c) && within(a.y, c.y, b.y)
	}
}

func ApplyLine2(line Line, board [][]int) {
	for x := range board {
		for y := range board[x] {
			if is_on(line.a, line.b, Point{x, y}) {
				board[x][y] += 1
			}
		}
	}
}

func ApplyLine(line Line, board [][]int) {
	if line.a.x == line.b.x {

		for i := min(line.a.y, line.b.y); i <= max(line.a.y, line.b.y); i++ {
			board[line.a.x][i] += 1
		}

	} else if line.a.y == line.b.y {

		for i := min(line.a.x, line.b.x); i <= max(line.a.x, line.b.x); i++ {
			board[i][line.a.y] += 1
		}
	}
}

func CountIntersect(board [][]int) int {
	intersect := 0
	for x := range board {
		for y := range board[x] {
			if board[x][y] >= 2 {
				intersect++
			}
		}
	}
	return intersect
}

func BPrint(board [][]int) {
	for y := range board {
		for x := range board[y] {
			if board[x][y] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(board[x][y])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
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

	strLines := strings.Split(string(f), "\n")

	var lines []Line = make([]Line, len(strLines))

	for i, sl := range strLines {
		pts := strings.Split(sl, " -> ")
		coA := strings.Split(pts[0], ",")
		coB := strings.Split(pts[1], ",")
		n, _ := strconv.Atoi(coA[0])
		lines[i].a.x = n
		n, _ = strconv.Atoi(coA[1])
		lines[i].a.y = n
		n, _ = strconv.Atoi(coB[0])
		lines[i].b.x = n
		n, _ = strconv.Atoi(coB[1])
		lines[i].b.y = n
	}

	var board [][]int = make([][]int, 1000)

	for i := range board {
		board[i] = make([]int, 1000)
	}

	for _, line := range lines {
		ApplyLine2(line, board)
	}

	// BPrint(board)

	print(CountIntersect(board))

}
