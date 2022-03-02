package app

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var clear map[string]func()

func Init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Prompt(name string) string {
	fmt.Print(" ", name, " ")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text == "" {
		fmt.Println("Did not enter a valid string!")

		text = Prompt(name)
	}

	return text
}

func GetFilePathForWriting(theme string) string {
	return "Journals/" + theme + ".json"
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		fmt.Println("Your platform is unsupported! I can't clear the terminal screen :( The output might look a bit funny at times.")
	}
}
