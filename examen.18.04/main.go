package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Grade int
}

func main() {
	var name string
	fmt.Scan(&name)

	file, err := os.OpenFile("student.json", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
	}

	defer file.Close()

	var students []Student
	mapa := make(map[string]int)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&students)

	for _, value := range students {
		mapa[value.Name] = value.Grade
		fmt.Println(mapa)

	}

}
