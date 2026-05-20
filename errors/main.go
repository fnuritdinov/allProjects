package main

// home work
// 1
/*
func main() {
	var str string
	fmt.Scanln(&str)

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
*/

// 2
/*
func main() {
	var num int
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println("Ошибка: отрицательное число")
	} else {
		fmt.Println(num * num)
	}

}
*/

//3
/*
func main() {

	var n int
	fmt.Scan(&n)
	f := checkEven(n)
	fmt.Println(f)

}

func checkEven(n int) error {
	err := errors.New("Число нечетное")

	if n%2 != 0 {
		return err
	}
	return nil
}
*/

// 4
/*
func main() {

	fmt.Println(divide(10, 0))
}

func divide(a, b int) (int, error) {
	err := errors.New("На ноль делить нельзя")
	if b > 0 {
		return a / b, nil
	} else {
		return 0, err
	}
}
*/

// 5
/*
func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 5
	num, err := find(arr, target)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

}

func find(arr []int, target int) (int, error) {
	err := errors.New("Элемент не найден в слайсе")
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, nil
		}
	}
	return -1, err
}
*/

// 6
/*
func main() {

	err := validateName("we")
	if err != nil {
		fmt.Println(err)
	}

}

func validateName(name string) error {
	nameRune := []rune(name)
	if len(nameRune) == ' ' || len(nameRune) < 3 {
		return errors.New("invalid name")
	}
	fmt.Println(string(nameRune))
	return nil
}
*/

// 7
/*
func main() {
	fmt.Println(register(9, ""))
}

func register(age int, name string) error {
	err := errors.New("Доступ запрещен")
	err1 := errors.New("Имя обязательное")

	if age < 18 {
		return err
	}
	if name == "" {
		return err1
	}
	fmt.Println(age, name)
	return nil

}
*/

// 8
/*
func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Ошибка при чтении строки")
		return
	}

	intStr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка при приоброазовании в число")
		return
	}

	if intStr <= 0 {
		fmt.Println("Ошибка: число должно быть положительным")
		return
	}

	if intStr%2 != 0 {
		fmt.Println("Ошибка: Число нечетное")
		return
	}
	fmt.Println("OK")
}
*/

// 9
/*
func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := range array {
		fmt.Scan(&array[i])
	}
	in, err := sumPositive(array)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(in)
}

func sumPositive(arr []int) (int, error) {
	count := 0
	for _, value := range arr {
		if value < 0 {
			return 0, errors.New("Число отрицательное")
		}
		count += value
	}

	return count, nil
}
*/

//10
/*
func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(safeSqrt(num))

}

func safeSqrt(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("Ошибка: нельзя извлечь корень из отрицательного числа")
	}
	return math.Sqrt(float64(n)), nil

}
*/

// 12
/*
func main() {
	num, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Errorf("Ошибка при преобразовании %v", err)
		return
	}
	fmt.Println(num)

}

*/
