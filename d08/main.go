// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"sort"
// 	"strings"
// )

// type Display struct {
// 	segments string
// 	value    int
// }

// type Input struct {
// 	displays [10]Display
// 	output   [4]string
// 	dmap     DMap
// }
// type DMap struct {
// 	top    rune
// 	topL   rune
// 	topR   rune
// 	center rune
// 	botL   rune
// 	botR   rune
// 	bot    rune
// }

// func StringToRuneSlice(s string) []rune {
// 	var r []rune
// 	for _, runeValue := range s {
// 		r = append(r, runeValue)
// 	}
// 	return r
// }

// func SortStringByCharacter(s string) string {
// 	r := StringToRuneSlice(s)
// 	sort.Slice(r, func(i, j int) bool {
// 		return r[i] < r[j]
// 	})
// 	return string(r)
// }

// func (m *DMap) Print() {
// 	fmt.Printf(" %c%c%c%c \n", m.top, m.top, m.top, m.top)
// 	fmt.Printf("%c    %c\n", m.topL, m.topR)
// 	fmt.Printf("%c    %c\n", m.topL, m.topR)
// 	fmt.Printf("%c    %c\n", m.topL, m.topR)
// 	fmt.Printf("%c    %c\n", m.topL, m.topR)
// 	fmt.Printf(" %c%c%c%c \n", m.center, m.center, m.center, m.center)
// 	fmt.Printf("%c    %c\n", m.botL, m.botR)
// 	fmt.Printf("%c    %c\n", m.botL, m.botR)
// 	fmt.Printf("%c    %c\n", m.botL, m.botR)
// 	fmt.Printf("%c    %c\n", m.botL, m.botR)
// 	fmt.Printf(" %c%c%c%c \n", m.bot, m.bot, m.bot, m.bot)
// }

// func GuessValues(input *Input) {
// 	sort.Slice(input.displays[:], func(i, j int) bool {
// 		return len(input.displays[i].segments) < len(input.displays[j].segments)
// 	})
// 	for i := range input.displays {
// 		// fmt.Printf("d: %v\n", d)
// 		switch len(input.displays[i].segments) {
// 		case 2:
// 			input.displays[i].value = 1
// 			input.dmap.topR = []rune(input.displays[i].segments)[0]
// 			input.dmap.botR = []rune(input.displays[i].segments)[1]
// 		case 3:
// 			input.displays[i].value = 7
// 			r := []rune(input.displays[i].segments)
// 			for _, ru := range r {
// 				if ru == input.dmap.topR || ru == input.dmap.botR {
// 					continue
// 				}
// 				input.dmap.top = ru
// 				break
// 			}
// 		case 4:
// 			input.displays[i].value = 4
// 			r := []rune(input.displays[i].segments)
// 			for _, ru := range r {
// 				if ru == input.dmap.topR || ru == input.dmap.botR || ru == input.dmap.top {
// 					continue
// 				}

// 				input.dmap.topL = ru
// 				break
// 			}
// 			for _, ru := range r {
// 				if ru == input.dmap.topR || ru == input.dmap.botR || ru == input.dmap.top || ru == input.dmap.topL {
// 					continue
// 				}
// 				input.dmap.center = ru
// 				break
// 			}
// 		case 7:
// 			input.displays[i].value = 8
// 			r := []rune(input.displays[i].segments)
// 			for _, ru := range r {
// 				if ru == input.dmap.top || ru == input.dmap.topR || ru == input.dmap.topL || ru == input.dmap.center || ru == input.dmap.botR {
// 					continue
// 				}
// 				input.dmap.botL = ru
// 				break
// 			}
// 			for _, ru := range r {
// 				if ru == input.dmap.top || ru == input.dmap.topR || ru == input.dmap.topL || ru == input.dmap.center || ru == input.dmap.botR || ru == input.dmap.botL {
// 					continue
// 				}

// 				input.dmap.bot = ru
// 				break
// 			}
// 		}
// 	}

// 	m := input.dmap
// 	two := SortStringByCharacter(string(m.top) + string(m.topR) + string(m.center) + string(m.botL) + string(m.bot))
// 	tree := SortStringByCharacter(string(m.top) + string(m.topR) + string(m.center) + string(m.botR) + string(m.bot))
// 	five := SortStringByCharacter(string(m.top) + string(m.topL) + string(m.center) + string(m.botR) + string(m.bot))
// 	zero := SortStringByCharacter(string(m.top) + string(m.topL) + string(m.topR) + string(m.botL) + string(m.botR) + string(m.bot))
// 	six := SortStringByCharacter(string(m.top) + string(m.topL) + string(m.center) + string(m.botL) + string(m.botR) + string(m.bot))
// 	nine := SortStringByCharacter(string(m.top) + string(m.topL) + string(m.topR) + string(m.center) + string(m.botL) + string(m.bot))

// 	for i := range input.displays {
// 		input.displays[i].segments = SortStringByCharacter(input.displays[i].segments)
// 		switch input.displays[i].segments {
// 		case two:
// 			input.displays[i].value = 2
// 		case tree:
// 			input.displays[i].value = 3
// 		case five:
// 			input.displays[i].value = 5
// 		case zero:
// 			input.displays[i].value = 0
// 		case six:
// 			input.displays[i].value = 6
// 		case nine:
// 			input.displays[i].value = 9
// 		}
// 	}
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		log.Fatalln("Program requires 1 argument")
// 	}

// 	fileName := os.Args[1]

// 	f, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Fatalln("File reading error", err)
// 	}

// 	lines := strings.Split(string(f), "\n")

// 	entries := make([]Input, len(lines))

// 	for i, line := range lines {

// 		parts := strings.Split(line, " | ")
// 		displays := strings.Split(parts[0], " ")
// 		outputs := strings.Split(parts[1], " ")
// 		copy(entries[i].output[:], outputs)
// 		for j := range entries[i].displays {
// 			entries[i].displays[j].segments = displays[j]
// 		}
// 	}

// 	for _, e := range entries {
// 		GuessValues(&e)

// 		fmt.Printf("e.displays: %v\n", e.displays)
// 	}

// }
