package join

import "rcm/internal/model"

func BuildRCMRows(
	claims []model.Claim,
	paymentLookup map[string]float64,
	cobLookup map[string]int,
) []model.RCMRow {
	var rows []model.RCMRow
	for _, claim := range claims {

		key := claim.PATID + "|" + claim.ServiceDate

		paid := paymentLookup[key]
		rank := cobLookup[claim.FinancialClass]

		rows = append(rows, model.RCMRow{
			PATID:        claim.PATID,
			ServiceDate:  claim.ServiceDate,
			Charge:       claim.Charge,
			Paid:         paid,
			Balance:      claim.Charge - paid,
			ExpectedRank: rank,
			IsDenied:     claim.DeniedCode != "",
		})
	}

	return rows
}
