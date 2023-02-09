package main

import "fmt"

type Task struct {
	title       string
	description string
	complete    bool
}

func (t *Task) printInfo() {
	fmt.Printf("Tarea: %s \n Descripcion: %s \n Status: %t \n", t.title, t.description, t.complete)
}

func (t *Task) changeStatus() {
	t.complete = true

}

type AllTask struct {
	tasks []*Task
}

func (tl *AllTask) addTaskList(data *Task) {
	tl.tasks = append(tl.tasks, data)

}

func (tl *AllTask) deleteTaskList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)

}

func main() {

	study := Task{"Estudiar", "Aprender sobre Golang", false}
	studyCSS := Task{"Repasar", "Leer sobre CSS en posteos", false}

	study.printInfo()
	study.changeStatus()
	study.printInfo()

	list := AllTask{}

	list.addTaskList(&study)
	list.addTaskList(&studyCSS)

	for i, task := range list.tasks {
		fmt.Println(i, task.title)
	}

}
