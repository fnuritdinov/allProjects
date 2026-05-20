package main

// 1
/*

func main() {
	var str string
	fmt.Scan(&str)
	intSt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка ввода", err)
		return
	}
	if intSt < 0 {
		fmt.Println("Отрицательное число")
		return
	}
	fmt.Println(intSt * intSt)
}
*/

// 2
/*
func main() {
	var num float64
	fmt.Scan(&num)
	a := num
	b := math.Floor(a)
	fmt.Println(a - b)
}
*/

// 3
/*
func main() {
	var str string
	fmt.Scan(&str)

	strb := []byte(str)
	for _, b := range strb {
		b = b + 1
		fmt.Println(b, string(b))
	}

}
*/

// 4
/*
func main() {
	var str string
	fmt.Scan(&str)

	vowels := "aeiouAEIOU"

	for _, v := range vowels {
		str = strings.ReplaceAll(str, string(v), "//")
	}

	fmt.Println(str)
}

*/

// 5
/*
func main() {
	var str string
	fmt.Scan(&str)

	for _, value := range str {
		if !unicode.IsDigit(value) {
			fmt.Println("Not a digit")
			return
		}

	}
	fmt.Println("Число: " + str)
}
*/

// 8
/*
func main() {
	var str string
	fmt.Scan(&str)

	countInt := 0
	countString := 0
	strRune := []rune(str)
	for _, value := range strRune {
		if unicode.IsDigit(value) {
			countInt++
		} else if unicode.IsLetter(value) {
			countString++
		}
	}
	fmt.Printf("Цифры: %d, Буквы: %d", countInt, countString)
}
*/
