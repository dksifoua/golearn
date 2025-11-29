package pointersanderrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestBankOperations(t *testing.T) {
	a := &Account{0, Dollar(25.0)}
	b := &Account{0, Dollar(00.0)}

	t.Run("Transfer", func(t *testing.T) {
		err := a.transfer(b, 10.0)
		if err != nil {
			t.Fatal("Something wrong happens.")
		}

		if a.Balance != Dollar(15.0) || b.Balance != Dollar(10.0) {
			errorMessage := fmt.Sprintf("Account %+v expected balance %v\n", a, Dollar(15.0))
			errorMessage += fmt.Sprintf("Account %+v expected balance %v", b, Dollar(10.0))
			t.Error(errorMessage)
		}
	})

	t.Run("Transfer insufficient funds", func(t *testing.T) {
		err := a.transfer(b, 25.0)
		if err == nil {
			t.Fatal("Something wrong happens.")
		}

		expected := errors.New("insufficient funds")
		if err.Error() != expected.Error() {
			t.Errorf("Expected: %+v - Actual: %+v", expected, err)
		}
	})
}
