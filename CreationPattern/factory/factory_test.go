package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	// The index number of Cash is 1
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Errorf("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment method message wasn't exist")
	}
	t.Log("Log:", msg)
}

func TestCreatePaymentDebitCard(t *testing.T) {
	// The index number of DebitCard is 2
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Errorf("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using debitCard") {
		t.Errorf("The debitCard payment method message wasn't exist")
	}
	t.Log("Log", msg)

}

func TesetGetPaymentMethodNonExistent(t *testing.T) {
	payment, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
		t.Log("LOG:", payment)
	}
	t.Log("LOG:", err)
}
