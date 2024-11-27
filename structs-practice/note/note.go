package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titele %v has the following content: \n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}
