package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// home work

// 1
/*
func main() {
	now := time.Now()

	fmt.Printf("Год: %d\n месяц: %d\n день: %d\n час: %d\n минуты: %d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
}

*/

// 2
/*
func main() {
	now := time.Now()

	tm := now.Format("02.01.2006 15:04")
	fmt.Println(tm)
}

*/

// 3
/*
func main() {
	now := time.Now()

	tm := now.Weekday()


	if tm == time.Saturday || tm == time.Sunday {
		fmt.Println("Выходной")
	} else {
		fmt.Println("Рабочий день")
	}
	fmt.Println(tm)
}

*/

// 4
/*
func main() {
	var tm int
	fmt.Scan(&tm)

	now := time.Now()

	newTime := now.AddDate(0, 0, tm)

	fmt.Printf("Дата через %d дней: %v", tm, newTime.Format("02-01-2006 15:04"))
}

*/

// 5
/*
func main() {
	birthday := time.Date(1994, time.April, 10, 12, 05, 0, 0, time.UTC)
	now := time.Now()

	diff := now.Sub(birthday)
	fmt.Println(diff.Hours() / 24)
}

*/

// 6
/*

func main() {
	date1 := time.Date(2026, time.May, 8, 16, 24, 0, 0, time.UTC)
	date2 := time.Date(2026, time.May, 1, 16, 24, 0, 0, time.UTC)

	fmt.Println(date1.After(date2))
	fmt.Println(date2.Before(date1))
	fmt.Println(!date1.Equal(date2))

}

*/
// 7
/*

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	tm, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Год:", tm.Year())
	fmt.Println("Месяц:", tm.Month())
	fmt.Println("День:", tm.Day())
}

*/
func main() {
	os.OpenFile()
	file, err := os.Open("grades.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
		return
	}
	defer file.Close()

	mp := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		gradesInfo := strings.Fields(line)

		if len(gradesInfo) != 2 {
			continue
		}

		name := gradesInfo[0]

		grade, err := strconv.Atoi(name)
		if err != nil {
			fmt.Println("Ошибка во время прсинга", err)
			continue
		}

		mp[name] = grade

	}

	for key, value := range mp {
		if value >= 60 {
			fmt.Printf(`Имя: %s возраст: %d`, key, value)
		}
	}

}
