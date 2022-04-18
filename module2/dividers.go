package main

import (
	"fmt"
	"math"
	"sort"
)

func checkIfCorrect(dividers []int, u, w, v int) bool {
	return (dividers[u]%dividers[w] == 0 && dividers[w]%dividers[v] == 0)
}
func main() {
	var (
		n        int
		dividers []int = make([]int, 0)
	)

	fmt.Scan(&n)

	var limit int = int(math.Sqrt(float64(n)))

	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			dividers = append(dividers, i)
			if Exist(dividers, n/i) == -1 {
				dividers = append(dividers, n/i)
			}
		}
	}

	sort.Ints(dividers)

	fmt.Println("graph {")

	for i := len(dividers) - 1; i >= 0; i-- {
		fmt.Printf("\t%d\n", dividers[i])
	}

	for u := len(dividers) - 1; u > 0; u-- {
		for v := u - 1; v >= 0; v-- {
			if dividers[u]%dividers[v] == 0 {
				var breaker bool = true

				for w := u - 1; w > v; w-- {
					if checkIfCorrect(dividers, u, w, v) {
						breaker = false
						break
					}
				}

				if breaker {
					fmt.Printf("\t%d -- %d\n", dividers[u], dividers[v])
				}
			}
		}
	}

	fmt.Println("}")

}

func Exist(mas []int, x int) int {
	for i := 0; i < len(mas); i++ {
		if mas[i] == x {
			return mas[i]
		}
	}
	return -1
}
