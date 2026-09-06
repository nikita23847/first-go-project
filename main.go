package main

import (
	"bufio"
	"fmt"
	Pkg "main/Pkg"
	"os"
	"strings"
	"time"
)

func help() {
	fmt.Println("all list of commands: ")
	fmt.Println("- help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("- add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("- list — эта команда позволяет получить полный список всех задач")
	fmt.Println("- del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("- done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("- events — эта команда позволяет получить список всех событий")
	fmt.Println("- exit — эта команда позволяет завершить выполнение программы")

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	app := Pkg.Alltasks{}
	inputerr := ""

	for {

		fmt.Print("Hello, you can typing: ")

		scanner.Scan()
		text := scanner.Text()

		fields := strings.Fields(text)
		if len(fields) == 0 {
			inputerr = "You wrote empty string"
			fmt.Println(inputerr)
			app.Addevent(text, inputerr)
			continue

		}
		cmd := fields[0]

		if cmd == "add" {
			if len(fields) < 2 {
				inputerr = "You didnt write task"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			if len(fields) == 2 {
				inputerr = "You didnt write descreption"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			if len(fields) > 2 {
				title := fields[1]
				created := time.Now()
				descreption := strings.Join(fields[2:], " ")
				app.Newtask(title, descreption, created)
				app.Addevent(text, inputerr)
				continue

			}
		}
		if cmd == "exit" {
			if len(fields) > 1 {
				inputerr = "you cant write more than one command"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			fmt.Println("goodbay!")
			app.Addevent(text, inputerr)
			return
		}
		if cmd == "done" {
			if len(fields) == 1 {
				inputerr = "you didnt write task"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			if len(fields) > 2 {
				inputerr = "You can complete only one task"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			app.Done(fields[1])
			app.Addevent(text, inputerr)
			continue

		}
		if cmd == "list" {
			if len(fields) > 1 {
				inputerr = "you cant write more than one command"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			if len(app.Tasks) == 0 {
				inputerr = "Your list is empty"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			app.Alltask()
			app.Addevent(text, inputerr)
			continue

		}
		if cmd == "help" {
			if len(fields) > 1 {
				inputerr = "you cant write more than one command"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			help()
			app.Addevent(text, inputerr)
			continue

		}
		if cmd == "del" {
			if len(fields) > 2 {
				inputerr = "you cant write more than one command"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			if len(fields) == 1 {
				inputerr = "you didnt write task"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			app.Del(fields[1])
			app.Addevent(text, inputerr)
			continue

		}
		if cmd == "events" {
			if len(fields) > 1 {
				inputerr = "you cant write more than one command"
				fmt.Println(inputerr)
				app.Addevent(text, inputerr)
				continue
			}
			app.Allevents()
			continue

		} else {
			fmt.Println("You wrote uncorrect value")
		}
	}
}
