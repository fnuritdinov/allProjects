package main

// 1
/*
type User struct {
	Name string
	Age  int
}

func printUser(u User) {
	fmt.Printf("User name: %s, Age: %d\n", u.Name, u.Age)
}
func main() {
	u := User{
		Name: "Ali",
		Age:  20,
	}
	printUser(u)
}
*/

// 2
/*type User struct {
	Name string
	Age  int
}

func increaseAge(u *User) {
	u.Age++
	fmt.Printf("name : %s, age : %d\n", u.Name, u.Age)
}
func main() {
	u := &User{
		Name: "Vali",
		Age:  10,
	}
	increaseAge(u)
}
*/

// 3
/*type User struct {
	Name string
	Age  int
}

func (u *User) reName(name string) {
	u.Name = name
}
func main() {
	u := &User{
		Name: "Bob",
		Age:  18,
	}
	u.reName("Alice")
	fmt.Printf("Name: %s, Age: %d\n", u.Name, u.Age)
}

*/

// 4
/*
type Adress struct {
	City string
}

type User struct {
	Name   string
	Age    int
	Adress Adress
}

func changeCity(u *User, city string) {
	u.Adress.City = city
}
func main() {
	u := &User{
		Name: "Farrukh",
		Age:  22,
		Adress: Adress{
			City: "London",
		},
	}
	changeCity(u, "Moscow")
	fmt.Printf("Name: %s, Age: %d, Adress: %s\n", u.Name, u.Age, u.Adress.City)
}
*/

// 5
/*
type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "Ali", Age: 30},
		{Name: "Tom", Age: 40},
		{Name: "Jack", Age: 50},
	}
	growAll(users)
	findOldest(users)
}

func growAll(users []User) {
	for _, user := range users {
		user.Age = user.Age + 1
		fmt.Println(user.Name, user.Age)
	}
}

func findOldest(users []User) User {
	oldest := users[0]
	for _, user := range users {
		if user.Age > oldest.Age {
			oldest = user
		}

	}
	fmt.Println(oldest.Age)
	return oldest

}
*/

// home work
// 1
/*
type User struct {
	Name string
	Age  int
}

func findAdults(users []User) []User {
	adults := []User{}
	for _, user := range users {
		if user.Age >= 18 {
			adults = append(adults, user)
		}
	}

	return adults
}

func main() {
	u := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 40},
		{Name: "Ali", Age: 14},
	}
	ad := findAdults(u)
	fmt.Println("Adults:", ad)
}

*/

// 2
/*type User struct {
	Name string
	Age  int
}

func increaseMinors(users []User) {
	for i := 0; i < len(users); i++ {
		if users[i].Age < 18 {
			users[i].Age++
		}
	}
	fmt.Println(users)
}

func main() {
	users := []User{
		{"Alice", 13},
		{"Bob", 20},
		{"Bt", 20},
	}
	increaseMinors(users)
}
*/

// 3
/*
type Adress struct {
	City string
}

type User struct {
	Name   string
	Age    int
	Adress Adress
}

func countFromCity(users []User, city string) int {
	count := 0
	for _, user := range users {
		if user.Adress.City == city {
			count++
		}
	}
	return count

}
func main() {
	users := []User{
		{Name: "Alice", Age: 30, Adress: Adress{City: "Moscow"}},
		{Name: "Bob", Age: 30, Adress: Adress{City: "Dc"}},
		{Name: "Charlie", Age: 30, Adress: Adress{City: "Dc"}},
		{Name: "David", Age: 30, Adress: Adress{City: "Dc"}},
	}
	count := countFromCity(users, "Dc")
	fmt.Println(count)
}
*/

// 4
/*
type User struct {
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	if u.Age > 18 {
		return true
	}
	return false
}

func main() {
	u := User{
		Name: "John Doe", Age: 25,
	}
	t := u.IsAdult()
	fmt.Println(t)
}
*/

// 5
/*
func main() {
	user := []*User{
		{Name: "Ali", Age: 22},
		{Name: "Vali", Age: 33},
	}
	reNameAll(user, "Alex")

}
func reNameAll(users []*User, name string) {
	for _, user := range users {
		user.Name = name
		fmt.Println(user.Name, user.Age)
	}
}

type User struct {
	Name string
	Age  int
}
*/

// 6
/*
func main() {
	u := []User{
		{Name: "Rustam", Age: 22},
		{Name: "Buzurg", Age: 33},
		{Name: "Farrukhruz", Age: 44},
		{Name: "Rich", Age: 55},
	}
	s := getAverageAge(u)
	fmt.Println("Average age is", s)
}

type User struct {
	Name string
	Age  float64
}

func getAverageAge(users []User) float64 {
	averageAge := 0.0
	count := 0.0
	sum := 0.0
	for _, user := range users {
		sum += user.Age
		count++
		averageAge = sum / count
	}
	return averageAge
}

*/
