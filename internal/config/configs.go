package config

import (
	"log"
	"os"
)

func EnviromentVariables() Env {
	token, exists := os.LookupEnv("BOT_TOKEN")
	if !exists {
		log.Fatal("Missing token..")
	}
	permissions, exists := os.LookupEnv("BOT_PERMISSIONS")
	if !exists {
		permissions = "18977313533056"
		log.Printf("Using default permissions: %s\n", permissions)
	}
	env := Env{
		Token:       token,
		Permissions: permissions,
	}
	return env
}

func IdentityVariables(token string) Identify {
	identify := Identify{
		Op: 2,
		D: IdentifyData{
			Token:    token,
			Intents:  513,
			Compress: true,
			Properties: IdentifyProperties{
				Os:      "Linux",
				Browser: "Menial",
				Device:  "Menial",
			},
		},
	}
	return identify
}

func (s *Setting) AddData() {
	s.Metadata = GetGateway(s.Env.Token, s.Gateway)
	s.Identify = Identify{
		Op: 2,
		D: IdentifyData{
			Token:    s.Env.Token,
			Intents:  513,
			Compress: true,
			Properties: IdentifyProperties{
				Os:      "Linux",
				Browser: "Menial",
				Device:  "Menial",
			},
		},
	}
	s.Resume = Resume{
		Op: 6,
		D: ResumeData{
			Token: s.Env.Token,
			//SessionId: ,
			//Seq: ,
		},
	}
}

func Settings() Setting {
	env := EnviromentVariables()
	setting := Setting{
		Env:     env,
		Gateway: "https://discord.com/api/gateway/bot",
		Api:     "https://discord.com/api/v10",
	}
	setting.AddData()
	return setting
}
