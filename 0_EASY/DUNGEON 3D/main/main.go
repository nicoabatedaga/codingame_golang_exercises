package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
var L, R, C, Ax, Ay, Az, Sx, Sy, Sz int
var Matrix [30][30][30]string
var minutes []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	fillObject(scanner)

	calculateTimeToGoal(Ax, Ay, Az)

	if len(minutes) == 0 {
		fmt.Println("NO PATH") // Write answer to stdout
	} else {
		sort.Ints(minutes)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("List of minutes: %v", minutes))
		fmt.Println(minutes[0]) // Write answer to stdout
	}
}

func recorrePath(x, y, z, px, py, pz, auxmin int) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("iteracion x:%v y:%v z:%v", x, y, z))
	if x == Sx && y == Sy && z == Sz {
		fmt.Fprintln(os.Stderr, "encontre la ubicacion")
		minutes = append(minutes, auxmin)
		return
	}
	//x==px && y==py && z==pz
	Matrix[z][y][x] = "*"
	if esPisoOpersona(z, y, x+1) {
		recorrePath(x+1, y, z, x, y, z, auxmin+1)
	}
	if esPisoOpersona(z, y, x-1) {
		recorrePath(x-1, y, z, x, y, z, auxmin+1)
	}
	if esPisoOpersona(z, y+1, x) {
		recorrePath(x, y+1, z, x, y, z, auxmin+1)
	}
	if esPisoOpersona(z, y-1, x) {
		recorrePath(x, y-1, z, x, y, z, auxmin+1)
	}
	if esPisoOpersona(z+1, y, x) {
		recorrePath(x, y, z+1, x, y, z, auxmin+1)
	}
	if esPisoOpersona(z-1, y, x) {
		recorrePath(x, y, z-1, x, y, z, auxmin+1)
	}
	Matrix[z][y][x] = "."
}

func calculateTimeToGoal(x0, y0, z0 int) {
	recorrePath(x0, y0, z0, x0, y0, z0, 0)
}

func esPisoOpersona(z, y, x int) bool {
	return (x >= 0 && x < C && y >= 0 && y < R && z >= 0 && z < L) && (Matrix[z][y][x] == "." || Matrix[z][y][x] == "S")
}

func fillObject(scanner *bufio.Scanner) {
	// L es la cantidad de niveles
	// R es la cantidad de row (files)
	// C cantidad de columnas
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L, &R, &C)

	var ln int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &ln)

	// Itero por los pisos
	for piso := 0; piso < L; piso++ {
		for fila := 0; fila < R; fila++ {
			scanner.Scan()
			line := scanner.Text()
			if len(line) > 0 {
				for columna, value := range line {
					switch string(value) {
					case "S":
						Sx = columna
						Sy = fila
						Sz = piso
						Matrix[piso][fila][columna] = string(value)

					case "A":
						Ax = columna
						Ay = fila
						Az = piso
						Matrix[piso][fila][columna] = "."

					default:
						Matrix[piso][fila][columna] = string(value)
					}
				}
			} else {
				fila--
			}
		}
	}

	//printMatrix(matrix, R, C, L)
}

func printMatrix() {
	for i := 0; i < L; i++ {
		fmt.Println("level ", i)
		for j := 0; j < R; j++ {
			linea := ""
			for ii := 0; ii < C; ii++ {
				linea = fmt.Sprintf("%s%s", linea, Matrix[i][j][ii])
			}
			fmt.Println(linea)
		}
	}
}
