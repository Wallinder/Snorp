package socket

type GatewayCloseCodes struct {
	Code      int
	Reconnect bool
}

var CloseCodes = map[string]GatewayCloseCodes{
	"UNKNOWN_ERROR":         {4000, true},
	"UNKNOWN_OPCODE":        {4001, true},
	"DECODE_ERROR":          {4002, true},
	"NOT_AUTHENTICATED":     {4003, true},
	"AUTHENTICATION_FAILED": {4004, false},
	"ALREADY_AUTHENTICATED": {4005, true},
	"INVALID_SEQ":           {4007, true},
	"RATE_LIMITED":          {4008, true},
	"SESSION_TIMEOUT":       {4009, true},
	"INVALID_SHARD":         {4010, false},
	"SHARDING_REQUIRED":     {4011, false},
	"INVALID_API_VERION":    {4012, false},
	"INVALID_INTENTS":       {4013, false},
	"DISALLOWED_INTENTS":    {4014, false},
}
