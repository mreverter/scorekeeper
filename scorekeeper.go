package scorekeeper

type Scorer struct {
	score int
}

func NewScorer() Scorer {
	return Scorer{score: 0}
}

func (s *Scorer) Get() int {
	return s.score
}

func (s *Scorer) Add(points int) {
	s.score += points
}


