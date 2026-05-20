package main

/*
	func main() {
		or := []Order{
			{ID: 1, Amount: 25.0, Status: "send"},
			{ID: 2, Amount: 20.0, Status: "paid"},
			{ID: 3, Amount: 10.0, Status: "online"},
		}
		OrdSl, err := OrderFunc(or)
		if err != nil {
			fmt.Println("Ошибка", err)
		}
		fmt.Println(OrdSl)
	}

	type Order struct {
		ID     int
		Amount float64
		Status string
	}

	func OrderFunc(ord []Order) ([]Order, error) {
		slice := make([]Order, 0)
		for _, value := range ord {
			if value.Status == "paid" {
				slice = append(slice, value)

			}
		}
		if len(slice) <= 0 {
			return ord, errors.New("Нет заказа со статусом paid в списке")
		}
		return slice, nil

}
*/
/*
func main() {
	SliceStr := []string{"golang", "goroutine", "paid", "Moscow"}
	wr, err := countWord(SliceStr, -4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wr)

}

func countWord(sl []string, n int) (int, error) {
	count := 0
	if len(sl) == 0 {
		return -1, errors.New("slice is empty")
	}
	if n < 0 {
		return -1, errors.New("n < 0")
	}
	for _, value := range sl {
		if len(value) >= n {
			count++

		}
	}
	return count, nil
}

*/
/*
func main() {
	p := []Payment{
		{ID: 1, UserID: 10, Amount: 20.0, Status: "success2"},
		{ID: 2, UserID: 11, Amount: 30.0, Status: "failed"},
		{ID: 3, UserID: 12, Amount: 45.0, Status: "new"},
		{ID: 4, UserID: 13, Amount: 20.0, Status: "blocked"},
		{ID: 5, UserID: 14, Amount: 80.0, Status: "success2"},
	}

	numb, errs := PaymentMn(p)
	fmt.Println(numb, errs)

}

type Payment struct {
	ID     int
	UserID int
	Amount float64
	Status string
}

func ValidFunc(p Payment) error {
	if p.Amount <= 0 {
		return errors.New("Неверная сумма")

	}
	if p.Status != "new" && p.Status != "success" && p.Status != "failed" {
		return errors.New("Неверный статус")
	}
	return nil

}

func PaymentMn(p []Payment) (float64, []error) {
	var total float64
	var sliceErr []error
	for _, value := range p {
		err := ValidFunc(value)
		if err != nil {
			sliceErr = append(sliceErr, errors.New("validation transaction"))
			continue
		}

		if value.Status == "success" {
			total += value.Amount
		}
	}
	errors.Join()
	return total, sliceErr

}
*/
// homework

// 2
/*
func main() {
	pr := []Product{
		{ID: 1, Name: "Alif", Price: 30},
		{ID: 2, Name: "Vali", Price: 45},
		{ID: 3, Name: "Askar", Price: 20},
	}
	price, err := PricePr(pr, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price)

}

type Product struct {
	ID    int
	Name  string
	Price int
}

func PricePr(pr []Product, minPrice int) (int, error) {
	if len(pr) == 0 {
		return 0, errors.New("Список товаров пуст")
	}
	if minPrice < 0 {
		return 0, errors.New("Минимальная цена не может отрицательная")
	}

	count := 0
	for _, value := range pr {
		if value.Price < 0 {
			return 0, errors.New("Цена товара не может быть отрицательной")
		}
		if value.Price > minPrice {
			count++
		}
	}

	return count, nil

}
*/

// 1
/*
func main() {
	us := []User{
		{ID: 1, Name: "Ali", Age: 18},
		{ID: 2, Name: "Vali", Age: 12},
		{ID: 3, Name: "Anton", Age: 33},
	}

	u, err := Person(us)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}

type User struct {
	ID   int
	Name string
	Age  int
}

func Person(u []User) ([]User, error) {
	if len(u) == 0 {
		return []User{}, errors.New("список пользователей пуст")
	}
	var slice []User
	for _, user := range u {
		if user.Age >= 18 {
			slice = append(slice, user)
		}
	}
	return slice, nil

}
*/
// 3
/*
func main() {
	st := []Student{
		{ID: 1, Name: "Melik", Score: 75},
		{ID: 2, Name: "Ali", Score: 101},
		{ID: 3, Name: "Vali", Score: 59},
		{ID: 4, Name: "Bob", Score: 88},
	}

	s, errs := StudentSt(st)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(s)

}

type Student struct {
	ID    int
	Name  string
	Score int
}

func Validate(s Student) error {
	if s.Name == "" {
		return errors.New("Поле имя не должно быть пустым")
	}
	if s.Score < 0 || 100 < s.Score {
		return errors.New("Не валидный бал")
	}
	return nil

}

func StudentSt(s []Student) ([]Student, []error) {
	var SuccessSl []Student
	var ErrorSl []error
	for _, student := range s {
		err := Validate(student)
		if err != nil {
			ErrorSl = append(ErrorSl, errors.New("Ошибка валидации"))

		} else {
			if student.Score >= 60 {
				SuccessSl = append(SuccessSl, student)
			}

		}

	}
	return SuccessSl, ErrorSl

}
*/

// 4
/*
func main() {
	orders := []Order{
		{ID: 1, ClientName: "", Amount: 20.0, Status: "new"},
		{ID: 2, ClientName: "Vali", Amount: 25.0, Status: "done"},
		{ID: 3, ClientName: "Oleg", Amount: 30.0, Status: "cancelled"},
		{ID: 4, ClientName: "Anton", Amount: 50.0, Status: "done"},
		{ID: 5, ClientName: "Banan", Amount: 50.0, Status: ""},
	}

	DoneSum, countOrders, countCancelOrders, errs := Analize(orders)
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Printf("Сумма успешных заказов на сумму: %v, количество заказов: %v, количество отменненых заказов: %v", DoneSum, countOrders, countCancelOrders)

}

type Order struct {
	ID         int
	ClientName string
	Amount     float64
	Status     string
}

func Validate(or Order) error {
	if or.ClientName == "" {
		return errors.New("Поле ClientName пустое")
	}
	if or.Amount < 0 {
		return errors.New("Поле суммы не может отрицательным")
	}
	if or.Status != "new" && or.Status != "done" && or.Status != "cancelled" {
		return errors.New("Невалидный статус - статус некорректный")
	}
	return nil
}

func Analize(ord []Order) (float64, int, int, []error) {
	var ErrSl []error
	var countOrders int
	var DoneSum float64
	var countCancelOrders int
	for _, order := range ord {
		err := Validate(order)
		if err != nil {
			ErrSl = append(ErrSl, errors.New("Ошибка валидации"))
			continue

		}
		if order.Status == "done" {
			countOrders++
			DoneSum += order.Amount

		}
		if order.Status == "cancelled" {
			countCancelOrders++
		}

	}

	return DoneSum, countOrders, countCancelOrders, ErrSl

}
*/
