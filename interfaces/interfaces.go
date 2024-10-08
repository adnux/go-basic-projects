package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adnux/go-basic-projects/colors"
	"github.com/adnux/go-basic-projects/interfaces/note"
	"github.com/adnux/go-basic-projects/interfaces/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func StartInterfaces() {
	printSomething("Hello")
	printSomething(1.5)

	title, content := getNoteData()
	todoText := getUserInput("Todo content:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Println(colors.Green, "Integer:", value, colors.Reset)
	case float64:
		fmt.Println(colors.Green, "Float:", value, colors.Reset)
	case string:
		fmt.Println(colors.Green, "String:", value, colors.Reset)
	}
}

func outputData[T outputtable](data T) error {
	data.Display()
	return saveData(data)
}

func getTypeName[T saver](data T) string {
	typeName := strings.Split(fmt.Sprintf("%T", data), ".")
	return typeName[len(typeName)-1]
}

func saveData[T saver](data T) error {
	err := data.Save()

	typeName := getTypeName(data)

	if err != nil {
		fmt.Printf("Saving the %s failed.\n", typeName)
		return err
	}

	fmt.Printf("Saving the %s succeeded!\n", typeName)
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
