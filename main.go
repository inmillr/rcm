package main

import (
	"encoding/csv" // For reading CSV files
	"fmt"          // For printing
	"os"           // For file operations
	"strconv"      // For converting strings to numbers
)

// Define the Claim structure
type Claim struct {
	PATID       string
	ServiceDate string
	DeniedCode  string
	Amount      float64
}

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

	// for _, row := range records { // Each row is a []string (slice of strings), records is a [][]string (slice of rows)
	// 	fmt.Println(row)
	// }

	var claims []Claim

	for i, row := range records {

		if i == 0 {
			continue
		}

		amount, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Printf("Bad amoun on row %d\n", i)
			continue
		}

		claim := Claim{
			PATID:       row[0],
			ServiceDate: row[1],
			DeniedCode:  row[2],
			Amount:      amount,
		}

		claims = append(claims, claim)
	}

	for _, claim := range claims {
		fmt.Printf("%+v\n", claim)
	}
}
