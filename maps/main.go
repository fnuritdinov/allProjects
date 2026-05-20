package main

import "fmt"

func main() {
	/*words := []string{"go", "java", "go", "python", "go", "java"}
	mapa := make(map[string]int)
	for _, word := range words {
		mapa[word]++

	}
	fmt.Println(mapa)

	*/
	/*
		scores := map[string]int{
			"Alice": 80,
			"Bob":   95,
			"John":  90,
		}
		maxScore := 0
		winner := ""
		for key, value := range scores {
			if value > maxScore {
				maxScore = value
				winner = key

			}
		}
		fmt.Println(winner, maxScore)

	*/
	/*
		products := map[string]int{
			"phone":  500,
			"laptop": 1500,
			"mouse":  50,
			"tablet": 700,
		}
		for key, value := range products {
			if value > 600 {
				fmt.Println(key, value)
			}
		}

	*/
	/*
		names := []string{"Ali", "Azam", "Bob", "Ben", "Charlie"}
		mp := make(map[string][]string)
		mp["names"] = names
		for k, v := range mp {
			fmt.Println(k, v)
		}

	*/

	str := "hello"

	mapa := make(map[rune]int)
	for _, v := range str {
		mapa[v]++
	}

	for k, v := range mapa {
		fmt.Printf("%c %d\n", k, v)
	}
}
