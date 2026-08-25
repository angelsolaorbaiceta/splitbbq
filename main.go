package main

import (
	"fmt"
	"os"

	"github.com/angelsolaorbaiceta/splitbbq/internal"
)

func main() {
	spendings := internal.ParseInput(os.Stdin)
	for _, payment := range internal.SettleDown(spendings) {
		fmt.Println(payment)
	}
}
