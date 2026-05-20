package main

import "fmt"

func main() {
	// 7 Калькулятор с защитой
	/*var x, y float64
	var symbol string
	fmt.Scan(&x, &y, &symbol)

	switch symbol {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("На ноль делить нельзя")
			return
		}
		fmt.Println(x / y)
	default:
		fmt.Println("Неверный символ")
	}
	*/

	// 1 Калькулятор логистики
	/*var kg, km int
	fmt.Scan(&kg, &km)

	var base int
	var extra int

	if kg <= 5 {
		base = 500
	} else if kg <= 20 {
		base = 1000
	} else {
		base = 2000
	}

	if km > 100 {
		extra = (km - 100) * 5
	}

	fmt.Println(base + extra)

	*/

	// 2 Максимум из трех
	/*
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		var mx int
		mx = a
		if b > mx {
			mx = b
		}
		if c > mx {
			mx = c
		}
		fmt.Println(mx)
	*/

	// 3 Тип чисел
	/*
		var a, b, c, d, e int
		fmt.Scan(&a, &b, &c, &d, &e)
		positive := 0
		negative := 0
		zero := 0
		if a > 0 {
			positive++
		} else if a < 0 {
			negative++
		} else {
			zero++
		}
		if b > 0 {
			positive++
		} else if b < 0 {
			negative++
		} else {
			zero++
		}
		if c > 0 {
			positive++
		} else if c < 0 {
			negative++
		} else {
			zero++
		}
		if d > 0 {
			positive++
		} else if d < 0 {
			negative++
		} else {
			zero++
		}
		if e > 0 {
			positive++
		} else if e < 0 {
			negative++
		} else {
			zero++
		}
		fmt.Println(positive, negative, zero)

	*/

	// 4 Существование и тип треугольника
	/*
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a+b > c && a+c > b && c+b > a {
			if a == b && b == c {
				fmt.Println("Треугольник равностороний")
			} else if a == b || a == c || b == c {
				fmt.Println("Треугольник равнобедреный")
			} else {
				fmt.Println("Треугольник разностороний")
			}

		} else {
			fmt.Println("Треугольник не существует")

		}

	*/
	// 5 Калькулятор дней в месяце
	/*
		var month int
		fmt.Scan(&month)
		if month == 4 || month == 6 || month == 9 || month == 11 {
			fmt.Println("30 дней")
		} else if month == 2 {
			fmt.Println("28 дней")
		} else {
			fmt.Println("31 день")
		}

	*/

	// 6 Время суток
	/*
		var hour int
		fmt.Scan(&hour)
		if hour > 0 && hour <= 5 {
			fmt.Println("Ночь")
		} else if hour > 6 && hour <= 11 {
			fmt.Println("Утро")
		} else if hour >= 12 && hour <= 17 {
			fmt.Println("День")
		} else if hour >= 18 && hour < 23 {
			fmt.Println("Ночь")
		} else {
			fmt.Println("Ошибка")
		}

	*/

	// 8 Белочка и шишки

	var N, M, K int
	fmt.Scan(&N, &M, &K)

	if N*M >= K {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
