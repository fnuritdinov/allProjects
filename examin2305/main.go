package main

//1
/*
func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"

	}()
	fmt.Println(<-ch)
}
*/

// 2
/*
func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	counter := 0
	go func() {
		for _, num := range nums {
			counter += num
		}
		ch <- counter
	}()

	fmt.Println(<-ch)
}

*/

// 3
/*
func main() {
	words := []string{"go", "lang", "channel"}
	ch := make(chan string, 3)
	ch1 := make(chan string, 3)

	go func() {
		for _, word := range words {
			ch <- word
		}
	}()

	go func() {
		for wordss := range ch {
			wr := strings.ToUpper(wordss)
			ch1 <- wr
		}
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
*/
// 4
/*
func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	sm := make(chan struct{}, 3)

	for _, val := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm <- struct{}{}
			fmt.Println(val)
			time.Sleep(1 * time.Second)
			<-sm
		}()
	}
	wg.Wait()
}
*/

// 5
/*
type SafeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.data[key] = value

}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	num, ok := s.data[key]
	if !ok {
		return 0, false
	}
	return num, ok
}

func main() {
	var wg sync.WaitGroup

	sf := SafeMap{
		data: map[string]int{
			"ID": 1,
		},
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		n, _ := sf.Get("ID")
		fmt.Println(n)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		sf.Set("Number", 2)
	}()
	wg.Wait()

	for key, value := range sf.data {
		fmt.Println(key, value)
	}
}

*/
