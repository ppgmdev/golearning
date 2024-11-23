package main

//minor changes
import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var initialInvestment float64
	var years float64
	var interestRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investrent Amount: ")
	fmt.Scan(&initialInvestment)

	//fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ")
	fmt.Scan(&interestRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	gains := initialInvestment * math.Pow(1+interestRate/100, years)
	futureRealValue := gains / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", gains)
	formattedRFV := fmt.Sprintf(`Future Value Adjusted 

	for inflation: %.1f\n`, futureRealValue)
	//outputs information
	//fmt.Println("Future Value:", gains)
	//fmt.Printf("Future Value: %.1f\n", gains)
	//fmt.Println("Future Value Adjusted for inflation:", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}
