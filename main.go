package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Height = 5
	Width  = 5
)

type Universe [][]bool

func NewUniverse() Universe {
	universe := make([][]bool, Height)

	for row := range universe {
		universe[row] = make([]bool, Width)
	}

	return universe
}

func SymbolMapping(universeRow []bool) string {
	var mappedSymbols string

	for i := 0; i < len(universeRow); i++ {
		if universeRow[i] == true {
			mappedSymbols += "*"
		} else {
			mappedSymbols += " "
		}
	}

	return mappedSymbols
}

func Step(current, future Universe) {
	for y := range current {
		for x := range current[y] {
			future[x][y] = current.Next(x, y)
		}
	}

	current = future
}

func (u Universe) Show() {
	for _, row := range u {
		println(SymbolMapping(row))
	}
}

func (u Universe) Seed() {
	// Seed
	rand.Seed(time.Now().UTC().UnixNano())

	for y := range u {
		for x := range u[y] {
			if rand.Intn(100) < 25 {
				u[y][x] = true
			} else {
				u[y][x] = false
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x += Width
	x %= Width

	y += Height
	y %= Height

	return u[y][x]
}

func (u Universe) Neighbours(x, y int) int {
	var count int

	for yRaise := -1; yRaise < 2; yRaise++ {
		for xRaise := -1; xRaise < 2; xRaise++ {
			if xRaise == 0 && yRaise == 0 {
				continue
			}

			if u.Alive(x+xRaise, y+yRaise) == true {
				count++
			}
		}
	}

	return count
}

func (u Universe) Next(x, y int) bool {
	count := u.Neighbours(x, y)

	if count == 3 && u.Alive(x, y) == false {
		return true
	} else if u.Alive(x, y) == true && count == 2 || count == 3 {
		return true
	}

	return false
}

func main() {
	fmt.Println("\033[H\033[2J")
	universe := NewUniverse()
	universe.Seed()

	tmp := NewUniverse()

	for i := 0; i < 10; i++ {
		universe.Show()
		time.Sleep(time.Second)
		fmt.Println("\033[H\033[2J")
		Step(universe, tmp)
		universe, tmp = tmp, universe
	}
}
