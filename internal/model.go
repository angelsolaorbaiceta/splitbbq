package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Spending struct {
	SpenderName string
	AmountCents int
}

func (s Spending) String() string {
	return fmt.Sprintf("%s paga %0.2f€", s.SpenderName, float64(s.AmountCents)/100)
}

type Payment struct {
	PayerName, PayeeName string
	AmountCents          int
}

func (p Payment) String() string {
	return fmt.Sprintf(
		"%s paga a %s %0.2f€",
		p.PayerName, p.PayeeName, float64(p.AmountCents)/100,
	)
}

func ParseInput(r io.Reader) []Spending {
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
