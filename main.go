package main

import (
	"fmt"
	"os"

	"github.com/angelsolaorbaiceta/splitbbq/internal"
)

func main() {
	spendings := internal.ParseInput(os.Stdin)
	payments, remainderCents := internal.SettleDown(spendings)

	fmt.Println("---------- Pagos ----------")
	for _, payment := range payments {
		fmt.Println(payment)
	}
	fmt.Println("---------------------------")
	fmt.Printf("Céntimos extra: %d\n", remainderCents)
}
