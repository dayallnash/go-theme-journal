package stages

import (
	app "dayallnash/theme-journal/src/App"
	structs "dayallnash/theme-journal/src/Structs"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func LoadTodaysPage(journal structs.Journal) (structs.Page, structs.Journal) {
	now := time.Now()

	for _, page := range journal.Pages {
		if page.Date == now.Format("02-01-2006") {
			return EditTodaysPage(journal), journal
		}
	}

	fmt.Println("No entry for today found. Let's start a new one!")

	return EditTodaysPage(journal), journal
}

func FindSpecificPage(journal structs.Journal, date string) (int, structs.Page) {
	i := 0
	for _, pageObj := range journal.Pages {
		i++
		if pageObj.Date == date {
			return i, pageObj
		}
	}

	var page structs.Page

	return i, page
}

func FindTodaysPage(journal structs.Journal) (int, structs.Page, bool) {
	now := time.Now()

	i := 0
	for _, pageObj := range journal.Pages {
		i++
		if pageObj.Date == now.Format("02-01-2006") {
			return i, pageObj, false
		}
	}

	var page structs.Page

	return i, page, true
}

func EditSpecificPage(journal structs.Journal, date string) structs.Page {
	i, page := FindSpecificPage(journal, date)

	if "" == page.Location {
		fmt.Print("Every entry needs a location, so you always know where you were when you were writing it. Where were you when you on this day?")

		location := app.Prompt(">")
		page.Location = location
	}

	theme := journal.Theme

	fmt.Print("Now let's start filling in your daily sections. If you don't want to fill in a section, just skip it. You can come back to it later.")

	page = WriteSections(journal, page, theme)

	journal.Pages[i-1] = page

	filePath := app.GetFilePathForWriting(theme)

	newJournal, err := json.Marshal(journal)

	if err != nil {
		log.Fatal(err)
	}

	if err = os.Truncate(filePath, 0); err != nil {
		log.Fatal("Failed to truncate " + filePath)
	}

	os.WriteFile(filePath, newJournal, 0755)

	return page
}

func EditTodaysPage(journal structs.Journal) structs.Page {
	newPage := true

	i, page, newPage := FindTodaysPage(journal)

	now := time.Now()

	page.Date = now.Format("02-01-2006")

	if "" == page.Location {
		fmt.Print("Every entry needs a location, so you always know where you were when you were writing it. Where are you today?")

		location := app.Prompt(">")
		page.Location = location
	}

	theme := journal.Theme

	fmt.Print("Now let's start filling in your daily sections. If you don't want to fill in a section, just skip it. You can come back to it later.")

	page = WriteSections(journal, page, theme)

	if newPage {
		journal.Pages = append(journal.Pages, page)
	} else {
		journal.Pages[i-1] = page
	}

	filePath := app.GetFilePathForWriting(theme)

	newJournal, err := json.Marshal(journal)

	if err != nil {
		log.Fatal(err)
	}

	if err = os.Truncate(filePath, 0); err != nil {
		log.Fatal("Failed to truncate " + filePath)
	}

	os.WriteFile(filePath, newJournal, 0755)

	return page
}

func WriteSections(journal structs.Journal, page structs.Page, theme string) structs.Page {
	if "" != journal.Section1 {
		page.Section1 = WriteSection(journal.Section1, page.Section1, theme)
	}

	if "" != journal.Section2 {
		page.Section2 = WriteSection(journal.Section2, page.Section2, theme)
	}

	if "" != journal.Section3 {
		page.Section3 = WriteSection(journal.Section3, page.Section3, theme)
	}

	if "" != journal.Section4 {
		page.Section4 = WriteSection(journal.Section4, page.Section4, theme)
	}

	if "" != journal.Section5 {
		page.Section5 = WriteSection(journal.Section5, page.Section5, theme)
	}

	return page
}

func WriteSection(journalSection string, pageSection string, theme string) string {
	swapFileName := "Journals/" + theme + "_Entry_Section5.swp"

	os.WriteFile(swapFileName, []byte(journalSection+":\n\n"+pageSection), 0755)
	cmd := exec.Command("nano", swapFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

	byte, _ := os.ReadFile("Journals/" + theme + "_Entry_Section5.swp")

	string := string(byte)
	string = strings.Replace(string, journalSection+":\n\n", "", 1)

	os.Remove(swapFileName)

	return string
}

func PromptToLoadSpecificPage(journal structs.Journal) (structs.Page, structs.Journal) {
	fmt.Print("Ok! Tell me a date that you'd like to load. Please do it in dd-mm-yyyy format.")
	date := strings.Replace(app.Prompt(">"), "\n", "", -1)

	fmt.Print(date)
	for _, page := range journal.Pages {
		fmt.Println(page.Date)
		if page.Date == date {
			return EditSpecificPage(journal, date), journal
		}
	}

	fmt.Println("Could not find an entry for that date :( Let's try again")

	return PromptToLoadSpecificPage(journal)
}

func ProcessEntryChoice(choice string, journal structs.Journal) (structs.Page, structs.Journal) {
	if "1" == strings.ToLower(choice) {
		return LoadTodaysPage(journal)
	} else {
		return PromptToLoadSpecificPage(journal)
	}
}
