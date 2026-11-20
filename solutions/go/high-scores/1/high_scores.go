package highscores

type HighScores struct{
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
        scores: scores,
    }
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
    mx := s.scores[0]
	for _, elem := range s.scores {
        mx = max(mx, elem)
    }
    return mx
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
    arr := make([]int, 3)

    for _, elem := range s.scores {
        if elem >= arr[0] {
            arr[2] = arr[1]
            arr[1] = arr[0]
            arr[0] = elem
        } else if elem >= arr[1] {
            arr[2] = arr[1]
            arr[1] = elem
        } else if elem > arr[2] {
            arr[2] = elem
        }
    }

    if len(s.scores) < 3 {
        arr = arr[:len(s.scores)]
    }

    return arr
}
