package event

type EventStructs struct {
	DiscordPayload struct {
		Op int    `json:"op"`
		S  int64  `json:"s"`
		T  string `json:"t"`
	}
	Hello struct {
		Op int
		D  struct {
			HeartbeatInterval float64 `json:"heartbeat_interval"`
		} `json:"d"`
	}
	HeartbeatAck struct {
		Op int `json:"op"`
	}
}
