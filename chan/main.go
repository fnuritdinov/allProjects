package main

//1
/*
func Sum(a, b int, ch chan int) {
	c := a + b
	ch <- c
}

func main() {
	ch := make(chan int)
	go Sum(10, 20, ch)
	fmt.Println(<-ch)
	close(ch)
}
*/

// 2
/*
func main() {

	ch := make(chan int)
	counter := 0
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		counter += value
		fmt.Println(value)
	}
	fmt.Println(counter)

}
*/

// 3
/*
func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "worker 1 finished"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		chan2 <- "worker 2 finished"
	}()

	select {
	case msg := <-chan1:
		fmt.Println(msg)

	case msg := <-chan2:
		fmt.Println(msg)
	}
}

*/

//homework
//1
/*
func main() {
	ch := make(chan string, 4)

	orders := []string{"Заказ 101", "Заказ 102", "Заказ 103", "Заказ 104"}

	go func() {
		for _, order := range orders {
			ch <- order
		}
		close(ch)
	}()

	for order := range ch {
		fmt.Printf("Обрабатывается заказ %s\n", order)
	}
}

*/

// 2
/*
func main() {
	nums := make(chan int, 6)
	numbers := []int{3, 8, 15, 20, 27, 40}

	go func() {
		for _, number := range numbers {
			if number > 10 {
				nums <- number
			}
		}
		close(nums)
	}()
	for n := range nums {
		fmt.Println(n)
	}
}

*/

// 3
/*
func main() {
	products := make(chan map[int]string, 3)
	prices := make(chan map[int]int, 3)

	mpProducts := make(map[int]string)
	mpPrises := make(map[int]int)

	mpProducts[1] = "Phone"
	mpProducts[2] = "Laptop"
	mpProducts[3] = "Tablet"

	mpPrises[1] = 1000
	mpPrises[2] = 2500
	mpPrises[3] = 800

	go func() {

		products <- mpProducts
		close(products)
	}()

	go func() {
		prices <- mpPrises
		close(prices)
	}()
	for data := range products {
		for data2 := range prices {

			for key, value := range data {
				for key1, value1 := range data2 {

					if key == key1 {
						fmt.Println(value, value1)
					}
				}
			}
		}
	}
}
*/

// 4
/*
func main() {
	files := make(chan string)
	completed := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		files <- "file loaded"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		completed <- "download completed"
	}()

	select {
	case msg := <-files:
		fmt.Println(msg)

	case msg := <-completed:
		fmt.Println(msg)
	}
}
*/

// 5
/*
func main() {
	errors := make(chan string)
	success := make(chan string)

	go func() {
		errors <- "server error"
		close(errors)
	}()

	go func() {
		success <- "request success"
		close(success)
	}()

	for {
		select {
		case msg := <-errors:
			fmt.Println(msg)
			return
		case msg := <-success:
			fmt.Println(msg)
			return
		}
	}
}
*/
