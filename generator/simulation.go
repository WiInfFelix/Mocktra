package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func SimulateDay() {

	simulDay, _ := time.Parse("YYYY-MM-DD", "2022-01-01")

	amountSalesOfDay := rand.Intn(20)

	SimulateSales(amountSalesOfDay)

	fmt.Println(amountSalesOfDay)

	simulDay.Add(time.Hour * 24)
	fmt.Println(simulDay)
}

func SimulateSales(amountSalesOfDay int) (sales []ItemPurchase) {

	for i := 0; i < amountSalesOfDay; i++ {
		loc := rand.Intn(len(DefaultLocations))

		trans := simulateTransactions()

		DefaultLocations[loc].Transactions = append(DefaultLocations[loc].Transactions, trans)
	}

	return

}

func simulateTransactions() (transaction Transaction) {

	return

}
