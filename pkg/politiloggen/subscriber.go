package politiloggen

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	BaseUrl = "https://api.politiet.no/politiloggen/v1/message"
)

type Message struct {
	Data     []Data   `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	ID                      string    `json:"id"`
	ThreadID                string    `json:"threadId"`
	Category                string    `json:"category"`
	District                string    `json:"district"`
	Municipality            string    `json:"municipality"`
	Area                    string    `json:"area"`
	IsActive                bool      `json:"isActive"`
	Text                    string    `json:"text"`
	CreatedOn               time.Time `json:"createdOn"`
	UpdatedOn               time.Time `json:"updatedOn"`
	ImageURL                any       `json:"imageUrl"`
	PreviouslyIncludedImage bool      `json:"previouslyIncludedImage"`
	IsEdited                bool      `json:"isEdited"`
}

type Metadata struct {
	RequestTime         time.Time `json:"requestTime"`
	APIVersion          string    `json:"apiVersion"`
	RequestLimitPerHour int       `json:"requestLimitPerHour"`
	TotalItems          int       `json:"totalItems"`
	PageSize            int       `json:"pageSize"`
	QueryParameters     []any     `json:"queryParameters"`
}

func sendRequest() ([]byte, error) {
	req, err := http.NewRequest("GET", BaseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	return data, err
}

func GetLastMessage() (Message, error) {
	var msg Message
	body, err := sendRequest()
	if err != nil {
		return msg, err
	}

	err = json.Unmarshal(body, &msg)
	return msg, err
}
