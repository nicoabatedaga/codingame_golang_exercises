package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
	// lightX: the X position of the light of power
	// lightY: the Y position of the light of power
	// initialTx: Thor's starting X position
	// initialTy: Thor's starting Y position
	var lightX, lightY, initialTx, initialTy int
	fmt.Scan(&lightX, &lightY, &initialTx, &initialTy)

	for {
		// remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
		var remainingTurns int
		fmt.Scan(&remainingTurns)

		var coord string
		initialTx, initialTy, coord = nextStep(lightX, lightY, initialTx, initialTy)
		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// A single line providing the move to be made: N NE E SE S SW W or NW
		fmt.Println(coord)
	}
}

/*
	  0 1 2 3 4
	0 * * N * *
	1 * * * * *
	2 W * X * E
	3 * * * * *
	4 * * S * *
*/
func nextStep(lightX, lightY, x, y int) (int, int, string) {
	nextCoord := ""
	if lightY != y {
		if lightY > y {
			nextCoord = nextCoord + "S"
		} else {
			nextCoord = nextCoord + "N"
		}
	}
	if lightX != x {
		if lightX > x {
			nextCoord = nextCoord + "E"
		} else {
			nextCoord = nextCoord + "W"
		}
	}
	switch nextCoord {
	case "N":
		return x, y - 1, nextCoord
	case "S":
		return x, y + 1, nextCoord
	case "E":
		return x + 1, y, nextCoord
	case "W":
		return x - 1, y, nextCoord
	case "NE":
		return x + 1, y - 1, nextCoord
	case "NW":
		return x - 1, y - 1, nextCoord
	case "SE":
		return x + 1, y + 1, nextCoord
	case "SW":
		return x - 1, y + 1, nextCoord
	default:
		return 0, 0, ""
	}
}
