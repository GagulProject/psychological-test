package main

import (
	"fmt"
	"sort"
)

func main() {
	sumCharacter("O", "A")
}

// 두 문자 를 정렬 후 합치는 함수
func sumCharacter(leftChromosome string, rightChromosome string) {
	chromosome := []string{leftChromosome, rightChromosome}
	fmt.Println("before sort chromosome : ", chromosome)
	sort.Strings(chromosome)
	fmt.Println("after sort chromosome : ", chromosome)

	retChromosome := chromosome[0] + chromosome[1]

	fmt.Println("sum string", retChromosome)
}
