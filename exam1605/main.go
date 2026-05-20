package main

/*
type Student struct {
	ID     int
	Name   string
	Age    int
	Grade  float64
	Passed bool
}

func TotalStudent(students []Student) int {
	countOfStudents := 0

	for _, value := range students {
		fmt.Println(value)
		countOfStudents++
	}
	return countOfStudents
}

func AvgOfStudents(students []Student) float64 {
	var summ float64
	var avg float64
	countOfStudetns := 0

	for _, value := range students {
		summ += value.Grade
		countOfStudetns++
		avg = summ / float64(countOfStudetns)

	}
	return avg

}

func SuccessStudents(students []Student) []Student {
	var SuccessStudentss []Student
	for _, student := range students {
		if student.Passed {
			SuccessStudentss = append(SuccessStudentss, student)
		}
	}
	return SuccessStudentss
}

func main() {

	file, err := os.ReadFile("std.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
		return
	}

	var students []Student

	err = json.Unmarshal(file, &students)
	if err != nil {
		fmt.Println("Ошибка при парсинге")
		return
	}
	fmt.Println(students)

	nums := TotalStudent(students)
	fmt.Printf("Колиество студентов %d\n", nums)

	avg := AvgOfStudents(students)
	fmt.Printf("Средний бал студентов %f\n", avg)

	success := SuccessStudents(students)
	fmt.Printf("Список успешных студентов %v\n", success)

	student := Student{
		ID:     3,
		Name:   "Anton",
		Age:    22,
		Grade:  83.0,
		Passed: true,
	}

	students = append(students, student)
	fmt.Println(students)

	data, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Ошибка при парсинге")
		return
	}
	err = os.WriteFile("std.json", data, 0644)
	if err != nil {
		fmt.Println("Ошибка при создании файла")
		return
	}
}

*/

//2
