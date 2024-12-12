package main

import "fmt"

func addExpense(expenses *[]float64, amount float64) {
	*expenses = append(*expenses, amount)
	fmt.Printf("Add expense: %.2f\n", amount)
}
func calculateTotal(expenses []float64) float64 {
	total := 0.0
	for _, expense := range expenses {
		total += expense
	}
	return total
}
func updateExpense(expenses *[]float64, index int, newAmount float64) {
	if index >= 0 && index < len(*expenses) {
		(*expenses)[index] = newAmount
		fmt.Printf("Update expense at index %d to %.2f\n", index, newAmount)
	} else {
		fmt.Println("Invalid index. No update made.")
	}
}

func main() {
	expenses := []float64{}

	addExpense(&expenses, 50.75)
	addExpense(&expenses, 25.40)
	addExpense(&expenses, 10.00)

	//display current expense
	fmt.Println("Current expense:", expenses)

	//Calculate total expenses
	total := calculateTotal(expenses)
	fmt.Printf("Total expenses: %.2f\n", total)

	//Update an expense
	updateExpense(&expenses, 1, 30.00)
	fmt.Println("Update expenses:", expenses)

	//Recalculate total
	total = calculateTotal(expenses)
	fmt.Printf("Updated total expenses: %.2f\n", total)
}
