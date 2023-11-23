package services

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	if len(file) == 0 {
		return err
	}

	unmarshalError := json.Unmarshal(file, t)

	if unmarshalError != nil {
		log.Fatal(err)
	}

	return nil
}

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Store(filename string) {
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
