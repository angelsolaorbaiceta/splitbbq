package internal

import "sort"

type balance struct {
	name        string
	amountCents int
}

func SettleDown(spendings []Spending) ([]Payment, int) {
	shareCents, remainderCents := calculateShareAndRemainder(spendings)
	debtors, creditors := splitDebtorsFromCreditors(spendings, shareCents)

	return calculatePayments(debtors, creditors, remainderCents), remainderCents
}

func splitDebtorsFromCreditors(spendings []Spending, shareCents int) ([]balance, []balance) {
	var (
		debtors, creditors []balance
	)

	for _, spending := range spendings {
		net := spending.AmountCents - shareCents
		if net < 0 {
			debtors = append(debtors, balance{spending.SpenderName, -net})
		} else if net > 0 {
			creditors = append(creditors, balance{spending.SpenderName, net})
		}
	}

	// Sort descending to greedily match largest magnitudes first.
	sort.Slice(
		debtors,
		func(i, j int) bool { return debtors[i].amountCents > debtors[j].amountCents },
	)
	sort.Slice(
		creditors,
		func(i, j int) bool { return creditors[i].amountCents > creditors[j].amountCents },
	)

	return debtors, creditors
}

func calculateShareAndRemainder(spendings []Spending) (int, int) {
	total := 0
	for _, spending := range spendings {
		total += spending.AmountCents
	}

	if total == 0 {
		return 0, 0
	}

	share := total / len(spendings)
	remainder := total % len(spendings)

	return share, remainder
}

func calculatePayments(debtors, creditors []balance, remainderCents int) []Payment {
	var (
		payments = make([]Payment, 0)
		i, j     = 0, 0
	)

	for i < len(debtors) && j < len(creditors) {
		d := &debtors[i]
		c := &creditors[j]

		transfer := min(d.amountCents, c.amountCents)
		payments = append(payments, Payment{
			PayerName:   d.name,
			PayeeName:   c.name,
			AmountCents: transfer,
		})

		d.amountCents -= transfer
		c.amountCents -= transfer

		if intCloseToZero(d.amountCents, remainderCents) {
			i++
		}
		if intCloseToZero(c.amountCents, remainderCents) {
			j++
		}
	}

	return payments
}
