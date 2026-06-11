package main

import (
	"encoding/csv" // For reading CSV files
	"fmt"          // For printing
	"os"           // For file operations
	"strconv"      // For converting strings to numbers
)

// ----- Define the Claim structure
type Claim struct {
	PATID       string
	ServiceDate string
	DeniedCode  string
	Amount      float64
}

// ----- Extract data from a CSV file (claims.csv)
func extract(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// ----- Validate the data (clean bad rows)
func validateRows(row []string) bool {
	if len(row) < 4 {
		return false
	}

	if row[0] == "" || row[3] == "" {
		return false
	}

	return true
}

// ----- Transform the data (convert to structs)
func transform(records [][]string) ([]Claim, error) {
	var claims []Claim

	for i, row := range records {

		// skip the header row
		if i == 0 {
			continue
		}

		if !validateRows(row) {
			continue
		}

		amount, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			continue
		}

		claims = append(claims, Claim{
			PATID:       row[0],
			ServiceDate: row[1],
			DeniedCode:  row[2],
			Amount:      amount,
		})
	}

	return claims, nil
}

// ----- Aggregate the data
func aggregate(claims []Claim) map[string]float64 {

	totals := make(map[string]float64)

	for _, c := range claims {
		totals[c.PATID] += c.Amount
	}

	return totals
}

// ----- Load the data (write CSV output)
func load(path string, totals map[string]float64) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"PATID", "TotalAmount"})

	for patid, total := range totals {
		writer.Write([]string{
			patid,
			fmt.Sprintf("%.2f", total),
		})
	}

	return nil
}

func main() {

	records, err := extract("claims.csv")
	if err != nil {
		panic(err)
	}

	claims, err := transform(records)
	if err != nil {
		panic(err)
	}

	totals := aggregate(claims)

	err = load("output.csv", totals)
	if err != nil {
		panic(err)
	}

	fmt.Println("ETL complete")
}
