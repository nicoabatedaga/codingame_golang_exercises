package main

import "fmt"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	alto, ancho, object := fillObject(scanner)
	altoGrid, anchoGrid, grid := fillObject(scanner)

	_, _, _ = alto, ancho, object
	_, _, _ = altoGrid, anchoGrid, grid
	//fmt.Println(alto, ancho, object)
	//fmt.Println(altoGrid, anchoGrid, grid)

	count := putObjectInGrid(&object, &grid, alto, ancho, altoGrid, anchoGrid)

	//fmt.Println("-------") // Write answer to stdout
	//

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(count) // Write answer to stdout
	//printMatrix(object, alto, ancho)
	if count == 1 {
		printMatrix(grid, altoGrid, anchoGrid)
	}
}

func printMatrix(matrix [10][10]int, alto int, ancho int) {
	for i := 0; i < alto; i++ {
		linea := ""
		for j := 0; j < ancho; j++ {
			switch matrix[i][j] {
			case 0:
				linea = fmt.Sprintf("%s.", linea)
			case 1:
				linea = fmt.Sprintf("%s#", linea)
			case 2:
				linea = fmt.Sprintf("%s*", linea)
			}
		}
		fmt.Println(linea)
	}
}

func fillObject(scanner *bufio.Scanner) (int, int, [10][10]int) {
	var a, b int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &a, &b)
	matrix := [10][10]int{}
	for i := 0; i < a; i++ {
		scanner.Scan()
		matrixLine := scanner.Text()
		for j := 0; j < b; j++ {
			var dato int
			switch matrixLine[j] {
			case '.':
				dato = 0
			case '#':
				dato = 1
			case '*':
				dato = 2
			}
			matrix[i][j] = dato
		}
	}
	//printMatrix(matrix, a, b)
	return a, b, matrix
}

func putObjectInGrid(objeto, grid *[10][10]int, alto, ancho, altogrid, anchogrid int) int {
	count := 0
	for i := 0; i < altogrid; i++ {
		for j := 0; j < anchogrid; j++ {
			entra := entraObjeto(*objeto, *grid, alto, ancho, altogrid, anchogrid, i, j)
			if entra {
				//fmt.Println("entra", i, j)
				count++
				for x := 0; x < alto; x++ {
					for y := 0; y < ancho; y++ {
						if objeto[x][y] == 2 {
							grid[i+x][j+y] = 2
						}
					}
				}
			}
		}
	}
	return count
}

func entraObjeto(objeto, grid [10][10]int, altoobjeto, anchoobjeto, altogrid, anchogrid, pii, pij int) bool {
	if pii+altoobjeto > altogrid || pij+anchoobjeto > anchogrid {
		return false
	}
	for i := 0; i < altoobjeto; i++ {
		for j := 0; j < anchoobjeto; j++ {
			gaux := grid[pii+i][pij+j]
			oaux := objeto[i][j]
			if (gaux == 1 /* || gaux == 2*/) && oaux == 2 {
				//fmt.Println(fmt.Sprintf("gaux:%v, oaux:%v", gaux, oaux))
				return false
			}
		}
	}
	return true
}
