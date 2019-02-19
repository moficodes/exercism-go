package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type mat struct {
	data [][]int
}

// New returns a new matrix generated from the string and an error
func New(data string) (*mat, error) {
	lines := strings.Split(data, "\n")

	if len(lines) == 0 {
		return nil, errors.New("no data")
	}
	m := make([][]int, len(lines))
	var rowLen int
	for i, line := range lines {
		numbers := strings.Split(strings.TrimSpace(line), " ")
		if i == 0 {
			rowLen = len(numbers)
		}
		if rowLen != len(numbers) {
			return nil, errors.New("col length mismatch")
		}
		col := make([]int, rowLen)
		for j, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			col[j] = val
		}
		m[i] = col
	}
	return &mat{data: m}, nil
}

func (m *mat) Rows() [][]int {
	r := make([][]int, len(m.data))
	for i, row := range m.data {
		r[i] = make([]int, len(row))
		for j, col := range row {
			r[i][j] = col
		}
	}
	return r
}

func (m *mat) Cols() [][]int {
	c := make([][]int, len(m.data[0]))

	for i := range c {
		c[i] = make([]int, len(m.data))
	}

	for i, row := range m.data {
		for j, col := range row {
			c[j][i] = col
		}
	}
	return c
}

func (m *mat) Set(row, col, val int) bool {
	if row >= len(m.data) || row < 0 {
		return false
	}
	if col >= len(m.data[0]) || col < 0 {
		return false
	}
	m.data[row][col] = val
	return true
}
