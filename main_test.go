package main

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	m := [][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
	}
	f := &Field{m, 3, 3}

	counter := [][]int{
		[]int{3, 5, 3},
		[]int{5, 8, 5},
		[]int{3, 5, 3},
	}

	for i := 0; i < f.width; i++ {
		for j := 0; j < f.height; j++ {
			c := f.countNeighbours(i, j)
			if c != counter[i][j] {
				fmt.Printf("problem at [%d, %d]", i, j)
				t.Fail()
			}
		}
	}
}

func TestRunCicle(t *testing.T) {
	m := [][]int{
		[]int{1, 1, 0},
		[]int{0, 1, 0},
		[]int{0, 1, 0},
	}
	f := &Field{m, 3, 3}

	f.runCicle()
	fmt.Println(f.String())
}
