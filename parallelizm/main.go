package main

// 1
/*
func main() {
	go fmt.Println(1, 2, 3)
	go fmt.Println(4, 5, 6)
	go fmt.Println(7, 8, 9)

	time.Sleep(1 * time.Second)
}

*/

// 3
/*
func Go() {
	for i := 0; i < 10; i++ {
		go fmt.Println("Goroutine number ", i)

		time.Sleep(1 * time.Second)
	}
}

*/
/*
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Goroutine number ", i)
			defer wg.Done()

		}()
		wg.Wait()
	}
	time.Sleep(5 * time.Second)

}

/*
func print(wg *sync.WaitGroup, start, end int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go print(&wg, 1, 3)
	wg.Wait()

	wg.Add(1)
	go print(&wg, 4, 6)
	wg.Wait()

	wg.Add(1)
	go print(&wg, 7, 9)
	wg.Wait()
}

*/
//
/*
func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println("tick")
	}

	time.Sleep(5 * time.Second)
}



// homework
*/

//1
/*

func Foods(wg *sync.WaitGroup, foods []string) {

	for _, food := range foods {
		wg.Add(1)
		go func(food string) {
			defer wg.Done()
			fmt.Printf("Cooking %s\n", food)
			time.Sleep(2 * time.Second)
			fmt.Printf("%s ready\n", food)
		}(food)
	}
}

func main() {

	var wg sync.WaitGroup

	foods := []string{
		"Pizza",
		"Burger",
		"Sushi",
		"Pasta",
	}

	Foods(&wg, foods)

	wg.Wait()
}

*/

//2
/*
func fast(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Fast worker")
		time.Sleep(time.Millisecond * 20)
	}
	wg.Done()
}

func slow(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Slow worker")
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go fast(&wg)
	go slow(&wg)
	wg.Wait()
}

*/

// 3
/*
func main() {
	var wg sync.WaitGroup
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Started %d\n", id)
			randomSleep := r.Int(5) + 1

			time.Sleep(time.Duration(randomSleep) * time.Second)
			fmt.Printf("Finished %d\n", id)
		}(i)
	}
	wg.Wait()
}

*/

// 4
/*
func main() {
	grades := []int{
		78, 90, 45, 88,
		67, 100, 56, 72,
	}
	var wg sync.WaitGroup
	p := len(grades) / 2
	first := grades[:p]
	second := grades[p:]

	wg.Add(2)

	go func() {
		defer wg.Done()

		countFirst := 0
		for _, value := range first {
			countFirst += value
		}
		go fmt.Println(countFirst)

	}()

	go func() {
		defer wg.Done()
		countSecond := 0
		for _, value := range second {
			countSecond += value
		}
		go fmt.Println(countSecond)
	}()
	wg.Wait()
}

*/

//5
/*
func main() {

	var wg sync.WaitGroup
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 5; i++ {

		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Runner %d started\n", id)
			randomSleep := r.Intn(5) + 1
			time.Sleep(time.Duration(randomSleep) * time.Second)

			fmt.Printf("Runner %d finished\n", id)

		}(i)
	}
	wg.Wait()
}

*/

// 6
/*
func main() {
	balance := 1000

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("CLient %d\n withdrowing money", id)

			time.Sleep(1 * time.Second)
			fmt.Printf("CLient %d\n finished", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("balance: %d", balance)
}

*/

// 7
/*
func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("request %d\n started", id)
			time.Sleep(time.Second * 1)
			fmt.Printf("request %d\n finished", id)
		}(i)
	}
	wg.Wait()
}

*/

// 8
/*
func main() {
	wg := sync.WaitGroup{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Student %d\n started exam", id)
			randomSl := r.Intn(5)
			time.Sleep(time.Duration(randomSl) * time.Second)
			fmt.Printf("Student %d\n finished exam", id)

		}(i)
	}
	wg.Wait()
}

*/
