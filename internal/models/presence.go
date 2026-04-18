package models

type Presence struct {
	User                 User         `json:"user"`
	Status               string       `json:"status,omitempty"`
	Since                int          `json:"since,omitempty"`
	ProcessedAtTimestamp int64        `json:"processed_at_timestamp,omitempty"`
	ClientStatus         ClientStatus `json:"client_status"`
	Activities           []*Activity  `json:"activities,omitempty"`
	AFK                  bool         `json:"afk,omitempty"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type ClientStatus struct {
	Desktop string `json:"desktop"`
}
