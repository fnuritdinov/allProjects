package main

import (
	"errors"
	"fmt"
)

type Storage struct {
	data map[string]interface{}
}

func (s *Storage) Set(key string, value interface{}) {
	if len(key) == 0 {
		fmt.Println(errors.New("Поле ключа пустое"))
	}
	if value == nil {
		fmt.Println(errors.New("Значение пустое"))
	}

	s.data[key] = value
}

func (s *Storage) Get(key string) (interface{}, bool) {
	if len(key) == 0 {
		return nil, false
	}
	return s.data[key], true
}

type User struct {
	Name string
	Age  int
}

func main() {
	mp := Storage{
		data: make(map[string]interface{}),
	}
	mp.Set("name", "Ali")
	mp.Set("numbers", 9929)
	mp.Set("user", User{Name: "Vali", Age: 33})
	fmt.Println(mp)

}
