package main

import "sync"
import "fmt"

/*
	1. Variable for Bank Balance
	2. Print out Starting values
	3. Define Weekly revenue
	4. Loop through 52 weeks and print out how much is made; keep a running total
	5. Print out Final Balance
*/

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// Variable for Bank Balance
	var bankBalance int
	var balance sync.Mutex

	// Print out Starting values
	fmt.Printf("Initial Bank Balance: $%d.00\n", bankBalance)

	// Define Weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 1000},
		{Source: "Side Hustle", Amount: 250},
		{Source: "Investments", Amount: 150},
		{Source: "Gift", Amount: 10},
	}
	wg.Add(len(incomes))
	// Loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := range 52 {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On Week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	// Print out Final Balance
	fmt.Printf("Final Bank Balance after 52 weeks: $%d.00\n", bankBalance)
}
