// Copyright 2017 Michael Meyling. All rights reserved.

package field

import (
	"fmt"
)

type Field struct {
  	Rows, Columns int
  	Matrix        [][]byte
}

const EMPTY = 32

func New(rows int, columns int) *Field {
	var f Field
	matrix := make([][]byte, rows, columns)
	for i := 0; i < rows; i++ {
		m := make([]byte, columns)
		for j := 0; j < columns; j++ {
			m[j] = EMPTY
		}
		matrix[i] = m
	}
	f.Rows = rows
	f.Columns = columns
	f.Matrix = matrix
	Frame(&f)
	return &f
}

func VLine(f *Field, x, y1, y2 int, v byte) {
	if y1 > y2 {
		return
	}
	for y := y1; y <= y2; y++ {
		if InField(f, x, y) {
			Set(f, x, y, v)
		}
	}
}

func HLine(f *Field, y, x1, x2 int, v byte) {
	if x1 > x2 {
		return
	}
	for x := x1; x <= x2; x++ {
		if InField(f, x, y) {
			Set(f, x, y, v)
		}
	}
}

func Frame(f *Field) {
	Square(f, 0, 0, f.Columns - 1, f.Rows - 1)
}

func Square(f *Field, x1, y1, x2, y2 int) {
	HLine(f, y1, x1 + 1, x2 - 1, '-')
	HLine(f, y2, x1 + 1, x2 - 1, '-')
	VLine(f, x1, y1 + 1, y2 - 1, '|')
	VLine(f, x2, y1 + 1, y2 - 1, '|')
	Set(f, x1, y1,  '+')
	Set(f, x1, y2,  '+')
	Set(f, x2, y1,  '+')
	Set(f, x2, y2,  '+')
}

func Get(f *Field, x, y int) byte {
	if NotInField(f, x, y) {
		panic(fmt.Sprintf("you tried to get at (%d, %d)", x, y))
	}
	return f.Matrix[f.Rows- 1 - y ][x]
}

func Set(f *Field, x, y int, v byte) {
	if NotInField(f, x, y) {
		panic(fmt.Sprintf("you tried to write at (%d, %d)", x, y))
	}
	f.Matrix[f.Rows- 1 - y][x] = v
}

func NotInField(f *Field, x, y int) bool {
	return x >= f.Columns || y >= f.Rows || x < 0 || y < 0
}

func InField(f *Field, x, y int) bool {
	return !NotInField(f, x, y)
}

func Empty(f *Field, x, y int) bool {
	if NotInField(f, x, y) {
		return false
	}
	return Get(f, x, y) == EMPTY
}

// draw field
func Draw(f *Field) {
	for i := 0; i < len(f.Matrix); i++ {
       	for j :=0 ; j < len(f.Matrix[i]); j++ {
		   fmt.Print(string(f.Matrix[i][j]));
	   	}
		fmt.Println()
	}
}
