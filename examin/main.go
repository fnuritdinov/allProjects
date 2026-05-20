package main

import (
	"errors"
	"fmt"
)

// 2
/*
func main() {
	file, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Fields(line)
		if len(field) != 2 {
			continue
		}
		nameOfLog := field[1]

		data[nameOfLog]++
	}
	maxCount := 0
	log := ""
	for key, value := range data {
		if value > maxCount {
			maxCount = value
			log = key
		}
	}
	fmt.Println(log, maxCount)
}

*/
// 1
/*
type Order struct {
	ID        int
	Customer  string
	Products  []string
	Total     float64
	CreatedAt time.Time
}

func AddOrders(order Order, id int, customer string, products []string, total float64, createdAt time.Time)  []Order{

	slOrders := []Order{}

	if order.ID == 0{
		fmt.Println("ID is empty")
		return []Order{}
	}
	order.ID == id


	if len(order.Customer) == 0{
		fmt.Println("Customer is empty")
		return []Order{}
	}
	order.Customer = customer


	if len(order.Products) == 0{
		fmt.Println("Products is empty")
		return []Order{}
	}
	order.Products = products


	if order.Total < 1.0{
		fmt.Println("total is empty")
		return []Order{}
	}
	order.Total = total


	order.CreatedAt = createdAt

	slOrders = append(slOrders, order)

	return slOrders

}

func GetOrders(orders []Order) {
	for _, value := range orders{
		fmt.Println(value)
	}
}

func ExpensiveOrder(orders []Order)float64, int{
expensive := 0
expensiveID := 0
for _, value := range orders{
if value.Total > float64(expensive)
expensive = value.Total
value.ID = expensiveID
}
fmt.Println(expensiveID, expensive)
return expensive, expensiveID
}

func TotalSumOfOrders(orders []Order)float64{
	sum := 0
	for _, value := range orders{
		sum += int(value.Total)
	}
	fmt.Println(float64(sum))
	return float64(sum)
}

func RichCustomer(orders []Order) string{
	lenOfProducts := 0

	customerName := ""

	for _, value := range orders{
		if len(value.Products) > lenOfProducts{
			lenOfProducts = len(value.Products)
			customerName = value.Customer
		}
	}
	fmt.Println(customerName)
	return customerName
}

func main() {
	file, err := os.OpenFile("orders.txt", |O_WRONLY| O_RDONLY, O_TRUNC, 0666)
	if err != nil{
		fmt.Println("Ошибка при создании файла")
	}
	file.WriteString()

	orders := []Order{}

	AddOrders(orders, 1, "Ali", ["apple"], 12.5, "2026-05-09 19:32")

	GetOrders(orders)

	totalSum := TotalSumOfOrders()
	richOrder := RichCustomer(orders)
}

*/

type Movie struct {
	ID          int
	Title       string
	Genres      []string
	Rating      float64
	ReleaseYear int
}

func AddMovie(movies []Movie, movie Movie) ([]Movie, error) {

	if movie.ID == 0 {
		return []Movie{}, errors.New("ID is empty")
	}

	if len(movie.Title) == 0 {
		return []Movie{}, errors.New("title is empty")
	}

	if len(movie.Genres) == 0 {
		return []Movie{}, errors.New("Genres is empty")
	}

	if movie.Rating < 0.0 || movie.Rating > 10.0 {
		return []Movie{}, errors.New("invalid rating")
	}

	if movie.ReleaseYear < 1900 {
		return []Movie{}, errors.New("invalid releaseYear")
	}

	movies = append(movies, movie)
	return movies, nil

}

func GetMovies(movies []Movie) {
	for _, value := range movies {
		fmt.Println(value)
	}
}

func BestMovies(movies []Movie) (float64, int) {
	bestRating := 0.0
	idOfBest := 0
	for _, value := range movies {
		if value.Rating > bestRating {
			bestRating = value.Rating
			idOfBest = value.ID
		}
	}
	return bestRating, idOfBest
}

func TotalRating(movies []Movie) float64 {
	sumOfRating := 0.0
	for _, value := range movies {
		sumOfRating += value.Rating
	}
	return sumOfRating
}

func MostGenresMovie(movies []Movie) (string, int) {
	mostGenre := 0
	nameOfMostGenre := ""
	for _, value := range movies {
		if len(value.Genres) > mostGenre {
			mostGenre = len(value.Genres)
			nameOfMostGenre = value.Title
		}
	}
	return nameOfMostGenre, mostGenre
}

func main() {
	movies := []Movie{
		{
			ID:          1,
			Title:       "Interstellar",
			Genres:      []string{"Sci-Fi", "Drama"},
			Rating:      9.5,
			ReleaseYear: 2014,
		},
		{
			ID:          2,
			Title:       "Titanic",
			Genres:      []string{"Drama"},
			Rating:      9.0,
			ReleaseYear: 1998,
		},
	}

	//GetMovies(movies)
	//BestRating, IDoFBestRating := BestMovies(movies)
	//fmt.Println("Best Rating: ", BestRating, IDoFBestRating)

	//total := TotalRating(movies)
	//fmt.Println("Total Rating: ", total)

	//name, count := MostGenresMovie(movies)
	//fmt.Println("Most Genres Movie: ", name, count)

	movies, err := AddMovie(movies, Movie{5, "Brigada", []string{"Drama"}, 10, 1990})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(movies)

}
