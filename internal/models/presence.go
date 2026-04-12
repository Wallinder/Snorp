package models

type Presence struct {
	User                 User         `json:"user"`
	Status               string       `json:"status"`
	ProcessedAtTimestamp int64        `json:"processed_at_timestamp"`
	ClientStatus         ClientStatus `json:"client_status"`
	Activities           []any        `json:"activities"`
}

type ClientStatus struct {
	Desktop string `json:"desktop"`
}
