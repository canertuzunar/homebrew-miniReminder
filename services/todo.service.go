package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := yellow(item.Task)
		done := red(fmt.Sprintf("%t", item.Done))
		if item.Done {
			task = green(item.Task)
			done = blue(fmt.Sprintf("%t", item.Done))

		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
