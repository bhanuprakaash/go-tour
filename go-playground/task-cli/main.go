package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

type Task struct {
	Id        int    `json:"id"`
	Task      string `json:"task,omitempty"`
	Completed bool   `json:"completed"`
}

var tasks = make(map[int]Task)

func add(args []string) int {
	if len(args) < 2 {
		fmt.Println("Invalid arguments!")
		return 0
	}
	id := rand.Intn(1000)
	newTask := Task{
		Id:   id,
		Task: args[1],
	}
	tasks[id] = newTask
	return id
}

func deleteTask(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Invalid arguments!")
		return false
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid id type")
		return false
	}
	if _, ok := tasks[id]; ok {
		delete(tasks, id)
		return true
	}
	return false
}

func ListAll() {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tStatus\tTask")

	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\n", task.Id, status, task.Task)
	}

	w.Flush()
}

func GetById(args []string) *Task {
	if len(args) < 2 {
		fmt.Println("Invalid arguments!")
		return nil
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid id type")
		return nil
	}
	if task, ok := tasks[id]; ok {
		return &task
	}
	return nil
}

func UpdateTask(args []string) bool {
	if len(args) < 3 {
		fmt.Println("Invalid arguments!")
		return false
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid id type")
		return false
	}
	if task, ok := tasks[id]; ok {
		task.Task = args[2]
		tasks[id] = task
		return true
	}
	return false
}

func UpdateStatus(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Invalid arguments!")
		return false
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid id type")
		return false
	}
	if task, ok := tasks[id]; ok {
		task.Completed = true
		tasks[id] = task
		return true
	}
	return false
}

func LoadTasks() error {
	var fileData []byte
	fileData, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}

	if len(fileData) == 0 {
		return nil
	}

	var tempTasks []Task
	err = json.Unmarshal(fileData, &tempTasks)
	if err != nil {
		return err
	}
	for _, task := range tempTasks {
		tasks[task.Id] = task
	}

	return nil
}

func SaveTasks() {
	var tempTasks []Task
	for _, task := range tasks {
		tempTasks = append(tempTasks, task)
	}

	fileData, err := json.MarshalIndent(&tempTasks, "", "  ")
	if err != nil {
		fmt.Println("Error while marshelling the data")
	}

	err = os.WriteFile("tasks.json", fileData, 0644)
	if err != nil {
		fmt.Println("Error while writing file")
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	var operation string
	if len(args) < 1 {
		fmt.Println("Please provide a command: add, delete, list, update, get")
		return
	}
	operation = args[0]

	err := LoadTasks()
	if err != nil {
		fmt.Println(err)
	}

	switch operation {
	case "add":
		id := add(args)
		if id != 0 {
			SaveTasks()
			fmt.Printf("Added task with ID: %d\n", id)
		}
	case "delete":
		if deleteTask(args) {
			SaveTasks()
			ListAll()
		}
	case "list":
		ListAll()
	case "update":
		updated := UpdateTask(args)
		if updated {
			SaveTasks()
			fmt.Println("Updated")
		} else {
			fmt.Println("Something went wrong")
		}
		ListAll()
	case "get":
		task := GetById(args)
		if task == nil {
			fmt.Println("something went wrong")
		}
		SaveTasks()
		fmt.Printf("[%d]: %s\n", task.Id, task.Task)
	case "done":
		done := UpdateStatus(args)
		if done {
			SaveTasks()
			ListAll()
		}
	}
}
