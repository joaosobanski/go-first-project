package main

import (
	"fmt"
)

func main() {
	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
	pontosPlanningPoker = append(pontosPlanningPoker, 40)
	fmt.Println((pontosPlanningPoker))
	fmt.Println(len(pontosPlanningPoker))
	// fmt.Println((pontosPlanningPoker[8]))
	for i := 0; i < len(pontosPlanningPoker); i++ {
		fmt.Println(pontosPlanningPoker[i])
	}
	for i, ponto := range pontosPlanningPoker {
		fmt.Println(i, ponto)
	}
}
