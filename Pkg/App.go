package Pkg

import (
	"fmt"
	"time"
)

type Task struct {
	Title       string
	Basetext    string
	Done        bool
	Created     time.Time
	Createdtask time.Time
}

type Alltasks struct {
	Tasks  []Task
	Events []Event
}

func (a *Alltasks) Newtask(Title string, Basetext string, Creted time.Time) {
	task := Task{Title: Title, Basetext: Basetext, Created: Creted}
	for i := range a.Tasks {
		if a.Tasks[i].Title == Title {
			fmt.Println("This task already exists")
			return
		}
	}
	a.Tasks = append(a.Tasks, task)
	fmt.Printf("You added: %s - %s\n", task.Title, task.Basetext)
	fmt.Println("time: ", task.Created.Format("15:04 02.01.2006"))
}

func (a *Alltasks) Done(Title string) {

	for index, b := range a.Tasks {
		if b.Title == Title {
			a.Tasks[index].Done = true
			fmt.Println("You marked the task done", Title)
			a.Tasks[index].Createdtask = time.Now()
			fmt.Println("Time done: ", a.Tasks[index].Createdtask.Format("15:04 02.01.2006"))
			return
		}
	}
	fmt.Println("The task dont exist")
}
func (a *Alltasks) Alltask() {
	fmt.Println("List all tasks: ")
	for i, _ := range a.Tasks {

		fmt.Println("Task: ", a.Tasks[i].Title)
		fmt.Println("Descreption: ", a.Tasks[i].Basetext)
		fmt.Println("Created: ", a.Tasks[i].Created.Format("15:04 02.01.2006"))
		if a.Tasks[i].Done == false {
			fmt.Println("Done: no")
		} else {
			fmt.Println("Done: yes")

		}
		if a.Tasks[i].Done == true {
			fmt.Println("Time done: ", a.Tasks[i].Createdtask.Format("15:04 02.01.2006"))
		}
		fmt.Println("===========================")
	}
}

func (a *Alltasks) Del(title string) {
	for i := range a.Tasks {
		if a.Tasks[i].Title == title {
			before := a.Tasks[:i]
			after := a.Tasks[i+1:]
			del := append(before, after...)
			a.Tasks = del
			fmt.Println("deleted successful")
			return
		}
	}
	fmt.Println("We didnt find this task")
}
