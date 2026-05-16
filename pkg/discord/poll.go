package discord

type Poll struct {
	Question         PollMedia
	Answer           []PollAnswer
	Duration         int  `json:"duration,omitempty"`
	AllowMultiselect bool `json:"allow_multiselect,omitempty"`
	LayoutType       int  `json:"layout_type,omitempty"`
}

type PollAnswer struct {
	AnswerID  int       `json:"answer_id,omitempty"`
	PollMedia PollMedia `json:"poll_media"`
}

type PollMedia struct {
	Text  string `json:"text,omitempty"`
	Emoji *Emoji `json:"emoji,omitempty"`
}
