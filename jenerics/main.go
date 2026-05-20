package main

// 1
/*
func main() {
	Print("Hello World")
	Print(10)
	Print(false)
	Print(3.5)
}

func Print[T any](value T) {
	fmt.Println(value)
}
*/
// 2
/*
func Min[T int | float64](a, b T) T {
	if a <= b {
		fmt.Println("first: ", a)
		return a
	}
	fmt.Println("second: ", b)
	return b

}

func main() {
	Min(10, 4)
	Min(1, 1.0)
}

*/

// 3
/*
func SumSlice[T int | float64](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	return sum
}

func main() {
	SumSlice([]int{1, 2, 3})
	SumSlice([]float64{1.0, 2.0, 3.0})
}

*/

// 4
/*
func Contains[T comparable](arr []T, value T) bool{
	for _, val := range arr{
		if val == value{
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3, 4}, 9))
}

*/

// 5
/*
func Filter[T any](arr []T, fn func(T) bool) []T {
	var newSl []T
	for _, value := range arr {
		ok := fn(value)
		if !ok {
			continue
		}
		newSl = append(newSl, value)
	}
	return newSl

}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	res := Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(res)
}
*/
// homework
// 1
/*
func Last[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}
	last := items[len(items)-1]
	fmt.Println(last)
	return last, true
}

func main() {
	sl := []int{1, 2, 3}
	n, ok := Last(sl)
	if ok {
		fmt.Println(n)
	}
}

*/

// 2
/*
func Reverse[T any](items []T) []T {
	var result []T

	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]

	}
	result = append(result, items...)
	fmt.Println(result)
	return result
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	Reverse(sl)
}

*/

// 3
/*
func Unique[T comparable](items []T) []T{
	var sl []T

	m := make(map[T]bool)
	for _, val := range items{
		if !m[val]{
			m[val] = true
			sl = append(sl, val)
		}
	}
	fmt.Print(sl)
	return sl
}

func main() {
	nSl := []int{1, 2, 3, 3, 2}
	Unique(nSl)

}

*/

// 4
/*
func CountValues[T comparable](items []T) map[T]int{
	mp := make(map[T]int)

	for _, val := range items{
		mp[val]++
	}
	fmt.Println(mp)
	return mp
}

func main() {
	sl := []string{"java", "go", "go", "java", "python"}
	CountValues(sl)
}

*/

// 5
/*
func Merge[T any](a, b []T) []T {
	sl := []T{}
	sl = append(sl, a...)
	sl = append(sl, b...)
	fmt.Println(sl)
	return sl
}

func main() {
	numSl := []int{1, 2, 3}
	numSl2 := []int{4, 5, 6}

	Merge(numSl, numSl2)
}

*/

// 6
/*
func Map[T any, R any](items []T, fn func(T) R) []R {
	newSl := []R{}
	for _, value := range items {
		r := fn(value)
		newSl = append(newSl, r)
	}
	fmt.Println(newSl)
	return newSl
}

func main() {
	nums := []int{1, 2, 3}
	Map(nums, func(n int) string {
		return fmt.Sprintf("number: %d", n)
	})
}

*/

// 7
/*
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
	fmt.Println(s.items)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	lastIndex := len(s.items) - 1

	value := s.items[lastIndex]

	s.items = s.items[:lastIndex]
	fmt.Println("lastIndex: ", value)
	fmt.Println(s.items)
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) != 0 {
		return false
	}
	return true
}

func main() {
	st := Stack[string]{}
	st.Push("go")
	st.Push("python")
	st.Pop()
	st.IsEmpty()

}

*/

// 8
/*
type Storage[T any] struct {
	data map[int]T
}

func (s *Storage[T]) Add(id int, value T){
	if id == 0{
		fmt.Println("id is empty")
		return
	}

	s.data[id] = value
	fmt.Println(s.data)
}


func (s *Storage[T]) Get(id int) (T, bool){
	if id == 0{
		fmt.Println("id is empty")
	}
	value, ok := s.data[id]
	if ok{
		return value, true
	}

	var zero T
	return zero, false
}


func (s *Storage[T]) Delete(id int){
	if id == 0{
		fmt.Println("id is empty")
	}
	_, ok := s.data[id]
	if ok{
		delete(s.data, id)
	}
	fmt.Println(s.data)
}
func (s *Storage[T]) All() []T{
	sl := []T{}
	for _, value := range s.data{
		sl = append(sl, value)
	}
	fmt.Println(sl)
	return sl
}

func main() {
	var mp Storage[string]
	mp.data = make(map[int]string)

	mp.Add(1, "nexia")
	mp.Add(2, "audi")
	mp.Add(3, "toyota")

	fmt.Println(mp.Get(2))

	mp.Delete(1)
}

*/
