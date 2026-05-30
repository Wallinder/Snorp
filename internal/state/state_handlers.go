package state

import "net/http"

func (s *SessionState) IsReady(w http.ResponseWriter, r *http.Request) {
	if s.Discord.ReadyData != nil {
		w.Write([]byte("ready"))
		return
	}
	http.Error(w, "instance is not ready yet", http.StatusInternalServerError)
}
