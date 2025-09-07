package event

const (
	DISPATCH                  = 0  //  Receive       An event was dispatched.
	HEARTBEAT                 = 1  //  Send/Receive  Fired periodically by the client to keep the connection alive.
	IDENTIFY                  = 2  //  Send	         Starts a new session during the initial handshake.
	PRESENCE_UPDATE           = 3  //  Send          Update the client's presence.
	VOICE_STATE_UPDATE        = 4  //  Send          Used to join/leave or move between voice channels.
	RESUME                    = 6  //  Send          Resume a previous session that was disconnected.
	RECONNECT                 = 7  //  Receive       You should attempt to reconnect and resume immediately.
	REQUEST_GUILD_MEMBERS     = 8  //  Send          Request information about offline guild members in a large guild.
	INVALID_SESSION           = 9  //  Receive       The session has been invalidated. You should reconnect and identify/resume accordingly.
	HELLO                     = 10 //  Receive       Sent immediately after connecting, contains the heartbeat_interval to use.
	HEARTBEAT_ACK             = 11 //  Receive       Sent in response to receiving a heartbeat to acknowledge that it has been received.
	REQUEST_SOUNDBOARD_SOUNDS = 31 //  Send          Request information about soundboard sounds in a set of guilds.

	UNKNOWN_ERROR         = 4000
	UNKNOWN_OPCODE        = 4001
	DECODE_ERROR          = 4002
	NOT_AUTHENTICATED     = 4003
	AUTHENTICATION_FAILED = 4004
	ALREADY_AUTHENTICATED = 4005
	INVALID_SEQ           = 4007
	RATE_LIMITED          = 4008
	SESSION_TIMEOUT       = 4009
	INVALID_SHARD         = 4010
	SHARDING_REQUIRED     = 4011
	INVALID_API_VERSION   = 4012
	INVALID_INTENTS       = 4013
	DISALLOWED_INTENTS    = 4014

	STATUS_NORMAL_CLOSURE             = 1000
	STATUS_GOING_AWAY                 = 1001
	STATUS_PROTOCOL_ERROR             = 1002
	STATUS_UNSUPPORTED_DATA           = 1003
	STATUS_NO_STATUS_RCVD             = 1005
	STATUS_ABNORMAL_CLOSURE           = 1006
	STATUS_INVALID_FRAME_PAYLOAD_DATA = 1007
	STATUS_POLICY_VIOLATION           = 1008
	STATUS_MESSAGE_TOO_BIG            = 1009
	STATUS_MANDATORY_EXTENSION        = 1010
	STATUS_INTERNAL_ERROR             = 1011
	STATUS_SERVICE_RESTART            = 1012
	STATUS_TRY_AGAIN_LATER            = 1013
	STATUS_BAD_GATEWAY                = 1014
	STATUS_TLS_HANDSHAKE              = 1015
)

var SocketErrors = map[int]bool{
	UNKNOWN_ERROR:         true,
	UNKNOWN_OPCODE:        true,
	DECODE_ERROR:          true,
	NOT_AUTHENTICATED:     true,
	AUTHENTICATION_FAILED: false,
	ALREADY_AUTHENTICATED: true,
	INVALID_SEQ:           true,
	RATE_LIMITED:          true,
	SESSION_TIMEOUT:       true,
	INVALID_SHARD:         false,
	SHARDING_REQUIRED:     false,
	INVALID_API_VERSION:   false,
	INVALID_INTENTS:       false,
	DISALLOWED_INTENTS:    false,

	STATUS_NORMAL_CLOSURE:             false,
	STATUS_GOING_AWAY:                 true,
	STATUS_PROTOCOL_ERROR:             true,
	STATUS_UNSUPPORTED_DATA:           false,
	STATUS_NO_STATUS_RCVD:             false,
	STATUS_ABNORMAL_CLOSURE:           false,
	STATUS_INVALID_FRAME_PAYLOAD_DATA: true,
	STATUS_POLICY_VIOLATION:           true,
	STATUS_MESSAGE_TOO_BIG:            true,
	STATUS_MANDATORY_EXTENSION:        false,
	STATUS_INTERNAL_ERROR:             true,
	STATUS_SERVICE_RESTART:            false,
	STATUS_TRY_AGAIN_LATER:            false,
	STATUS_BAD_GATEWAY:                false,
	STATUS_TLS_HANDSHAKE:              false,
}
