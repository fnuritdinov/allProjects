package main

// 1
/*
func main() {
	us := []User{
		{Name: "Ali", Email: "vl@gmail.com"},
		{Name: "Vali", Email: "vk@gmail.com"},
	}

	fmt.Println(IsEmailExist(us, "vl@gmail.com"))

}

type User struct {
	Name  string
	Email string
}

func IsEmailExist(users []User, email string) bool {

	for _, user := range users {
		err := Validate(user)
		if err != nil {
			return false
		}
		if user.Email == email {
			return true
		}
	}
	return false
}

func Validate(user User) error {
	if user.Email == "" {
		return errors.New("Поле емайла пустое")
	}
	return nil
}
*/

//2
/*
func main() {
	var cache Cache
	cache.data = make(map[string]Item)

	cache.Set("Laptop", "Lenovo", 10)
	cache.Set("Phone", "Iphone", 10)

	laptop, ok := cache.Get("Laptop")
	if !ok {
		fmt.Println("Такого товара нет")
	} else {
		fmt.Println(laptop, ok)
	}

}

type Item struct {
	Value      string
	ExpireTime int
}

type Cache struct {
	data map[string]Item
}

func (c *Cache) Set(key, value string, ttl int) {
	c.data[key] = Item{
		Value:      value,
		ExpireTime: ttl,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	item, ok := c.data[key]
	if !ok {
		return "", false
	}
	return item.Value, true
}
*/

// 3
/*
func main() {
	prods := []Product{
		{"Phone", 100},
		{"TV", 300},
		{"Book", 20},
	}
	Products, err := FilterExpensive(prods, 299)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Products)
}

type Product struct {
	Name  string
	Price int
}

func FilterExpensive(products []Product, minPrice int) ([]Product, error) {
	var slProd []Product

	for _, product := range products {
		if product.Price > minPrice {
			slProd = append(slProd, product)
		}

	}

	if len(slProd) == 0 {
		fmt.Println("No products found")
	}

	return slProd, nil
}

*/
/*
func main() {
	a := []int{1, 2, 3, 4, 5, 1000000}
	b := []int{3, 4, 5, 6, 7, 1000000}

	mapa := make(map[int]struct{})
	sliceInt := []int{}
	for _, value := range a {
		mapa[value] = struct{}{}

	}
	for _, value := range b {
		if _, ok := mapa[value]; ok {
			sliceInt = append(sliceInt, value)
		}
	}

	fmt.Println(sliceInt)

}
*/

// 4
/*

type Order struct {
	ID     int
	UserID int
	Amount int
	Status string
}

func GetSuccessfulOrders(orders []Order) []Order {
	var results []Order
	for _, order := range orders {
		if order.Status == "success" {
			results = append(results, order)
		}
	}
	return results
}

func TotalAmountByUser(orders []Order) map[int]int {
	mp := make(map[int]int)
	for _, value := range orders {
		mp[value.UserID] = value.Amount
	}
	return mp
}

func main() {
	orders := []Order{
		{1, 10, 100, "failed"},
		{2, 20, 50, "failed"},
		{3, 10, 200, "success"},
		{4, 30, 300, "success"},
	}
	fmt.Println(GetSuccessfulOrders(orders))
	fmt.Println(TotalAmountByUser(orders))
}
*/

// homework

// 1
/*
type User struct {
	ID       int
	Username string
}

func AddUser(users []User, id int, userName string) ([]User, error) {
	userName = strings.TrimSpace(userName)
	userName = strings.ToLower(userName)

	if userName == "" {
		return users, errors.New("Поле имя не должно быть пустым")
	}
	if len(userName) < 3 {
		return users, errors.New("username is too short")
	}

	for _, user := range users {
		if user.Username == userName {
			return users, errors.New("Username already exist")
		}
	}

	newUser := User{
		ID:       id,
		Username: userName,
	}

	users = append(users, newUser)

	return users, nil
}

func main() {
	users := []User{
		{ID: 1, Username: "vadud"},
		{ID: 2, Username: "ali"},
	}

	users, err := AddUser(users, 3, " VALI ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}
*/
// 2
/*
func main() {
	var SliceInt []int

	a := []int{1, 2, 2, 3, 4, 5, 5, 6}
	b := []int{2, 4, 7}

	mp := make(map[int]bool)
	for _, valueA := range a {
		mp[valueA] = true
	}

	for _, ValueB := range b {
		delete(mp, ValueB)
	}

	for key := range mp {
		SliceInt = append(SliceInt, key)
	}
	sort.Ints(SliceInt)
	fmt.Println(SliceInt)
}
*/
// 3
/*
type Profile struct {
	Name  string
	Email string
	Age   int
}

func UpdateProfile(p *Profile, name string, email string, age int) error {
	if p == nil {
		return errors.New("profile is nil")
	}
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))

	if name == "" {
		return errors.New("name is empty")
	}

	if email == "" || !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	if age <= 1 || age >= 120 {
		errors.New("invalid age")
	}
	return nil

}

func main() {
	profile := Profile{Name: "Ali", Email: "ali@mail.com", Age: 20}

	err := UpdateProfile(&profile, " UMED", "lake@gmail.com", 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile)

}
*/
