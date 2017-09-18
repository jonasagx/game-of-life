package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	sleepInterval = 300 * time.Millisecond
	seed          = time.Now().Unix()
)

func NewField(width, height int) *Field {
	return &Field{width: width, height: height}
}

type Field struct {
	matrix        [][]int
	width, height int
}

func (f *Field) Populate() {
	rand.Seed(seed)

	for i := 0; i < f.width; i++ {
		row := make([]int, f.width)

		for j := 0; j < f.height; j++ {
			row[j] = rand.Intn(2)
		}
		f.matrix = append(f.matrix, row)
	}
}

func (f *Field) String() string {
	s := ""

	for i := 0; i < f.width; i++ {
		for j := 0; j < f.width; j++ {
			if f.matrix[i][j] == 1 {
				s += "# "
			} else {
				s += ". "
			}
		}
		s += "\n"
	}

	return s
}

func (f *Field) countNeighbours(i, j int) int {
	var n int

	f.checkCellAndAdd(i-1, j-1, &n)
	f.checkCellAndAdd(i-1, j, &n)
	f.checkCellAndAdd(i-1, j+1, &n)
	f.checkCellAndAdd(i, j-1, &n)
	f.checkCellAndAdd(i, j+1, &n)
	f.checkCellAndAdd(i+1, j-1, &n)
	f.checkCellAndAdd(i+1, j, &n)
	f.checkCellAndAdd(i+1, j+1, &n)

	return n

}

func (f *Field) checkCellAndAdd(i, j int, c *int) {
	// fmt.Printf("checking [%d, %d]\n", i, j)

	if i < 0 || j < 0 {
		return
	}

	if i >= f.width || j >= f.height {
		return
	}

	// fmt.Printf("add %d at [%d, %d]\n", f.matrix[i][j], i, j)
	*c += f.matrix[i][j]
}

// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
// Any live cell with two or three live neighbours lives on to the next generation.
// Any live cell with more than three live neighbours dies, as if by overpopulation.
// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func (f *Field) runCicle() {
	var newField [][]int

	for i := 0; i < f.width; i++ {
		newRow := make([]int, f.width)

		for j := 0; j < f.height; j++ {
			c := f.countNeighbours(i, j)

			if f.matrix[i][j] == 1 && c < 2 {
				// f.matrix[i][j] = 0
				newRow[j] = 0
				// fmt.Printf("killing cell [%d %d] - c: %d\n", i, j, c)
				continue
			}

			if f.matrix[i][j] == 1 && (c == 3 || c == 2) {
				newRow[j] = 1
				// fmt.Printf("keeping cell [%d %d] - c: %d\n", i, j, c)
				continue
			}

			if f.matrix[i][j] == 1 && c > 3 {
				// f.matrix[i][j] = 0
				newRow[j] = 0
				// fmt.Printf("killing cell [%d %d] - c: %d\n", i, j, c)
				continue
			}

			if f.matrix[i][j] == 0 && c > 3 {
				// f.matrix[i][j] = 1
				newRow[j] = 1
				// fmt.Printf("new cell [%d %d] - c: %d\n", i, j, c)
				continue
			}
		}

		newField = append(newField, newRow)
	}

	f.matrix = newField
}

func (f *Field) Run(cicles int) {
	fmt.Println(f.String())
	var lastField string
	for i := 0; i < cicles; i++ {
		time.Sleep(sleepInterval)
		f.runCicle()
		fieldNow := f.String()
		fmt.Println(fieldNow)

		if fieldNow == lastField {
			fmt.Printf("field got into repited state after %d iterations\n", i)
			break
		}

		lastField = fieldNow
	}
}

func main() {
	f := NewField(30, 30)
	f.Populate()

	// fmt.Println(f.String())
	// f.runCicle()
	// fmt.Println(f.String())

	f.Run(30)
}
