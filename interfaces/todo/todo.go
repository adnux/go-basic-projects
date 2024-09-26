package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo has the following content:\n%v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text:      content,
		CreatedAt: time.Now(),
	}, nil
}
