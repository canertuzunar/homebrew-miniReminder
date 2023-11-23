package main

import (
	"log"
	"os"

	"github.com/canertuzunar/miniReminder/services"
)

func main() {
	todo := &services.Todos{}
	os.WriteFile("tasks.json", []byte{}, 0644)
	if err := todo.Load("tasks.json"); err != nil {
		log.Fatal(err)
	}

	todo.Add("first task")
	todo.Add("second task")
	todo.Add("third task")
	todo.Print()
	err := todo.Complete(1)
	if err != nil {
		log.Fatal(err)
	}
	todo.Print()
	todo.Store("tasks.json")
}
