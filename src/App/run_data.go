package app

import structs "dayallnash/theme-journal/src/Structs"

type runData struct {
	currentLoadedJournal structs.Journal
	currentLoadedPage    structs.Page
}

func GetCurrentLoadedJournal() structs.Journal {
	var runData runData
	return runData.currentLoadedJournal
}

func SetCurrentLoadedJournal(journal structs.Journal) {
	var runData runData
	runData.currentLoadedJournal = journal
}

func GetCurrentLoadedPage() structs.Page {
	var runData runData
	return runData.currentLoadedPage
}

func SetCurrentLoadedPage(page structs.Page) {
	var runData runData
	runData.currentLoadedPage = page
}
