package load

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ToCSV(path string, data map[string]float64) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"PATID", "TotalAmount"})

	for k, v := range data {
		writer.Write([]string{
			k,
			fmt.Sprintf("%.2f", v),
		})
	}

	return nil
}
