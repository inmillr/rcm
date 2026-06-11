package model

type Claim struct {
	PATID          string
	ServiceDate    string
	Charge         float64
	FinancialClass string
	DeniedCode     string
	Amount         float64
}

type Payment struct {
	PATID         string
	ServiceDate   string
	PaymentAmount float64
}

type COBReference struct {
	FinancialClass string
	ExpectedRank   int
}

type RCMRow struct {
	PATID        string
	ServiceDate  string
	Charge       float64
	Paid         float64
	Balance      float64
	ExpectedRank int
	IsDenied     bool
}
