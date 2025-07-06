package config

type Setting struct {
	Env       Env
	Identify  Identify
	Heartbeat Heartbeat
	Resume    Resume
}

// Env
type Env struct {
	Token       string
	Permissions string
	Gateway     string
	Api         string
}

// Identify
type Identify struct {
	Op int       `json:"op"`
	D  IdentifyD `json:"d"`
}

type IdentifyD struct {
	Token      string     `json:"token"`
	Intents    int        `json:"intents"`
	Compress   bool       `json:"compress"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

// Heartbeat
type Heartbeat struct {
	Op int `json:"op"`
	D  int `json:"d"`
}

// Resume
type Resume struct {
	Op int     `json:"op"`
	D  ResumeD `json:"d"`
}
type ResumeD struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int    `json:"seq"`
}
