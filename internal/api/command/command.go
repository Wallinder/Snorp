package command

const (
	SUB_COMMAND       = 1
	SUB_COMMAND_GROUP = 2
	STRING            = 3
	INTEGER           = 4
	BOOLEAN           = 5
	USER              = 6
	CHANNEL           = 7
	ROLE              = 8
	MENTIONABLE       = 9
	NUMBER            = 10
	ATTACHMENT        = 11
)
const (
	CHAT_INPUT          = 1 //Slash commands; a text-based command that shows up when a user types /
	USER_COMMAND        = 2 //A UI-based command that shows up when you right click or tap on a user
	MESSAGE             = 3 //A UI-based command that shows up when you right click or tap on a message
	PRIMARY_ENTRY_POINT = 4 //A UI-based command that represents the primary way to invoke an app's Activity
)

const (
	GUILD_INSTALL = 0 //App is installable to servers
	USER_INSTALL  = 1 //App is installable to users
)

const (
	GUILD           = 0 //Interaction can be used within servers
	BOT_DM          = 1 //Interaction can be used within DMs with the app's bot user
	PRIVATE_CHANNEL = 2 //Interaction can be used within Group DMs and DMs other than the app's bot user
)

const (
	APP_HANDLER             = 1 //The app handles the interaction using an interaction token
	DISCORD_LAUNCH_ACTIVITY = 2 //Discord handles the interaction by launching an Activity and sending a follow-up message without coordinating with the app
)
