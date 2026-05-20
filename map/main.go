package main

func main() {
	// 1
	/*
		slice := []int{1, 2, 3, 1, 2, 1, 3}
		mp := make(map[int]int)
		for _, v := range slice {
			mp[v]++
		}
		fmt.Println(mp)

	*/
	// 2
	/*
		m := map[string]int{
			"apple":  20,
			"banana": 40,
			"cherry": 80,
		}

		mk, ok := m["apple"]
		if !ok {
			fmt.Println("Нет такого продукта в словаре")
		}
		fmt.Println(mk, ok)

	*/

	//3
	/*
		mp := map[string]int{"Ali": 20, "Umed": 23, "Vali": 19}
		delete(mp, "Vali")
		fmt.Println(mp)

	*/

	//4
	/*
		mp := map[string]int{"Ali": 80, "Umed": 95, "Vali": 78}
		max := 0
		for key := range mp {
			if mp[key] > max {
				max = mp[key]
			}
		}
		fmt.Println(max)

	*/

	//5
	/*
		sl := []string{"go", "golang", "code"}
		mp := make(map[string]int)

		for _, value := range sl {
			mp[value] = len(value)
		}
		fmt.Println(mp)

	*/

	// home work

	// 1
	/*
		slStr := []string{"go", "java", "go", "python", "java", "rust"}
		mp := make(map[string]struct{})
		for _, value := range slStr {
			mp[value] = struct{}{}
		}
		fmt.Println(len(mp))

	*/

	// 2
	/*
		slInt := []int{4, 7, 1, 9, 7, 2, 1}
		mp := make(map[int]bool)

		for _, value := range slInt {
			if mp[value] {
				fmt.Println(value)
				return
			}
			mp[value] = true
		}

		fmt.Println("Нет повторов")

	*/
	// 3
	/*
		slStr := []string{"go", "cat", "dog", "hi", "code", "sun"}
		mp := make(map[string][]string)
		for _, value := range slStr {
			mp[strconv.Itoa(len(value))] = append(mp[strconv.Itoa(len(value))], value)
		}
		fmt.Println(mp)

	*/
	// 4
	/*
		str := "banana"
		mp := make(map[rune]int)

		for _, ch := range str {
			mp[ch]++
		}

		fmt.Println(mp)

	*/

}
