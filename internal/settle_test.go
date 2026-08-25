package internal

import (
	"reflect"
	"testing"
)

func TestSettle(t *testing.T) {
	t.Run("No spendings", func(t *testing.T) {
		payments := SettleDown([]Spending{})
		if len(payments) != 0 {
			t.Fatal("Want no spendings")
		}
	})

	t.Run("Two people, one pays everything", func(t *testing.T) {
		payments := SettleDown([]Spending{
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
}
