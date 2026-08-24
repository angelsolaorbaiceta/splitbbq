package main

import (
	"fmt"
	"os"

	"github.com/angelsolaorbaiceta/splitbbq/internal"
)

func main() {
	for i, spending := range internal.ParseInput(os.Stdin) {
		fmt.Printf("%d - %s\n", i+1, spending)
	}
}
