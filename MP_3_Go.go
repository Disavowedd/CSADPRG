/*
Last names:

	Abaniel, Aaron
	Blasco, Gian
	Gutierrez, Allen
	Singson, Railey

Language: Go
Paradigm(s): Imperative Paradigm
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

func compsTax(input int) float64 {
	mInc := float64(input)
	var tax float64

	switch {
	case mInc < 4250:
		tax = 4000 * 0.045

	case mInc >= 4250 && mInc <= 4749.99:
		tax = 4500 * 0.045

	case mInc >= 4750 && mInc <= 5249.99:
		tax = 5000 * 0.045

	case mInc >= 5250 && mInc <= 5749.99:
		tax = 5500 * 0.045

	case mInc >= 5750 && mInc <= 6249.99:
		tax = 6000 * 0.045

	case mInc >= 6250 && mInc <= 6749.99:
		tax = 6500 * 0.045

	case mInc >= 6750 && mInc <= 7249.99:
		tax = 7000 * 0.045

	case mInc >= 7250 && mInc <= 7749.99:
		tax = 7500 * 0.045

	case mInc >= 7750 && mInc <= 8249.99:
		tax = 8000 * 0.045

	case mInc >= 8250 && mInc <= 8749.99:
		tax = 8500 * 0.045

	case mInc >= 8750 && mInc <= 9249.99:
		tax = 9000 * 0.045

	case mInc >= 9250 && mInc <= 9749.99:
		tax = 9500 * 0.045

	case mInc >= 9750 && mInc <= 10249.99:
		tax = 10000 * 0.045

	case mInc >= 10250 && mInc <= 10749.99:
		tax = 10500 * 0.045

	case mInc >= 10750 && mInc <= 11249.99:
		tax = 11000 * 0.045

	case mInc >= 11250 && mInc <= 11749.99:
		tax = 11500 * 0.045

	case mInc >= 11750 && mInc <= 12249.99:
		tax = 12000 * 0.045

	case mInc >= 12250 && mInc <= 12749.99:
		tax = 12500 * 0.045

	case mInc >= 12750 && mInc <= 13249.99:
		tax = 13000 * 0.045

	case mInc >= 13250 && mInc <= 13749.99:
		tax = 13500 * 0.045

	case mInc >= 13750 && mInc <= 14249.99:
		tax = 14000 * 0.045

	case mInc >= 14250 && mInc <= 14749.99:
		tax = 14500 * 0.045

	case mInc >= 14750 && mInc <= 15249.99:
		tax = 15000 * 0.045

	case mInc >= 15250 && mInc <= 15749.99:
		tax = 15500 * 0.045

	case mInc >= 15750 && mInc <= 16249.99:
		tax = 16000 * 0.045

	case mInc >= 16250 && mInc <= 16749.99:
		tax = 16500 * 0.045

	case mInc >= 16750 && mInc <= 17249.99:
		tax = 17000 * 0.045

	case mInc >= 17250 && mInc <= 17749.99:
		tax = 17500 * 0.045

	case mInc >= 17750 && mInc <= 18249.99:
		tax = 18000 * 0.045

	case mInc >= 18250 && mInc <= 18749.99:
		tax = 18500 * 0.045

	case mInc >= 18750 && mInc <= 19249.99:
		tax = 19000 * 0.045

	case mInc >= 19250 && mInc <= 19749.99:
		tax = 19500 * 0.045

	case mInc >= 19750 && mInc <= 20249.99:
		tax = 20000 * 0.045

	case mInc >= 20250 && mInc <= 20749.99:
		tax = 20500 * 0.045

	case mInc >= 20750 && mInc <= 21249.99:
		tax = 21000 * 0.045

	case mInc >= 21250 && mInc <= 21749.99:
		tax = 21500 * 0.045

	case mInc >= 21750 && mInc <= 22249.99:
		tax = 22000 * 0.045

	case mInc >= 22250 && mInc <= 22749.99:
		tax = 22500 * 0.045

	case mInc >= 22750 && mInc <= 23249.99:
		tax = 23000 * 0.045

	case mInc >= 23250 && mInc <= 23749.99:
		tax = 23500 * 0.045

	case mInc >= 23750 && mInc <= 24249.99:
		tax = 24000 * 0.045

	case mInc >= 24250 && mInc <= 24749.99:
		tax = 24500 * 0.045

	case mInc >= 24750 && mInc <= 25249.99:
		tax = 25000 * 0.045

	case mInc >= 25250 && mInc <= 25749.99:
		tax = 25500 * 0.045

	case mInc >= 25750 && mInc <= 26249.99:
		tax = 26000 * 0.045

	case mInc >= 26250 && mInc <= 26749.99:
		tax = 26500 * 0.045

	case mInc >= 26750 && mInc <= 27249.99:
		tax = 27000 * 0.045

	case mInc >= 27250 && mInc <= 27749.99:
		tax = 27500 * 0.045

	case mInc >= 27750 && mInc <= 28249.99:
		tax = 28000 * 0.045

	case mInc >= 28250 && mInc <= 28749.99:
		tax = 28500 * 0.045

	case mInc >= 28750 && mInc <= 29249.99:
		tax = 29000 * 0.045

	case mInc >= 29250 && mInc <= 29749.99:
		tax = 29500 * 0.045

	case mInc >= 29750:
		tax = 30000 * 0.045
	}

	return tax
}

func compphTax(input int) float64 {
	mInc := float64(input)
	var tax float64

	if mInc <= 10000 {
		tax = 500 / 2
	} else if mInc >= 10000.01 && mInc <= 99999.99 {
		tax = (mInc * 0.05) / 2
	} else {
		tax = 5000 / 2
	}

	return tax
}

func comppiTax(input int) float64 {
	mInc := float64(input)
	var tax float64

	if mInc <= 1500 {
		tax = mInc * 0.01
	} else if mInc > 1500 && mInc <= 5000 {
		tax = mInc * 0.02
	} else {
		tax = 5000 * 0.02
	}

	return tax
}

func compMTax(input float64) float64 {
	tInc := input
	var tax float64

	switch {
	case tInc <= 20833:
		tax = 0

	case tInc >= 20833 && tInc <= 33332:
		if tInc == 20833 {
			tax = 0
		} else {
			tax = 0 + ((tInc - 20833) * 0.15)
		}

	case tInc >= 33333 && tInc <= 66666:
		if tInc == 33333 {
			tax = 0 + ((tInc - 20833) * 0.15)
		} else {
			tax = 1875 + ((tInc - 33333) * 0.20)
		}

	case tInc >= 66667 && tInc <= 166666:
		if tInc == 66667 {
			tax = 1875 + ((tInc - 33333) * 0.20)
		} else {
			tax = 8541.80 + ((tInc - 66667) * 0.25)
		}

	case tInc >= 166667 && tInc <= 666666:
		if tInc == 166667 {
			tax = 8541.80 + ((tInc - 66667) * 0.25)
		} else {
			tax = 33541.80 + ((tInc - 166667) * 0.30)
		}

	case tInc >= 666667 && tInc <= math.MaxFloat64:
		if tInc == 666667 {
			tax = 33541.80 + ((tInc - 166667) * 0.30)
		} else {
			tax = 183541.80 + ((tInc - 666667) * 0.35)
		}
	}

	return tax
}

func main() {
	var origIncome int
	var opt string

	for {
		fmt.Printf("\nEnter an your monthly income: ")
		fmt.Scan(&origIncome)

		fmt.Println("=======Monthly Contributions=======")

		sTax := compsTax(origIncome)
		fmt.Printf("SSS tax: Php %.2f\n", sTax)

		phTax := compphTax(origIncome)
		fmt.Printf("PhilHealth tax: Php %.2f\n", phTax)

		piTax := comppiTax(origIncome)
		fmt.Printf("Pag-IBIG tax: Php %.2f\n", piTax)

		cTax := sTax + phTax + piTax
		fmt.Printf("Total Contribution tax: Php %.2f\n", cTax)

		mTax := compMTax(float64(origIncome) - cTax)
		fmt.Println("\n=======Tax Computation=======")
		fmt.Printf("Monthly income tax: Php %.2f\n", mTax)
		fmt.Printf("Net pay after tax: Php %.2f\n", float64(origIncome)-mTax)

		fmt.Println("\n=======End Result=======")
		fmt.Printf("Total Deductions: Php %.2f\n", mTax+cTax)
		fmt.Printf("Net pay after deductions: Php %.2f\n", (float64(origIncome)-mTax)-cTax)

		fmt.Printf("\nDo you want to continue? (Y/N): ")
		fmt.Scan(&opt)

		if strings.EqualFold(opt, "n") {
			break
		}
	}
	fmt.Println("=======Calculator Terminated=======")
}
