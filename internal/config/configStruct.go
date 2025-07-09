package config

type Setting struct {
	Env      Env
	Metadata Metadata
	Gateway  string
	Api      string
	Identify Identify
	Resume   Resume
}

// Env
type Env struct {
	Token       string
	Permissions string
}

// Metadata
type Metadata struct {
	Url     string          `json:"url"`
	Shards  int             `json:"shards"`
	Session MetadataSession `json:"session_start_limit"`
}

type MetadataSession struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

// Identify
type Identify struct {
	Op int          `json:"op"`
	D  IdentifyData `json:"d"`
}

type IdentifyData struct {
	Token      string             `json:"token"`
	Intents    int                `json:"intents"`
	Compress   bool               `json:"compress"`
	Properties IdentifyProperties `json:"properties"`
}

type IdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

// Resume
type Resume struct {
	Op int        `json:"op"`
	D  ResumeData `json:"d"`
}
type ResumeData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int64  `json:"seq"`
}
