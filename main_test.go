package main

import (
	"strings"
	"testing"
)

func CreateUniverse(t *testing.T) Universe {
	t.Helper()

	return Universe{
		{true, true, true, true, true},
		{false, false, false, false, true},
		{false, false, false, false, false},
		{true, false, false, false, false},
		{true, true, true, true, false},
	}
}

func TestRightSeizedUniverse(t *testing.T) {
	universe := NewUniverse()

	receivedHeight := len(universe)
	receivedWidth := len(universe[0])

	if receivedHeight != Height || receivedWidth != Width {
		t.Errorf("Expected %d as height but got %d. Expected %d as width but got %d", Height, receivedHeight, Width, receivedWidth)
	}
}

func TestCorrectSymbolMapping(t *testing.T) {
	universeRow := []bool{true, true, true, false, true}

	mappedUniverse := SymbolMapping(universeRow)

	if strings.Compare(mappedUniverse, "*** *") != 0 {
		t.Errorf("Expected mapping of %s but got %s", "**  *", mappedUniverse)
	}
}

func TestLivingCellWithinArray(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Alive(0, 0)

	if alive != true {
		t.Errorf("Expected that the cell at [0,0] is alive but it was dead")
	}
}

func TestDeadCellWithinArray(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Alive(1, 0)

	if alive == true {
		t.Errorf("Expected that the cell at [0,0] is dead but it was alive")
	}
}

func TestAliveCellPositiveBeyondOfArray(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Alive(Width, Height)

	if alive != true {
		t.Errorf("Expected that the cell at [0,0] is alive but it was dead")
	}
}

func TestDeadCellNegativeBeyondOfArray(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Alive(-1, -1)

	if alive == true {
		t.Errorf("Expected that the cell at [0,0] is dead but it was alive")
	}
}

func TestLivingNeighbours(t *testing.T) {
	universe := CreateUniverse(t)

	count := universe.Neighbours(4, 3)

	if count != 4 {
		t.Errorf("Expected 4 living neighbours but got %d", count)
	}
}

func TestLessThanTwoLivingNeighbours(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Next(2, 2)

	if alive == true {
		t.Errorf("Expected that cell will be dead in the next generation but it will be alive")
	}
}

func TestWithTwoLivingNeighbours(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Next(3, 0)

	if alive != true {
		t.Errorf("Expected that cell will be alive in the next generation but it will be dead")
	}
}

func TestWithThreeLivingNeighbours(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Next(3, 4)

	if alive != true {
		t.Errorf("Expected that cell will be alive in the next generation but it will be dead")
	}
}

func TestWithThreeLivingNeighboursAndDead(t *testing.T) {
	universe := CreateUniverse(t)

	alive := universe.Next(1, 1)

	if alive != true {
		t.Errorf("Expected that cell will be alive in the next generation but it will be dead")
	}
}
