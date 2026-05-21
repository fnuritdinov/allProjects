package main

// 1
/*
func main() {

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 100)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
*/
// 2
/*
func worker(id int, stop *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if atomic.LoadInt64(stop) == 1 {
			fmt.Printf("worker %d stopped \n", id)
			return
		}

		fmt.Printf("worker %d working...\n", id)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup
	var stop int64

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &stop, &wg)
	}
	go func() {
		time.Sleep(time.Second * 2)
		atomic.StoreInt64(&stop, 1)
	}()
	wg.Wait()
}
*/
//3
/*
func main() {
	var mu sync.Mutex
	var count int64

	for i := 0; i < 10; i++ {
		mu.Lock()
		atomic.AddInt64(&count, 1000)
		mu.Unlock()
	}
	fmt.Println(count)
}
*/
// 4
/*
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Withdraw(amount int) (int, error) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	if amount > b.balance {
		return 0, errors.New("error")
	}
	b.balance -= amount
	fmt.Println(b.balance)
	return b.balance, nil

}

func (b *BankAccount) Deposit(amount int) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if amount < 0 {
		return 0, errors.New("invalid amount")
	}

	b.balance += amount
	fmt.Printf("deposit balance %d\n", b.balance)
	return b.balance, nil
}

func main() {

	var wg sync.WaitGroup

	b := BankAccount{
		balance: 100,
	}
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			b.Withdraw(40)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			b.Deposit(100)
		}()
	}
	wg.Wait()
}
*/
//5
/*
type St struct {
	data string
	rwmu sync.RWMutex
}

func (s *St) reader() {
	s.rwmu.RLock()

	fmt.Println(s.data)
	s.rwmu.RUnlock()

}

func (s *St) writer(data string) {
	s.rwmu.Lock()
	s.data = data
	s.rwmu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	data := "red"
	pp := St{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pp.reader()
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		pp.writer(data)
	}()
	wg.Wait()
}

*/
// 6
/*
type Config struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *Config) Get() {
	c.mu.RLock()
	fmt.Println(c.data)
	c.mu.RUnlock()
}

func (c *Config) Set(key, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	//data := make(map[string]string)
	//data["colour"] = "green"
	conf := Config{
		data: make(map[string]string),
	}
	conf.data["key2"] = "value"
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conf.Get()
		}()

	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conf.Set("key", "val")

		}()

	}
	wg.Wait()

}

*/
// 7
/*
func main() {

	var m sync.Map
	m.Store("name", "Ali")
	m.Store("number", 1)
	m.Store("car", "bmw")
	m.Store("colour", "red")
	m.Store("sport", "tennis")

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}

*/
//8
/*

func main() {
	var wg sync.WaitGroup
	var users sync.Map

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			value, ok := users.Load("Vali")

			if ok {
				count := value.(int)
				count++
				users.Store("Vali", count)
				fmt.Println(count)
			} else {

				users.Store("Vali", 1)
			}
		}()
	}
	wg.Wait()
	value, _ := users.Load("Vali")
	fmt.Println(value)
}

*/
// 9
/*
func InitConfig() {
	fmt.Println("config loaded")
}

var once sync.Once

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(InitConfig)
		}()
	}
	wg.Wait()
}
*/
// 10
/*
func connectDB() {
	fmt.Println("successfull connection", time.Now())
}

var once sync.Once

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(connectDB)
		}()
	}
	wg.Wait()
}

*/

//homework 21 05
//5
/*
func loadSettings() {
	time.Sleep(2 * time.Second)
	fmt.Println("settings loaded")
}

var once sync.Once

func main() {

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(loadSettings)
		}()
	}
	wg.Wait()
}

*/

// 2
/*
type Warehouse struct {
	products int
	mu       sync.Mutex
}

func (w *Warehouse) Add(count int) {
	w.mu.Lock()
	if count != 0 {
		w.products += count
	}
	fmt.Println(w.products)
	w.mu.Unlock()

}

func (w *Warehouse) Take(count int) {
	w.mu.Lock()
	if count <= w.products {
		w.products -= count
	}
	fmt.Println(w.products)
	w.mu.Unlock()

}

func main() {
	var wg sync.WaitGroup
	w := Warehouse{
		products: 20,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Take(2)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Add(20)
		}()
	}
	wg.Wait()
}

*/
// 3
/*
type Stats struct {
	views int
	mu    sync.RWMutex
}

func (s *Stats) AddView() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.views++
	fmt.Println(s.views)
}

func (s *Stats) GetViews() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Println(s.views)
}

func main() {
	var wg sync.WaitGroup

	st := Stats{
		views: 10,
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			st.AddView()
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			st.GetViews()
		}()
	}
	wg.Wait()
}

*/
// 4
/*
func main() {
	var mp sync.Map
	var wg sync.WaitGroup

	usernames := []string{
		"Ali",
		"Vali",
		"Anton",
		"Andrey",
		"Alex",
	}

	for _, username := range usernames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			mp.Store(name, true)
			fmt.Println(name)
		}(username)
	}

	wg.Wait()

	if value, ok := mp.Load("Ali"); ok {
		fmt.Println(value)
	}

	mp.Delete("Anton")
	mp.Delete("Vali")

	fmt.Println("Содержимые мапы")
	mp.Range(func(key, value any) bool {
		usernames := key.(string)
		online := value.(bool)

		if online {
			fmt.Println(usernames, online)
		}
		return true
	})
}

*/

// 1
/*
func main() {

}

*/
