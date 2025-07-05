package main

import "fmt"

func (t *Tasks) Add(message string) error {
	t.Count = len(t.TaskList)
	t.TaskList = append(t.TaskList, Task{t.Count, message, false, false})
	return nil
}

func (t *Tasks) Update(message string, id int) error {
	for i := range t.TaskList {
		if i == id {
			t.TaskList[i].Message = message
			return nil
		}
	}
	return fmt.Errorf("Error update, id is not valid: %v %v", id, message)
}

func (t *Tasks) Delete(id int) error {
	for i := range t.TaskList {
		if i == id {
			t.TaskList = append(t.TaskList[:i], t.TaskList[i+1:]...)
			for j := i; j < len(t.TaskList); j++ {
				t.TaskList[j].ID = j
			}
			t.Count--
			return nil
		}
	}
	return fmt.Errorf("Error delete, id is not valid: id: %v", id)
}

func (t *Tasks) List() {
	for _, v := range t.TaskList {
		fmt.Println(v.ID, v.Message)
	}
}

func (t *Tasks) MarkDone(id int) {
	for i := range t.TaskList {
		if i == id {
			t.TaskList[i].Done = true
		}
	}
}

func (t *Tasks) MarkInProgress(id int) {
	for i := range t.TaskList {
		if i == id {
			t.TaskList[i].InProgress = true
		}
	}
}

func (t *Tasks) ListDone() {
	for _, v := range t.TaskList {
		if v.Done {
			fmt.Println(v.ID, v.Message)
		}
	}
}

func (t *Tasks) ListInProgress() {
	for _, v := range t.TaskList {
		if v.InProgress {
			fmt.Println(v.ID, v.Message)
		}
	}
}

func (t *Tasks) ListToDo() {
	for _, v := range t.TaskList {
		if !v.Done && !v.InProgress {
			fmt.Println(v.ID, v.Message)
		}
	}
}
