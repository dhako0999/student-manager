package student

import "strconv"

type Student struct {
	Name string
	Score int
}

func (s Student) LetterGrade() string {
	if s.Score >= 90 {
		return "A"
	} else if s.Score >= 80 {
		return "B"
	} else if s.Score >= 70 {
		return "C"
	} else if s.Score >= 60 {
		return "D"
	}

	return "F"
}

func (s *Student) AddPoints(points int) {
	(*s).Score += points

	if (*s).Score > 100 {
		(*s).Score = 100
	} else if (*s).Score < 0 {
		(*s).Score = 0
	}
}

func New(name string, score int) Student {
	return Student{Name: name, Score: score}
}

func (s Student) String() string {
	return s.Name + " | Score: " + strconv.Itoa(s.Score) + " | Grade: " + s.LetterGrade()
}