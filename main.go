package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("claims.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file) // Create CSV Reader

	records, err := reader.ReadAll() // Load file into memory
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range records { // Each row is a []string (slice of strings), records is a [][]string (slice of rows)
		fmt.Println(row)
	}
}
