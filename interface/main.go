package main

// homework
//1
/*

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Address string
}

func (e *EmailNotifier) Send(message string) error {
	if message == "" {
		return errors.New("message is empty")
	}
	return nil

}

type SmsNotifier struct {
	Text   string
	Number string
}

func (s *SmsNotifier) Send(message string) error {
	if message == "" {
		return errors.New("message is empty")
	}

	return nil
}

type TelegramNotifier struct {
	Text  string
	Login string
}

func (t *TelegramNotifier) Send(message string) error {
	if message == "" {
		return errors.New("message is empty")
	}

	return nil
}

func Notify(n Notifier, message string) error {
	err := n.Send(message)
	if err != nil {
		return err
	}
	fmt.Println("notified: ", message)
	return nil
}

func main() {
	email := &EmailNotifier{
		Address: "123.123.123.123",
	}
	sms := &SmsNotifier{
		Number: "123456",
	}
	telegram := &TelegramNotifier{
		Login: "123.123.123.123",
	}

	Notify(email, "Hello Farrukh")
	Notify(sms, "Hello Azam")
	Notify(telegram, "Hello Aziz")

}\

*/

// 2
/*
type Validator interface {
	Validate(value string) error
}

type EmailSt struct{}

func (e EmailSt) Validate(value string) error {
	if len(value) < 5 {
		return errors.New("email too short")
	}
	if !strings.Contains(value, "@") {
		return errors.New("invalid email address")
	}
	return nil
}

type PasswordSt struct{}

func (p PasswordSt) Validate(value string) error {
	if len(value) < 6 {
		return errors.New("password too short")
	}

	hasDigit := false
	for _, r := range value {
		if unicode.IsDigit(r) {
			hasDigit = true
			break
		}
	}

	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	return nil
}

type User struct{}

func (u User) Validate(value string) error {
	if len(value) < 3 {
		return errors.New("invalid username")
	}

	return nil
}

func ValidateInput(v Validator, value string) {
	err := v.Validate(value)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Println(value)
}

func main() {
	email := &EmailSt{}
	password := &PasswordSt{}
	user := &User{}
	ValidateInput(email, "Vali@alif.tj")
	ValidateInput(password, "qwert")
	ValidateInput(user, "Antonio")

}
*/

// 3
/*
type Filter interface {
	Apply(nums []int) []int
}

type EvenFilter struct{}

func (e EvenFilter) Apply(nums []int) []int {
	newSl := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			newSl = append(newSl, num)
		}
	}
	return newSl
}

type OddFilter struct{}

func (o OddFilter) Apply(nums []int) []int {
	newSlice := []int{}

	for _, num := range nums {
		if num%2 != 0 {
			newSlice = append(newSlice, num)
		}
	}
	return newSlice
}

type PositiveFilter struct{}

func (p PositiveFilter) Apply(nums []int) []int {
	positiveSl := []int{}
	for _, num := range nums {
		if num >= 0 {
			positiveSl = append(positiveSl, num)
		}
	}
	return positiveSl
}

func ProcessFilter(f Filter, nums []int) {
	err := f.Apply(nums)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	ProcessFilter(&EvenFilter{}, []int{1, -2, 3, 4, -5})
	ProcessFilter(OddFilter{}, []int{1, -2, 3, 4, -5})
	ProcessFilter(PositiveFilter{}, []int{1, -2, 3, 4, -5})

*/
// 4
/*
type Order struct {
	ID     int
	Status string
}

type OrderProcces interface {
	Process(order *Order) error
}

var ErrInvalidStatus = errors.New("invalid status")

type NewOrder struct{}

func (n *NewOrder) Process(order *Order) error {
	if order.Status == "new" {
		order.Status = "processing"
		return nil

	}
	return ErrInvalidStatus
}

type CancelOrder struct{}

func (c *CancelOrder) Process(order *Order) error {
	if order.Status == "new" || order.Status == "processing" {
		order.Status = "canceled"
		return nil
	}
	return ErrInvalidStatus
}

type CompleteOrder struct{}

func (comp *CompleteOrder) Process(order *Order) error {
	if order.Status == "processing" {
		order.Status = "complete"
		return nil
	}
	return ErrInvalidStatus
}

func HandleOrder(p OrderProcces, order *Order) {
	err := p.Process(order)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order status is: ", order.Status)
}

func main() {
	ord := &Order{
		ID:     1,
		Status: "new",
	}
	newOr := &NewOrder{}
	complete := &CompleteOrder{}
	cancel := &CancelOrder{}
	HandleOrder(newOr, ord)
	HandleOrder(complete, ord)
	HandleOrder(cancel, ord)

}

*/

//func main() {
//var i interface{} = "Hello"
/*
	s, ok := i.(string)
	if ok {
		fmt.Println("Это строка: ", s)
	} else {
		fmt.Println("Это не строка")
	}

*/
