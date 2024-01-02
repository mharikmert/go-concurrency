package main

import (
	"fmt"
	"sync"
)

type Income struct {
	source string
	amount int
}

func main() {

	totalBalance := 0

	incomes := []Income{
		{source: "Job", amount: 1000},
		{source: "Freelance", amount: 100},
		{source: "Other", amount: 10},
		{source: "Investments", amount: 50},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(len(incomes))

	for i, income := range incomes {

		// function literal
		go func(i int, income Income) {

			defer wg.Done()

			for month := 1; month <= 12; month++ {

				mutex.Lock()
				totalBalance += income.amount
				mutex.Unlock()

				fmt.Printf("Month: %d, source: %s, amount: %d\n", month, income.source, income.amount)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Total income at the end of the year: %d\n", totalBalance)

}
