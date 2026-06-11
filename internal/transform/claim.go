package transform

import (
	"fmt"
	"rcm/internal/model"
	"strconv"
)

func RowToClaim(row []string) (model.Claim, error) {

	if len(row) < 4 {
		return model.Claim{}, fmt.Errorf("invalid row: %+v", row)
	}

	amount, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return model.Claim{}, fmt.Errorf("invalid amount %q: %w", row[3], err)
	}

	return model.Claim{
		PATID:       row[0],
		ServiceDate: row[1],
		DeniedCode:  row[2],
		Charge:      amount,
	}, nil
}
