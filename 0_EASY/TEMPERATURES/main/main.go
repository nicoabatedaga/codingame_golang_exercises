package main

import (
	"fmt"
	"math"
)
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	list := []int64{}
	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		list = append(list, t)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(cercanoAcero(list)) // Write answer to stdout
}

func cercanoAcero(list []int64) int64 {
	if len(list) == 0 {
		return 0
	}
	cercano := int64(9999)
	for _, temp := range list {
		//fmt.Print("iteracion", i)
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Comparo cercano: %v con temp: %v", cercano, temp))
		if math.Abs(float64(cercano)) >= math.Abs(float64(temp)) {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("Comparo ABS cercano: %v con temp: %v", math.Abs(float64(cercano)), math.Abs(float64(temp))))
			if math.Abs(float64(cercano)) == math.Abs(float64(temp)) {
				if cercano < 0 {
					cercano = temp
				}
			} else {
				cercano = temp
			}
		}
	}
	return cercano
}
