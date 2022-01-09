package main

import "fmt"
import "os"

/**
 * The while loop represents the game.
 * Each iteration represents a turn of the game
 * where you are given inputs (the heights of the mountains)
 * and where you have to print an output (the index of the mountain to fire on)
 * The inputs you are given are automatically updated according to your last actions.
 **/

func main() {
	for {
		ordenParaAplicarMafia := map[int]int{}
		for i := 0; i < 8; i++ {
			// mountainH: represents the height of one mountain.
			var mountainH int
			fmt.Scan(&mountainH)
			ordenParaAplicarMafia[i] = mountainH
		}
		fmt.Fprintln(os.Stderr, ordenParaAplicarMafia)

		fmt.Println(masAlta(ordenParaAplicarMafia)) // The index of the mountain to fire on.
	}
}

func masAlta(mapa map[int]int) int {
	x := 0
	for montania, alto := range mapa {
		if mapa[x] < alto {
			x = montania
		}
	}
	return x
}
