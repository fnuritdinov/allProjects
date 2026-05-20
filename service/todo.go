package service

import (
	"errors"
	"fmt"
)

type TodoService struct {
	Tasks  []Task
	NextID int
}

func (t *TodoService) AddTask(title, description string) {
	if title == "" {
		errors.New("Please enter a title")
	}
	if description == "" {
		errors.New("Please enter a description")
	}

	newTask := Task{
		ID:          t.NextID,
		Title:       title,
		Description: description,
	}
	t.Tasks = append(t.Tasks, newTask)
	t.NextID++
}

func (t *TodoService) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	var status string
	for _, task := range t.Tasks {
		status = "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d.%s: %s [%s]\n]", task.ID, task.Title, task.Description, status)
	}
}

var ErrTaskNotFound = errors.New("Task not found")

func (t *TodoService) DoneTask(id int) error {
	for idx, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[idx].Done = true
			return nil
		}
	}
	return ErrTaskNotFound
}

func (t *TodoService) Delete(id int) error {
	for idx, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:idx], t.Tasks[idx+1:]...)
		}
	}
	return nil
}
