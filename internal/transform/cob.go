package transform

import (
	"rcm/internal/model"
	"strconv"
)

func RowToCOBReference(row []string) (model.COBReference, error) {

	rank, err := strconv.Atoi(row[1])
	if err != nil {
		return model.COBReference{}, err
	}

	return model.COBReference{
		FinancialClass: row[0],
		ExpectedRank:   rank,
	}, nil
}
