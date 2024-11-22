package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	revenue, err1 := userInput("Revenue: ")
	expenses, err2 := userInput("Expenses: ")
	taxRate, err3 := userInput("Tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error")
		fmt.Println(err3)
		fmt.Println("-------")
		panic("Can't continue. Sorry.")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeToFile(ebt, profit, ratio)

	fmt.Printf("$%.1f\n", ebt)
	fmt.Printf("$%.1f\n", profit)
	fmt.Printf("$%.1f\n", ratio)
}

func userInput(text string) (input float64, err error) {
	fmt.Print(text)
	fmt.Scan(&input)
	if input < 0 {
		err = errors.New("error value should be greater than 0")
		return 0, err
	}
	return input, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeToFile(ebt, profit, ratio float64) {
	content := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	err := os.WriteFile("Financials.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
