package account

import (
	client "module/src/account/clients"
	"strconv"
)

type DebitAccount struct {
	PersonalClient                         client.ClientData
	AgencyNumber, AccountNumber, Operation int32
	balance                                float64
}

func (account *DebitAccount) WithdrawFromAccount(value float64) string {
	canWithdraw := value > 0 && value <= account.balance
	if canWithdraw {
		account.balance -= value
		return "Successfully done"
	} else {
		return "You can't withdraw the money"
	}
}

func (account *DebitAccount) DepositToAccount(value float64) string {
	if value > 0 {
		account.balance += value
		return strconv.FormatFloat(value, 'f', 2, 64) + " added to your account"
	} else {
		return "The value has to be bigger than R$0,00"
	}
}

func (account *DebitAccount) TransferFromAccounts(value float64, transfAccount *DebitAccount) string {
	if value > 0 && account.balance > value {
		account.balance -= value
		transfAccount.DepositToAccount(value)
		return "The value of your balance now is R$" + strconv.FormatFloat(transfAccount.balance, 'f', 2, 64)
	} else {
		return "You cannot transfer this amount! Try it again with other value!"
	}
}

func (account *DebitAccount) GetBalance() float64 {
	return account.balance
}
