package scorekeeper

type Scorer struct {
	score map[string]int
}

func NewScorer() Scorer {
	return Scorer{score: make(map[string]int)}
}

func (s *Scorer) Get(player string) int {
	return s.score[player]
}

func (s *Scorer) Add(points int, player string) {
	s.score[player] += points
}
