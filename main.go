package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	var tasks Tasks
	if err := json.Unmarshal(content, &tasks); err != nil {
		fmt.Println("Error unmarshalling: ", err)
	}

	if tasks.TaskList == nil {
		tasks.TaskList = []Task{}
	}

	argsCount := len(os.Args)
	if argsCount == 1 {
		log.Fatal("Error - commands are not found")
	}
	if argsCount == 3 && os.Args[1] != "list" {
		switch os.Args[1] {
		case "add":
			if err := tasks.Add(os.Args[2]); err != nil {
				fmt.Println("Error ADD: ", err)
			}
		case "delete":
			val, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Error converting, cannot delete: ", err)
			}
			if err := tasks.Delete(val); err != nil {
				fmt.Println("Error DELETE: ", err)
			}
		case "mark-in-progress":
			val, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Error converting, cannot mark-in-progress: ", err)
			}
			tasks.MarkInProgress(val)
		case "mark-done":
			val, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Error converting, cannot mark-done: ", err)
			}
			tasks.MarkDone(val)
		}
	} else if argsCount == 3 && os.Args[1] == "list" {
		switch os.Args[2] {
		case "done":
			tasks.ListDone()
		case "todo":
			tasks.ListToDo()
		case "in-progress":
			tasks.ListInProgress()
		}
	} else if argsCount == 2 && os.Args[1] == "list" {
		tasks.List()
	} else if argsCount == 4 && os.Args[1] == "update" {
		val, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Error converting, cannot update: ", err)
		}
		if err := tasks.Update(os.Args[3], val); err != nil {
			fmt.Println("Error UPDATE: ", err)
		}
	} else {
		log.Fatal("Error, this command is not valid: ", os.Args[1:])
	}

	// newJsonContent, err := json.Marshal(tasks)
	newJsonContent, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		log.Fatal("Error marshalling type to json:", err)
	}
	err = os.WriteFile(FileName, newJsonContent, 0644)
	if err != nil {
		log.Fatal("Error writing json to file:", err)
	}
}
