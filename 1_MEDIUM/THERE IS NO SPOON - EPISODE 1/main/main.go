package main

import (
	"fmt"
)
import "os"
import "bufio"

/**
 * Don't let the machines win. You are humanity's last hope...
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: the number of cells on the X axis
	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	// height: the number of cells on the Y axis
	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)
	matrix := [30][30]string{}

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j, c := range line {
			matrix[j][i] = string(c)
		}
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("%#v", matrix))
	fmt.Fprintln(os.Stderr, width, height)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if matrix[i][j] == "0" {
				printNodeNeighbor(matrix, i, j, width, height)
			}
		}
	}
}

func printNodeNeighbor(matrix [30][30]string, x int, y int, width int, height int) {
	xr := -1
	yr := -1
	xb := -1
	yb := -1
	for i := x + 1; i < width; i++ {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("chequeo x:%v i:%v %v", x, i, matrix[i][y]))
		if matrix[i][y] == "0" {
			xr = i
			yr = y
			i = width
		}
	}
	for i := y + 1; i < height; i++ {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("chequeo y:%v i:%v %v", y, i, matrix[x][i]))
		if matrix[x][i] == "0" {
			xb = x
			yb = i
			i = height
		}
	}
	fmt.Println(fmt.Sprintf("%v %v %v %v %v %v", x, y, xr, yr, xb, yb))
}
