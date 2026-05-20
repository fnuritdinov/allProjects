package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
func main() {

	x := 10
	y := &x
	fmt.Println(y, *y)
}
*/

/*
func main() {
	x := 5
	y := &x
	*y = 20
	fmt.Println(x)
}
*/

/*
func main() {
	x := 6
	Double(&x)
	fmt.Println(x)

}

func Double(x *int) {
	*x = 2 * *x
}
*/

/*
func main() {
	a, b := 10, 20
	Swap(&a, &b)
	fmt.Println(a, b)

}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
*/

/*
func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		s := &nums[i]
		*s += 10
	}
	fmt.Println(nums)
}
*/

// Home work

// 1
/*
func main() {
	a := 5
	Increment(&a)
	fmt.Println(a)
}

func Increment(a *int) {
	*a++
}
*/

// 2
/*
func main() {
	a := 8
	Zero(&a)
	fmt.Println(a)
}

func Zero(a *int) {
	*a = *a - *a
}
*/

//3
/*
func main() {
	a, b := 3, 4
	fmt.Println(Max(&a, &b))

}

func Max(a *int, b *int) int {
	max1 := *a
	if *b > max1 {
		max1 = *b
	}
	return max1
}
*/

// 4
/*
func main() {
	a, b := 9, 8
	ReplaceIfGreater(&a, &b)
	fmt.Println(a, b)

}

func ReplaceIfGreater(a *int, b *int) {
	if *a > *b {
		*b = *a
	}
}
*/

// 5
/*
func main() {
	nums := []int{2, 4, 6, 8, 9}
	for i := 0; i < len(nums); i++ {
		s := &nums[i]
		if *s%2 == 0 {
			*s += 5
		}

	}
	fmt.Println(nums)

}
*/

// 6
/*
func main() {
	a := []int{1, 2, 3, 4}
	swapFirstLast(a)
	fmt.Println(a)

}

func swapFirstLast(nums []int) {
	first := &nums[0]
	last := &nums[len(nums)-1]
	*first, *last = *last, *first

}
*/

// 7
/*
func main() {
	a := 3
	b := 9
	result := a + b
	Sum(&a, &b, &result)
	fmt.Println(result)

}

func Sum(a *int, b *int, result *int) {
	*result = *a + *b
}

*/
func main() {
	numb := adding("jjjjk34", "jijijij56")
	fmt.Println(numb)
}

func adding(a, b string) int64 {
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	aRune := []rune(a)
	bRune := []rune(b)

	var numA int64
	for _, valueA := range aRune {
		if unicode.IsDigit(valueA) {
			A, err := strconv.Atoi(string(valueA))
			if err != nil {
				fmt.Println(err)
			}
			numA += int64(A)

		}
	}
	var numB int64
	for _, valueB := range bRune {
		if unicode.IsDigit(valueB) {
			B, err := strconv.Atoi(string(valueB))
			if err == nil {
				fmt.Println(err)
			}
			numB += int64(B)
		}
	}
	return numA + numB
}
