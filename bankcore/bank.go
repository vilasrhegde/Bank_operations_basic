package bankcore

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
    Name    string
    Address string
    Phone   string
}

// Account ...
type Account struct {
    Customer
    Number  int32
    Balance float64
}
// Deposit ...
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to deposit should be greater than zero")
    }

    a.Balance += amount
    return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("the amount to withdraw should be greater than zero")
    }

    if a.Balance < amount {
        return errors.New("the amount to withdraw should be greater than the account's balance")
    }

    a.Balance -= amount
    return nil
}

// Statement ...
// func (a *Account) Statement() string {
//     return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
// }
func (a Account) Statement() string {
	jsonString, _ := json.Marshal(a)
	return string(jsonString)
}
type Statementer interface{
    Statement() string
}
func Statement(s Statementer) {
	fmt.Println(s.Statement())
}