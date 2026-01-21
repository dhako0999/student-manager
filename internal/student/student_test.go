package student

import "testing"


func TestLetterGrade(t *testing.T) {
	


	tests := []struct {
		name string
		score int
		expected string
	} {
		{"A grade", 95, "A"},
		{"B grade", 85, "B"},
		{"C grade", 75, "C"},
		{"D grade", 65, "D"},
		{"F grade", 40, "F"},
	}

	for _, tt := range tests {
		s := Student{Name: "Test", Score: tt.score}
		result := s.LetterGrade()

		if result != tt.expected {
			t.Errorf("%s: expected %s, got %s", tt.name, tt.expected, result)
		}
	}
}


func TestAddPoints(t *testing.T) {
	tests := []struct {
		start int
		points int
		expected int
	} {
		{80, 10, 90},
		{80, 90, 100},
		{5, -10, 0},
	}


	for _, tt := range tests {
		s := Student{Name: "Test", Score: tt.start}
		s.AddPoints(tt.points)


		if s.Score != tt.expected {
			t.Errorf("AddPoints(%d): expected %d, got %d", tt.points, tt.expected, s.Score)
		}
	}
}


