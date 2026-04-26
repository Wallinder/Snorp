package politiloggen

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

var (
	BaseUrl = "https://api.politiet.no/politiloggen/v1/message"
)

type Subscriber map[string]chan Data

func (r Subscriber) AddSubscriber(id string) chan Data {
	ch := make(chan Data)
	r[id] = ch
	return ch
}

func (r Subscriber) RemoveSubscriber(id string) {
	delete(r, id)
}

func (r Subscriber) broadcast(msg Data) {
	for _, ch := range r {
		ch <- msg
	}
}

type Log struct {
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

func getLog() (Log, error) {
	var log Log
	body, err := sendRequest()
	if err != nil {
		return log, err
	}

	err = json.Unmarshal(body, &log)
	return log, err
}

func NewSubscriber() Subscriber {
	subscriber := make(Subscriber)
	return subscriber
}

func (r Subscriber) Listen(ctx context.Context) error {
	if r == nil {
		return fmt.Errorf("no subscriber registered")
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		if len(r) == 0 {
			continue
		}

		log, err := getLog()
		if err != nil {
			slog.Error("unable to fetch log", "error", err, "url", BaseUrl)
			continue
		}

		for _, data := range log.Data {
			r.broadcast(data)
		}
		select {
		case <-ticker.C:
			continue

		case <-ctx.Done():
			return nil
		}
	}
}
