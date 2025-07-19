package state

type SessionState struct {
	Seq       int64
	Running   bool
	Metadata  Metadata
	ReadyData ReadyData
}

func (s *SessionState) UpdateSeq(seq int64) {
	s.Seq = seq
}

func (s *SessionState) GetSessionId() string {
	return s.SessionId
}

func (s *SessionState) GetSeq() int64 {
	return s.Seq
}
