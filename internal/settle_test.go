package internal

import (
	"reflect"
	"testing"
)

func TestSettle(t *testing.T) {
	t.Run("No spendings", func(t *testing.T) {
		payments, _ := SettleDown([]Spending{})
		if len(payments) != 0 {
			t.Fatal("Want no payments")
		}
	})

	t.Run("Two people, one pays everything", func(t *testing.T) {
		payments, _ := SettleDown([]Spending{
			{"Alice", 50},
			{"Bob", 0},
		})

		if len(payments) != 1 {
			t.Fatalf("Want 1 payment, got %d", len(payments))
		}

		got := payments[0]
		want := Payment{PayerName: "Bob", PayeeName: "Alice", AmountCents: 25}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Want %v, got %v", want, got)
		}
	})

	t.Run("Two people, each pays the same", func(t *testing.T) {
		payments, _ := SettleDown([]Spending{
			{"Alice", 50},
			{"Bob", 50},
		})

		if len(payments) != 0 {
			t.Fatal("Want no payments")
		}
	})

	t.Run("Four people, non-divisible total", func(t *testing.T) {
		// Total = 403 cents. 100 cents each, and a 3 cent surplus that's used
		// as tolerance but not added to anyone.
		payments, _ := SettleDown([]Spending{
			{"Alice", 203},
			{"Bob", 150},
			{"Candance", 50},
			{"Denisse", 0},
		})

		if len(payments) != 2 {
			t.Fatalf("Want 2 payments, got %d", len(payments))
		}

		wantPayments := []Payment{
			{PayerName: "Denisse", PayeeName: "Alice", AmountCents: 100},
			{PayerName: "Candance", PayeeName: "Bob", AmountCents: 50},
		}
		if !reflect.DeepEqual(wantPayments, payments) {
			t.Fatalf("Want %v, got %v", wantPayments, payments)
		}
	})

	t.Run("Big group", func(t *testing.T) {
		payments, _ := SettleDown([]Spending{
			{"olaia", 250},
			{"pablo", 0},
			{"raquel", 1480},
			{"iñigo", 0},
			{"angel", 680},
			{"jen", 0},
			{"vanessa", 0},
			{"jose", 0},
		})

		wantPayments := []Payment{
			{PayerName: "pablo", PayeeName: "raquel", AmountCents: 301},
			{PayerName: "iñigo", PayeeName: "raquel", AmountCents: 301},
			{PayerName: "jen", PayeeName: "raquel", AmountCents: 301},
			{PayerName: "vanessa", PayeeName: "raquel", AmountCents: 276},
			{PayerName: "vanessa", PayeeName: "angel", AmountCents: 25},
			{PayerName: "jose", PayeeName: "angel", AmountCents: 301},
			{PayerName: "olaia", PayeeName: "angel", AmountCents: 51},
		}
		if !reflect.DeepEqual(wantPayments, payments) {
			t.Fatalf("Want %v, got %v", wantPayments, payments)
		}
	})
}
