package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
// 1
type User struct {
	Login    string
	Password string
}

var users = make(map[string]User)

func Login(login, password string) bool {
	user, ok := users[login]
	if !ok {
		return false
	}
	if user.Password == password {
		return true
	}
	return false
}

func main() {
	users["Farrukh"] = User{Login: "Farrukh", Password: "123"}
	users["Alice"] = User{Login: "Alice", Password: "123"}
	fmt.Println(Login("Farrukh", "123"))
}

*/

// 2
/*
var vt = map[string]int{}

func main() {
	votes := []string{"Ali", "Umed", "Ali", "Azam", "Ali", "Umed"}

	for _, vote := range votes {
		vt[vote]++
	}

	maxValue := 0
	winner := ""
	for key, value := range vt {
		if value > maxValue {
			maxValue = value
			winner = key

		}
	}
	fmt.Println(winner, maxValue)
}

// 3
/*
type Employee struct {
	Name   string
	Salary int
}

func main() {
	sliceEmployee := []Employee{
		Employee{
			Name:   "Bob",
			Salary: 600,
		},
		Employee{
			Name:   "Alice",
			Salary: 800,
		},
		Employee{
			Name:   "Vali",
			Salary: 1000,
		},
		Employee{
			Name:   "DD",
			Salary: 1200,
		},
	}
	maxSal := 0
	sum := 0
	avg := 0
	betAvg := 0

	for _, employee := range sliceEmployee {
		sum += employee.Salary
		if employee.Salary > maxSal {
			maxSal = employee.Salary
		}
	}
	avg = sum / len(sliceEmployee)

	for _, employee := range sliceEmployee {
		if employee.Salary > avg {
			betAvg++
		}
	}

	fmt.Printf("Максимальная зарплата: %d, сумма всех зарплат: %d, среднее: %d, больше среднего %d", maxSal, sum, avg, betAvg)

*/

// 4
/*
type Client struct {
	Name  string
	Money int
}

func main() {

	clients := []Client{
		Client{
			Name:  "Ali",
			Money: 10,
		},
		Client{
			Name:  "Vali",
			Money: 2000,
		},
		Client{
			Name:  "Bob",
			Money: 300,
		},
		Client{
			Name:  "GG",
			Money: 4000,
		},
	}
	fn := func() ([]Client, []Client, int) {
		var richSlice []Client
		var Aslice []Client

		sumBalance := 0
		for _, cl := range clients {
			if cl.Money > 1000 {
				richSlice = append(richSlice, cl)
			}
			if strings.HasPrefix(cl.Name, "A") {
				Aslice = append(Aslice, cl)
			}
			sumBalance += cl.Money
		}
		fmt.Printf("Sum Balance: %d, Preffix A: %v, Better 1000: %v\n", sumBalance, Aslice, richSlice)
		return richSlice, Aslice, sumBalance
	}
	fn()
*/

func main() {

	reader := bufio.NewReader(os.Stdin)
	mapa := make(map[string]int)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	t := strings.Fields(text)

	for _, val := range t {
		mapa[val]++
	}
	for key, val := range mapa {
		fmt.Printf("%s %d\n", key, val)
	}

}
