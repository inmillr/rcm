package join

import "rcm/internal/model"

func BuildPaymentLookup(payments []model.Payment) map[string]float64 {

	result := make(map[string]float64)

	for _, p := range payments {
		key := p.PATID + "|" + p.ServiceDate
		result[key] += p.PaymentAmount
	}

	return result
}

func BuildCOBLookup(refs []model.COBReference) map[string]int {

	result := make(map[string]int)

	for _, r := range refs {
		result[r.FinancialClass] = r.ExpectedRank
	}

	return result
}
