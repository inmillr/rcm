package transform

import (
	"csvparser/internal/model"
	"strconv"
)

func RowToClaim(row []string) (model.Claim, error) {

	amount, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return model.Claim{}, err
	}

	return model.Claim{
		PATID:       row[0],
		ServiceDate: row[1],
		DeniedCode:  row[2],
		Amount:      amount,
	}, nil
}
