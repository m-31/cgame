// Copyright 2017 Michael Meyling. All rights reserved.

package field

import (
	"fmt"
)

type Field struct {
  	Rows, Columns int
  	Matrix        [][]byte
}

func New(rows int, columns int) *Field {
	var f Field
	matrix := make([][]byte, rows, columns)
	for i := 0; i < rows; i++ {
		m := make([]byte, columns)
		for j := 0; j < columns; j++ {
			m[j] = 32
		}
		matrix[i] = m
	}
	f.Rows = rows
	f.Columns = columns
	f.Matrix = matrix
	return &f
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

func Empty(f *Field, x, y int) bool {
	if NotInField(f, x, y) {
		return false
	}
	return Get(f, x, y) == 32
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
