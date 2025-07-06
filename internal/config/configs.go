package config

import (
	"log"
	"os"
)

func EnviromentVariables() Env {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("Missing token")
	}
	permissions, exists := os.LookupEnv("BOT_PERMISSIONS")
	if !exists {
		permissions = "18977313533056"
		log.Printf("Using default permissions: %s\n", permissions)
	}
	gateway, exists := os.LookupEnv("DISC_GATEWAY")
	if !exists {
		gateway = "https://discord.com/api/gateway/bot"
		log.Printf("Using default gateway: %s\n", gateway)
	}
	api, exists := os.LookupEnv("DISC_API")
	if !exists {
		api = "https://discord.com/api/v10"
		log.Printf("Using default api: %s\n", api)
	}
	env := Env{
		Token:       token,
		Permissions: permissions,
		Gateway:     gateway,
		Api:         api,
	}
	return env
}

func IdentityVariables(token string) Identify {
	properties := Properties{
		Os:      "Linux",
		Browser: "Menial",
		Device:  "Menial",
	}
	identifyd := IdentifyD{
		Token:      token,
		Intents:    513,
		Compress:   false,
		Properties: properties,
	}
	identify := Identify{
		Op: 2,
		D:  identifyd,
	}
	return identify
}

func ResumeData(token string) Resume {
	resumed := ResumeD{
		Token: token,
	}
	resume := Resume{
		Op: 6,
		D:  resumed,
	}
	return resume
}

func Settings() Setting {
	env := EnviromentVariables()
	setting := Setting{
		Env:       env,
		Identify:  IdentityVariables(env.Token),
		Heartbeat: Heartbeat{Op: 1, D: 0},
		Resume:    ResumeData(env.Token),
	}
	return setting
}
