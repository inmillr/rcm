package transform

import (
	"rcm/internal/model"
	"strconv"
)

func RowToPayment(row []string) (model.Payment, error) {
	amount, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return model.Payment{}, err
	}

	return model.Payment{
		PATID:         row[0],
		ServiceDate:   row[1],
		PaymentAmount: amount,
	}, nil
}
