package main

import (
	"log"
	"os"
	"strconv"

	"github.com/canertuzunar/miniReminder/services"
	"github.com/urfave/cli/v2"
)

var (
	todo = &services.Todos{}
)

func main() {
	todo := &services.Todos{}
	if err := todo.Load("tasks.json"); err != nil {
		log.Fatal(err)
	}

	// todo.Add("first task")
	// todo.Add("second task")
	// todo.Add("third task")
	// todo.Print()
	// err := todo.Complete(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// todo.Print()
	// todo.Save("tasks.json")
	app := &cli.App{
		Name: "todo",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(ctx *cli.Context) error {
					todo.Add(ctx.Args().First())
					todo.Save("tasks.json")
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "add a task to the list",
				Action: func(ctx *cli.Context) error {
					todo.Load("tasks.json")
					idx, _ := strconv.Atoi(ctx.Args().First())
					todo.Complete(int(idx))
					todo.Save("tasks.json")
					return nil
				},
			},
			{
				Name: "print",
				Action: func(ctx *cli.Context) error {
					todo.Load("tasks.json")
					todo.Print()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// func onReady() {
// 	systray.SetIcon(icon.Data)
// 	systray.SetTitle("Awesome App")
// 	systray.SetTooltip("Pretty awesome")
// 	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

// 	// Sets the icon of a menu item. Only available on Mac and Windows.
// 	mQuit.SetIcon(icon.Data)
// }

// func onExit() {
// 	// clean up here
// }
