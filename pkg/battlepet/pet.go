package battlepet

var SaveState map[string]*Battlepet

type Battlepet struct {
	Name  string
	Stats Stats
	Moves Moves
}

func newBattlepet(name string) Battlepet {
	return Battlepet{
		Name: name,
		Stats: Stats{
			HP:      1,
			Damage:  1,
			Defense: 1,
			Speed:   1,
		},
	}
}
