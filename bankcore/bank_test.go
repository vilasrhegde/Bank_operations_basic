package bankcore
import "testing"

//test customer account and name
func TestAccount(t *testing.T) {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
    }

    if account.Name == "" {
        t.Error("can't create an Account object")
    }
}

// test deposit method
func TestDeposit(t *testing.T) {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
    }

    account.Deposit(10)//give negative to get FAIL (-10)

    if account.Balance != 10 {
        t.Error("balance is not being updated after a deposit")
    }
}

//for withdraw
func TestWithdraw(t *testing.T) {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
    }

    account.Deposit(10)
    account.Withdraw(10)

    if account.Balance != 0 {
        t.Error("balance is not being updated after withdraw")
    }
}

//statement
func TestStatement(t *testing.T) {
    account := Account{
        Customer: Customer{
            Name:    "John",
            Address: "Los Angeles, California",
            Phone:   "(213) 555 0147",
        },
        Number:  1001,
        Balance: 0,
    }

    account.Deposit(100)
    statement := account.Statement()
    if statement != "1001 - John - 100" {
        t.Error("statement doesn't have the proper format")
    }
}