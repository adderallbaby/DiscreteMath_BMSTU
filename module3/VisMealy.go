package main

import "fmt"
func ratatatata(exit [][]string, crossover [][]int, start int)  {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")


	
	for i, _ := range crossover {
		for j, x := range crossover[i] {
			fmt.Printf("%d -> %d [label = \"%s(%s)\"]\n", i, x, string(j + 97), exit[i][j])
		}
	}
	fmt.Println("}")
}
func main() {
	var n, m, start, cr int
	var sym string
	_, _ = fmt.Scan(&n)
	crossover := make([][]int, n)
	exit := make([][]string, n)
	_, _ = fmt.Scan(&m)
	_, _ = fmt.Scan(&start)
	for i := 0; i < n; i++ {
		crossover[i] = make([]int, 0)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&cr)
			crossover[i] = append(crossover[i], cr)
		}
 	}
	for i := 0; i < n; i++ {
		exit[i] = make([]string, 0)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&sym)
			exit[i] = append(exit[i], sym)
		}
	}
	ratatatata(exit, crossover, start)
 }
