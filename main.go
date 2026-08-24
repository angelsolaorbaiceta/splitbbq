package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Spending struct {
	spenderName string
	amountCents int
}

func (s Spending) String() string {
	return fmt.Sprintf("%s paga %0.2f€", s.spenderName, float64(s.amountCents)/100)
}

func parseInput(r io.Reader) []Spending {
	var (
		scanner   = bufio.NewScanner(r)
		spendings []Spending
		line      string
	)

	for scanner.Scan() {
		line = scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		amountIdx := len(fields) - 1
		name := strings.Join(fields[:amountIdx], " ")
		amount, err := strconv.ParseFloat(fields[amountIdx], 64)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Warning: can't parse amount from %s\n",
				fields[len(fields)-1],
			)
		}

		spendings = append(spendings, Spending{name, int(amount * 100)})
	}

	return spendings
}

func main() {
	for i, spending := range parseInput(os.Stdin) {
		fmt.Printf("%d - %s\n", i+1, spending)
	}
}
