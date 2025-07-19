package state

type SessionState struct {
	Seq       int64
	Running   bool
	Metadata  Metadata
	ReadyData ReadyData
}

func (s *SessionState) GetSeq() int64 {
	return s.Seq
}
