package main

type Actions interface {
	Add(message string) error
	Update(message string, id int) error
	Delete(id int) error
	List()
	MarkDone(id int)
	MarkInProgress(id int)
	ListDone()
	ListInProgress()
	ListToDo()
}

type Tasks struct {
	TaskList []Task `json:"tasks"`
	Count    int    `json:"count"`
}

type Task struct {
	ID         int    `json:"id"`
	Message    string `json:"message"`
	Done       bool   `json:"done"`
	InProgress bool   `json:"in progress"`
}

const FileName string = "file.json"
