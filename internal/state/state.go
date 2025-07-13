package state

type Session struct {
	SessionId string
	Seq       int64
	Running   bool
	Metadata  Metadata
}

func (s *Session) UpdateSessionId(sessionId string) {
	s.SessionId = sessionId
}

func (s *Session) UpdateSeq(seq int64) {
	s.Seq = seq
}

func (s *Session) GetSessionId() string {
	return s.SessionId
}

func (s *Session) GetSeq() int64 {
	return s.Seq
}
