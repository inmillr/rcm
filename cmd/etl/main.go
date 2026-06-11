package main

import (
	"fmt"

	"rcm/internal/extract"
	"rcm/internal/join"
	"rcm/internal/model"
	"rcm/internal/transform"
)

func main() {

	// ----- CLAIMS ----------
	var claims []model.Claim

	claimReader, claimFile, err := extract.Stream("data/claims.csv")
	if err != nil {
		panic(err)
	}
	defer claimFile.Close()

	for i := 0; ; i++ {
		row, err := claimReader.Read()
		if err != nil {
			break
		}

		if i == 0 {
			continue
		}

		claim, err := transform.RowToClaim(row)
		if err != nil {
			fmt.Println("Claim Transform Error:", err)
			continue
		}

		claims = append(claims, claim)
	}

	// ----- PAYMENTS ----------
	var payments []model.Payment

	paymentReader, paymentFile, err := extract.Stream("data/payments.csv")
	if err != nil {
		panic(err)
	}
	defer paymentFile.Close()

	for i := 0; ; i++ {

		row, err := paymentReader.Read()
		if err != nil {
			break
		}

		if i == 0 {
			continue
		}

		payment, err := transform.RowToPayment(row)
		if err != nil {
			fmt.Println("Payment Transform Error:", err)
			continue
		}

		payments = append(payments, payment)
	}

	// ----- COB REFERENCE ----------
	var cobRefs []model.COBReference
	cobReader, cobFile, err := extract.Stream("data/cob_order.csv")
	if err != nil {
		panic(err)
	}
	defer cobFile.Close()

	for i := 0; ; i++ {

		row, err := cobReader.Read()
		if err != nil {
			break
		}

		if i == 0 {
			continue
		}

		ref, err := transform.RowToCOBReference(row)
		if err != nil {
			fmt.Println("COB Transform Error:", err)
			continue
		}

		cobRefs = append(cobRefs, ref)
	}

	// ----- VERIFY LOADS ----------

	fmt.Println("Claims:", len(claims))
	fmt.Println("Payments:", len(payments))
	fmt.Println("COB Refs:", len(cobRefs))

	// ----- BUILD LOOKUPS ----------

	paymentLookup := join.BuildPaymentLookup(payments)

	cobLookup := join.BuildCOBLookup(cobRefs)

	// ----- JOIN ----------

	rcmRows := join.BuildRCMRows(claims, paymentLookup, cobLookup)

	// ----- OUTPUT ----------

	fmt.Println()
	fmt.Println("RCM Results")
	fmt.Println("----------")

	for _, row := range rcmRows {
		fmt.Printf("%+v\n", row)
	}
}
