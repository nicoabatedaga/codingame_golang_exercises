package main

import (
	"fmt"
	"math"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)
	var XF, YF int
	XF = W
	YF = H
	var X, Y int
	X = X0
	Y = Y0
	X0 = 0
	Y0 = 0

	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		fmt.Fprintln(os.Stderr, fmt.Sprintf("PUNTO X: %v - Y: %v", X, Y))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("X0:%v, Y0:%v", X0, Y0))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("XF:%v YF:%v", XF, YF))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("bombDir: %v", bombDir))
		X, Y, X0, Y0, XF, YF = nextJump(X, Y, X0, Y0, XF, YF, bombDir)

		// the location of the next window Batman should jump to.
		fmt.Println(fmt.Sprintf("%v %v", X, Y))
	}
}

/*
  0 1 2 3 4
0 * * * * *
1 * * * * *
2 * * X * *
3 * * * * *
4 * * * * *
*/
func nextJump(X, Y, X0, Y0, XF, YF int, direction string) (int, int, int, int, int, int) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("direction switch -> %v", direction))
	switch direction {
	case "D":
		/*
			  0 1 2 3 4
			0 * * * * *
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 * * 0 * *
		*/
		return X,
			numMedio(Y, YF),
			X,
			Y,
			X,
			YF
	case "DR":
		/*
			  0 1 2 3 4
			0 * * * * *
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 * * * * 0
		*/
		return numMedio(X, XF),
			numMedio(Y, YF),
			X,
			Y,
			XF,
			YF
	case "R":
		/*
			  0 1 2 3 4
			0 * * * * *
			1 * * * * *
			2 * * X * O
			3 * * * * *
			4 * * * * *
		*/
		return numMedio(X, XF),
			Y,
			X,
			Y,
			XF,
			Y
	case "UR":
		/*
			  0 1 2 3 4
			0 * * * * 0
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 * * * * *
		*/
		return numMedio(X, XF),
			numMedio(Y, Y0),
			X,
			Y0,
			XF,
			Y
	case "U":
		/*
			  0 1 2 3 4
			0 * * O * *
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 * * * * *
		*/
		return X,
			numMedio(Y, Y0),
			X,
			Y0,
			X,
			Y
	case "UL":
		/*
			  0 1 2 3 4
			0 0 * * * *
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 * * * * *
		*/
		return numMedio(X, X0),
			numMedio(Y, Y0),
			X0,
			Y0,
			X,
			Y
	case "L":
		/*
			  0 1 2 3 4
			0 * * * * *
			1 * * * * *
			2 0 * X * *
			3 * * * * *
			4 * * * * *
		*/
		return numMedio(X, X0),
			Y,
			X0,
			Y,
			X,
			Y
	case "DL":
		/*
			  0 1 2 3 4
			0 * * * * *
			1 * * * * *
			2 * * X * *
			3 * * * * *
			4 0 * * * *
		*/
		return numMedio(X, X0),
			numMedio(Y, YF),
			X0,
			Y,
			X,
			YF
	default:
		return X,
			Y,
			X,
			Y,
			X,
			Y

	}
}

func numMedio(a, b int) int {
	var m int
	if a == b {
		if a == 0 {
			return 0
		} else {
			return a - 1
		}
	}
	if b == 0 {
		m = int(math.Trunc(float64(a / 2)))
	} else {
		m = int(math.Round(float64((a + b) / 2)))
	}
	fmt.Fprintln(os.Stderr, fmt.Sprintf("medio de %v %v: %v", a, b, m))
	return m
}
