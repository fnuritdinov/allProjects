package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(IsPalindrome(n))
}

func IsPalindrome(n int) string {
	a := n / 1000
	b := (n / 100) % 10
	c := (n / 10) % 10
	d := n % 10
	if a+b == c+d {
		return "YES"
	} else {
		return "NO"

	}

}
