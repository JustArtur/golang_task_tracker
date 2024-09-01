package notes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang_task_tracker/app/types"
	"golang_task_tracker/config"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type SpellerResponse struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func Correct(note *types.NotePayload) {
	correctText(&note.Title)
	correctText(&note.Body)
}

func correctText(text *string) {
	url := config.Envs.YNDXSpellerURL

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	err := writer.WriteField("text", *text)
	if err != nil {
		log.Printf("Error write body: %v", err)
		return
	}

	writer.Close()

	response, err := http.Post(url, "multipart/form-data", &requestBody)
	if err != nil {
		fmt.Println("Failed to request yandex:", err)
		return
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response: %v", err)
		return
	}

	var spellerResponse []SpellerResponse
	err = json.Unmarshal(respBody, &spellerResponse)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	*text = replaceIncorrectWords(&spellerResponse, []rune(*text))
}

func replaceIncorrectWords(spellerResponse *[]SpellerResponse, text []rune) string {
	for _, element := range *spellerResponse {
		offset := element.Pos
		correctWord := []rune(element.S[0])
		for i := range element.Len {
			text[offset+i] = correctWord[i]
		}
	}

	return string(text)
}
