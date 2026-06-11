package extract

import (
	"encoding/csv"
	"os"
)

func Stream(path string) (*csv.Reader, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(file)
	return reader, file, nil
}
