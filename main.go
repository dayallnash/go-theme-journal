package main

import (
	app "dayallnash/theme-journal/src/App"
	stages "dayallnash/theme-journal/src/Stages"
	"fmt"
	"log"
	"strconv"
)

func main() {
	app.Init()

	journal := app.GetCurrentLoadedJournal()

	if true == stages.CreateRootDirIfNotExists() {
		name := stages.RunGreeting()
		journal = stages.RunSetup(name)
	} else {
		fmt.Println("Which journal do you wish to load?")

		stages.PrintListOfExistingJournals()
		choice := app.Prompt(">")

		intChoice, err := strconv.Atoi(choice)

		if err != nil {
			log.Fatal(err)
		}

		journal = stages.LoadExistingJournal(intChoice)
	}

	app.ClearScreen()

	for {
		fmt.Println("MENU\n----\n")

		fmt.Println("Want to work on today's entry? [1]")
		fmt.Println("How about working on a specific previous entry? [2]")

		choice := app.Prompt("[1|2]")

		stages.ProcessEntryChoice(choice, journal)

		app.ClearScreen()

		fmt.Println("Entry saved!\n\n")
	}
}
