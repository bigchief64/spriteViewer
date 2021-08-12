package tools

import (
	"embed"
	"encoding/csv"
)

func ReadData(fileName string, fs embed.FS) ([][]string, error) {
	f, err := fs.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	/*
		// skip first line
		if _, err := r.Read(); err != nil {
			return [][]string{}, err
		}*/

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
