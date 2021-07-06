package accounts

// BankAccount struct (이것의 단점은 외부에서 Account 객체에 접근이 가능하도록 열려있게됨.)
// type Account struct {
// 	Owner   string
// 	Balance int
// }

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) *Account {
	a.balance += amount

	return &a
}

// Balance of your account
func (a Account) GetBalance() int {
	return a.balance
}
