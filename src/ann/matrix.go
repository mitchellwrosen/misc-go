package main

import (
	"fmt"
)

type Matrix struct {
	rows, cols int
	data       [][]float32
}

func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float32, cols)
	}

	return &Matrix{rows, cols, data}
}

func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%f ", m.data[i][j])
		}
		fmt.Println()
	}
}

func (m *Matrix) SetData(data [][]float32) {
	m.data = data
}

type elem struct {
	row, col int
	val      float32
}

func dotMultiply(m1, m2 *Matrix, row, col int, ch chan elem) {
	if m1.cols != m2.rows {
		panic(fmt.Sprintf("%dx%d X %dx%d", m1.rows, m1.cols, m2.rows, m2.cols))
	}

	sum := float32(0)
	for i := 0; i < m1.cols; i++ {
		sum += m1.data[row][i] * m2.data[i][col]
	}

	ch <- elem{row, col, sum}
}

func (m *Matrix) Times(m2 *Matrix) *Matrix {
	if m.cols != m2.rows {
		panic(fmt.Sprintf("%dx%d X %dx%d", m.rows, m.cols, m2.rows, m2.cols))
	}

	ch := make(chan elem, m.rows*m2.cols)

	prod := NewMatrix(m.rows, m2.cols)
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m2.cols; col++ {
			go dotMultiply(m, m2, row, col, ch)
		}
	}

	for i := 0; i < m.rows*m2.cols; i++ {
		e := <-ch
		prod.data[e.row][e.col] = e.val
	}

	return prod
}

func main() {
	m1 := NewMatrix(3, 2)
	m1.SetData([][]float32{[]float32{8, 1}, []float32{2, 2}, []float32{3, 3}})
	m1.Print()
	fmt.Println()

	m2 := NewMatrix(2, 1)
	m2.SetData([][]float32{[]float32{1}, []float32{2}})
	m2.Print()
	fmt.Println()

	m3 := m1.Times(m2)
	m3.Print()
}
