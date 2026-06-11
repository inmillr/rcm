package main

import (
	"csvparser/config"
	"csvparser/internal/aggregate"
	"csvparser/internal/extract"
	"csvparser/internal/load"
	"csvparser/internal/model"
	"csvparser/internal/transform"
	"csvparser/internal/validate"
	"fmt"
)

func main() {

	cfg := config.Load()

	reader, file, err := extract.Stream(cfg.InputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var claims []model.Claim
	var rejects int

	for i := 0; ; i++ {

		row, err := reader.Read()
		if err != nil {
			break
		}

		if i == 0 {
			continue // ----- skip header row
		}

		ok, reason := validate.Row(row)
		if !ok {
			rejects++
			fmt.Println("reject:", reason, row)
			continue
		}

		claim, err := transform.RowToClaim(row)
		if err != nil {
			rejects++
			fmt.Println("transform error:", err)
			continue
		}

		claims = append(claims, claim)
	}

	fmt.Println("records:", len(claims), "rejects:", rejects)

	totals := aggregate.ByPATID(claims)

	err = load.ToCSV(cfg.OutputFile, totals)
	if err != nil {
		panic(err)
	}

	fmt.Println("ETL Complete!")
}
