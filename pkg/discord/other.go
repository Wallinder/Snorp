package discord

const (
	RoleMentions     = "roles"
	UserMentions     = "users"
	EveryoneMentions = "everyone"
)

type AllowedMentions struct {
	Parse []string `json:"parse,omitempty"`
	Users []string `json:"users,omitempty"`
}
