package aggregate

import "rcm/internal/model"

func ByPATID(claims []model.Claim) map[string]float64 {

	totals := make(map[string]float64)

	for _, c := range claims {
		totals[c.PATID] += c.Amount
	}

	return totals
}
