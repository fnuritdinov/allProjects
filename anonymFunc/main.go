package main

func main() {
	/*				//classwork

				fn := func(text string) {
					fmt.Println(text)
				}
				fn("Hello Go")

				func(a, b int) int {
					result := a + b
					fmt.Println(result)
					return result
				}(2, 3)



				check := func(n int) bool {
					if n%2 == 0 {
						fmt.Println("true")
						return true
					}
					fmt.Println("false")
					return false
				}
				check(3)



				slFun := func(nums []int) []int {
					var result []int
					for _, n := range nums {
						n = n * 2
						result = append(result, n)
					}
					fmt.Println(result)
					return nums
				}
				slFun([]int{1, 2, 3, 4, 5})
				}


			slMax := func(nums []int) int {
				maxNum := 0
				for _, n := range nums {
					if n > maxNum {
						maxNum = n
					}
				}

				return maxNum
			}
			fmt.Println(slMax([]int{1, 2, 3}))


		fn := func(nums []int) []int {
			slice := []int{}
			for _, n := range nums {
				if n%2 == 0 {
					slice = append(slice, n)
				}
			}
			return slice
		}

		fmt.Println(fn([]int{1, 2, 3, 4, 5}))


	count := 1
	fn := func() int {

		count = count + 1
		return count
	}

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	*/

	// homework
	// 1
	/*
		fn := func(pass string) bool {
			if len(pass) < 6 {
				return false
			}

			hasLetter := false
			hasDigit := false

			passRunes := []rune(pass)
			for _, r := range passRunes {
				if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
					hasLetter = true
				}
				if r >= '0' && r <= '9' {
					hasDigit = true
				}
			}
			return hasLetter && hasDigit
		}

		fmt.Println(fn("abc12"))

	*/
	// 2
	/*
		wr := func(words []string, num int) int {
			count := 0
			for _, word := range words {
				if len(word) > num {
					count++
				}
			}
			return count
		}
		fmt.Println(wr([]string{"go", "backend", "api", "database", "sql"}, 7))

	*/
	//3
	/*
		fn := func(sliceInt []int) []int {
			var sliceNil []int
			for _, value := range sliceInt {
				if value > 0 {
					sliceNil = append(sliceNil, value)
				}

			}
			return sliceNil
		}

		fmt.Println(fn([]int{4, -2, 7, -9, 0, 3}))

	*/
	// 4
	/*

		fn := func(sliceStr []string) []string {
			var sliceTT []string
			for _, str := range sliceStr {
				str = strings.ToLower(str)
				str = strings.TrimSpace(str)

				runeSlice := []rune(str)
				runeSlice[0] = unicode.ToUpper(runeSlice[0])
				sliceTT = append(sliceTT, string(runeSlice))
			}
			return sliceTT

		}

		fmt.Println(fn([]string{"ali", "UMED", " rustam ", "SaId"}))

	*/
	// 5
	/*
		func(str string) int {
			runes := []rune(str)
			count := 0
			for _, value := range runes {
				if unicode.IsDigit(value) {
					count += int(value - '0')

				}
			}
			fmt.Println(count)
			return count
		}("12")

	*/

	// 6

	// 7
	/*
		fn := func(a, b int, str string) int {
			if str == "+" {
				return a + b
			} else if str == "-" {
				return a - b
			} else if str == "*" {
				return a * b
			} else if str == "/" {
				if b == 0 {
					fmt.Println("На ноль делить нельзя")
					return 0
				}
				return a / b
			}
			return 0

		}
		fmt.Println(fn(3, 0, "/"))

	*/
	// 8
	/*
		type Order struct {
			ID     int
			Amount int
			Status string
		}

		orders := []Order{{1, 500, "new"}, {2, 1200, "paid"}, {3, 300, "paid"}, {4, 900, "failed"}}
		fn := func(orders []Order) int {
			count := 0
			for _, value := range orders {
				if value.Status == "paid" {
					count += value.Amount
				}
			}
			return count
		}
		fmt.Println(fn(orders))

	*/
	//9
	/*
		count := 0
		nextID := func() int {
			count++
			return count
		}
		fmt.Println(nextID())
		fmt.Println(nextID())

	*/
	// 10
	/*
		type User struct {
			Name string
			Age  int
		}

		usrs := []User{
			{"Ali", 17},
			{"Umed", 20},
			{"Rustam", 15},
			{"Said", 22},
		}

		user := func(user []User) []User {
			sliceUser := []User{}
			for _, u := range user {
				if u.Age >= 18 {
					sliceUser = append(sliceUser, u)
				}
			}
			return sliceUser
		}
		fmt.Println(user(usrs))

	*/
}
