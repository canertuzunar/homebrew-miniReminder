package main

import (
	"log"
	"os"

	"github.com/canertuzunar/miniRemainder/services"
)

func main() {
	todo := &services.Todos{}
	os.WriteFile("tasks.json", []byte{}, 0644)
	if err := todo.Load("tasks.json"); err != nil {
		log.Fatal(err)
	}

	todo.Add("first task")
	todo.Store("tasks.json")
}
