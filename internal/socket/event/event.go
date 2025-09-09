package event

import (
	"context"
	"log"
	"menial/internal/state"

	"github.com/coder/websocket"
)

const APIversion = "10"

func EventListener(ctx context.Context, cancel context.CancelFunc, session *state.SessionState) {
	if session.Conn != nil {
		log.Println("Connection already open")
		return
	}

	url := session.Metadata.Url

	if session.Resume {
		url = session.ReadyData.ResumeGatewayURL
	}

	url += "/?v=" + APIversion + "&encoding=json"

	log.Printf("Connecting to socket: %s\n", url)

	conn, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		log.Printf("Error opening connection: %v\n", err)
		return
	}
	session.Conn = conn

	defer func() {
		session.Conn.Close(1006, "Normal Closure")
		session.Conn = nil
		cancel()
	}()

	EventHandler(ctx, session)
}
