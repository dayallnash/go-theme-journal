package stages

import (
	app "dayallnash/theme-journal/src/App"
	structs "dayallnash/theme-journal/src/Structs"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CreateRootDirIfNotExists() bool {
	themeJournalDir := "Journals"
	_, err := os.Stat(themeJournalDir)

	if os.IsNotExist(err) {
		err := os.Mkdir(themeJournalDir, 0755)
		if err != nil {
			log.Fatal(err)
		}

		return true
	}

	return false
}

func RunGreeting() string {
	println()
	print("Welcome to your new Theme Journal. Let's start by finding out who you are. What's your name?")

	name := app.Prompt(">")

	println("Great to meet you, " + name + "!")
	print("The Theme Journal system can help increase productivity, improve mindfulness, and make you more likely to keep up with healthy habits. Want to get started by finding out more about the Theme Journal system?")

	choice := app.Prompt("(Y/N)")

	if "y" == strings.ToLower(choice) {
		RunTutorial()
	}

	return name
}

func RunSetup(name string) structs.Journal {
	println()
	println("First, let's setup a theme for your journal.")
	println("A theme is kind of like an overarching 'idea' of your goals for this journal. You can have one theme per journal, but as many journals as you like.")
	println("One example of a theme might be 'Health'. In a Health themed journal, you might want to organise your workouts, log new recipes you have found, and track your muscle mass.")
	println("Take a moment to think, then let me know what your theme for your first journal will be.")

	theme := app.Prompt(">")
	journal := CreateJournalFile(theme, name)

	println("Great! Say hello to your new journal.")
	println("      ______ ______\n" + "    _/      Y      \\_\n" + "   // ~~ ~~ | ~~ ~  \\\\\n" + "  // ~ ~ ~~ | ~~~ ~~ \\\\\n" + " //________.|.________\\\\\n" + "`----------`-'----------'")

	println()
	println("Now we need to create some daily sections for your " + theme + " themed journal. These sections will appear on each day's entry for you to fill in. You can have up to 5 sections per day.")
	print("What do you want your first section to be called?")

	section1 := app.Prompt(">")

	journal = AddSection(theme, section1, journal)

	print("Do you want to add another section?")

	choice := app.Prompt("(Y/N)")

	if "n" == strings.ToLower(choice) {
		return journal
	}

	print("What do you want your second section to be called?")

	section2 := app.Prompt(">")

	journal = AddSection(theme, section2, journal)

	print("Do you want to add another section?")

	choice = app.Prompt("(Y/N)")

	if "n" == strings.ToLower(choice) {
		return journal
	}

	print("What do you want your third section to be called?")

	section3 := app.Prompt(">")

	journal = AddSection(theme, section3, journal)

	print("Do you want to add another section?")

	choice = app.Prompt("(Y/N)")

	if "n" == strings.ToLower(choice) {
		return journal
	}

	print("What do you want your fourth section to be called?")

	section4 := app.Prompt(">")

	journal = AddSection(theme, section4, journal)

	print("Do you want to add another section?")

	choice = app.Prompt("(Y/N)")

	if "n" == strings.ToLower(choice) {
		return journal
	}

	print("What do you want your fifth (final) section to be called?")

	section5 := app.Prompt(">")

	return AddSection(theme, section5, journal)
}

func CreateJournalFile(theme string, name string) structs.Journal {
	filePath := app.GetFilePathForWriting(theme)

	emptyFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	emptyFile.Chmod(0755)

	contents := structs.Journal{Owner: name, Theme: theme}
	jsonBytes, err := json.Marshal(contents)

	if err != nil {
		log.Fatal(err)
	}

	_, err = emptyFile.Write(jsonBytes)

	if err != nil {
		log.Fatal(err)
	}

	app.SetCurrentLoadedJournal(contents)

	return contents
}

func AddSection(theme string, section string, journal structs.Journal) structs.Journal {
	filePath := app.GetFilePathForWriting(theme)

	currentContents, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(currentContents, &journal)

	if journal.Section1 == "" {
		journal.Section1 = section
	} else if journal.Section2 == "" {
		journal.Section2 = section
	} else if journal.Section3 == "" {
		journal.Section3 = section
	} else if journal.Section4 == "" {
		journal.Section4 = section
	} else if journal.Section5 == "" {
		journal.Section5 = section
	}

	newJournal, err := json.Marshal(journal)

	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(filePath, newJournal, 0755)

	app.SetCurrentLoadedJournal(journal)

	return journal
}

func RunTutorial() {

}

func PrintListOfExistingJournals() {
	i := 0
	err := filepath.Walk("Journals", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "/") {
			i++
			fmt.Println(path + " [" + strconv.Itoa(i) + "]")
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func LoadExistingJournal(choice int) structs.Journal {
	i := 0

	filePath := ""

	err := filepath.Walk("Journals", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, "/") {
			i++
			if i == choice {
				filePath = path
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	currentContents, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	var journal structs.Journal
	json.Unmarshal(currentContents, &journal)

	app.SetCurrentLoadedJournal(journal)

	return journal
}
