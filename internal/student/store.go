package student

import (
	"encoding/json"
	"os"
)


type Store struct {
	Path string
}

func NewStore(path string) Store {
	return Store{Path: path}
}

func (st Store) Load() ([]Student, error) {
	data, err := os.ReadFile(st.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Student{}, nil
		}

		return nil, err
	}

	var students []Student
	if err := json.Unmarshal(data, &students); err != nil {
		return nil, err
	}

	return students, nil

}

func (st Store) Save (students []Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(st.Path, data, 0644)
}