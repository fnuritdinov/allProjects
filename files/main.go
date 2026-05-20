package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*

// home work
// 1
/*
func main() {
	file, err := os.Open("newFile.txt")
	if err != nil{
		fmt.Println("Ошибка при открытии файла", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan(){
		line := scanner.Text()
		count ++
		fmt.Println(line, count)
	}
	if err := scanner.Err(); err != nil{
		fmt.Println("Ошибка чтения", err)
	}
}

*/

// 2
/*
func main() {
	data, err := os.ReadFile("newFile.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	l := string(data)

	text := strings.Fields(l)

	fmt.Println(len(text))
}

*/
// 3
/*
func main() {
	file, err := os.OpenFile("dFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при создании файла")
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите имена (exit для выхода):")

	for {
		fmt.Print(">")
		scanner.Scan()

		name := strings.TrimSpace(scanner.Text())

		if name == "exit" {
			break
		}

		if name == "" {
			continue
		}

		_, err := writer.WriteString(name + "\n")
		if err != nil {
			fmt.Println("Ошибка записи", err)
			return
		}
		writer.Flush()
	}
	fmt.Println("Готово! Данные сохранены в users.txt")
}


*/

// 4
/*
func main() {
	file, err := os.ReadFile("newFile.txt")
	if err != nil{
		fmt.Println("Ошибка при чтении файла")
	}

	st := string(file)

	words := strings.Fields(st)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	text := scanner.Text()

	found := false

	for _, w := range words{
		if w == text{
			found = true
		}
	}

	if found{
		fmt.Println("Найдено")
	}else{
		fmt.Println("Не найдено")
	}
}

*/

// 5
/*
func main() {
	var login, password string

	fmt.Scan(&login, &password)

	if len(login) == 0{
		fmt.Println("login is empty")
		return
	}

	if len(password) < 6{
		fmt.Println("invalid password")
		return
	}

	lBt := []byte(login + "\n")
	pBt := []byte(password)

	var sliceBt []byte

	sliceBt = append(sliceBt, lBt...)
	sliceBt = append(sliceBt, pBt...)

	file, err := os.OpenFile("fireFox.txt", )
	if err != nil {
		fmt.Println("Ошибки при создании файла")
	}
	defer file.Close()

	_, err = file.Write(sliceBt)
	if err != nil{
		fmt.Println("Ошибка при записи файла", err)
		return
	}

	fmt.Println("Запись успешно добавлена в файл")
}

*/
// 6
/*
func main() {
	file, err := os.Open("fireFox.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла")
		return
	}
	defer file.Close()

	newFile, err := os.OpenFile("copyFile.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Ошибка при создании файла")
		return
	}
	defer newFile.Close()

	cop, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("Ошибка пи копии файла")
		return
	}
	fmt.Println(cop)
}

*/
/*
func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Split(line, ",")

	tm := time.DateTime

	first, err1 := time.Parse(tm, strings.TrimSpace(parts[0]))
	second, err2 := time.Parse(tm, strings.TrimSpace(parts[1]))

	if err1 != nil || err2 != nil {
		fmt.Println("parse error")
		return
	}

	diff := first.Sub(second)
	if diff < 0 {
		diff = -diff
	}

	fmt.Println(diff)
}


*/
/*
func main() {
	reader := bufio.NewsReader(os.Stdin)

	for {
		fmt.Println("Введите команду: ")
		fmt.Println("1.Добавить задачу")
		fmt.Println("2.Показать задачу")
		fmt.Println("0.Выйти")
		fmt.Print("")

		var choice int

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Выберите задачу: ")

			reader.ReadString('\n')

			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			err := addTask(task)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Taks added")

		}
		case 2:
		err := showTask()
		if err != nil{
		fmt.Println("Error: ", err)
	}case 3:
		fmt.Println("Exit...")
		return

	default:
		fmt.Println("Invalid choice")
	}
}

func addTask(task string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибк при создании файла")
	}
	defer file.Close()

	_, err := file.WriteString(task + '\n')
	if err != nil {
		fmt.Println("Ошибка при записи")
		return err
	}
}

func showTask() {
	file, err := os.Open(fileName)
	{
		if err != nil {
			return err
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println('\n')

	index := 1

	for scanner.Scan() {
		fmt.Println("%d. %s\n", index, scanner.Text())
		index++
	}

	return scanner.Err()
}

*/
