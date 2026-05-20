package main

import "fmt"

func main() {
	countPr := map[string]int{}
	sales := []string{"phone", "laptop", "phone", "tablet", "phone", "laptop"}
	for _, phone := range sales {
		countPr[phone]++
	}
	fmt.Println(countPr)
}
