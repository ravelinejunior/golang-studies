package main

import (
	"fmt"
	"module/src/account"
	client "module/src/account/clients"
)

func main() {
	testingWithdrawInterface()
}

func PayBill(account verifiedAccount, billValue float64) {
	account.WithdrawFromAccount(billValue)
}

type verifiedAccount interface {
	WithdrawFromAccount(billValeu float64) string
}

func testingWithdrawInterface() {
	franklinData := client.ClientData{
		Name:             "Franklin",
		SocialDataNumber: "3161531",
		Profession:       "Engineer",
	}

	franklinAccount := account.DebitAccount{
		PersonalClient: franklinData,
		AgencyNumber:   12542,
		AccountNumber:  1553,
		Operation:      10,
	}

	juniorData := client.ClientData{
		Name:             "Junior",
		SocialDataNumber: "551223",
		Profession:       "Businness",
	}

	juniorAccount := account.PersonalAccount{
		PersonalClient: juniorData,
		AgencyNumber:   14585,
		AccountNumber:  1697,
	}

	franklinAccount.DepositToAccount(1800)
	juniorAccount.DepositToAccount(2040)

	PayBill(&franklinAccount, 600)
	PayBill(&juniorAccount, 1600)

	fmt.Println(&franklinAccount)
	fmt.Println(&juniorAccount)

}

func testingDebitAccount() {
	franklinData := client.ClientData{
		Name:             "Franklin",
		SocialDataNumber: "3161531",
		Profession:       "Engineer",
	}

	franklinAccount := account.DebitAccount{
		PersonalClient: franklinData,
		AgencyNumber:   12542,
		AccountNumber:  1553,
		Operation:      10,
	}

	juniorData := client.ClientData{
		Name:             "Junior",
		SocialDataNumber: "551223",
		Profession:       "Businness",
	}

	juniorAccount := account.DebitAccount{
		PersonalClient: juniorData,
		AgencyNumber:   14585,
		AccountNumber:  1697,
		Operation:      10,
	}

	franklinAccount.DepositToAccount(1800)
	fmt.Println("First deposit", franklinAccount)

	juniorAccount.DepositToAccount(2040)
	fmt.Println("First deposit", juniorAccount)

	franklinAccount.TransferFromAccounts(1200, &juniorAccount)
	fmt.Println("After transfer", juniorAccount, franklinAccount)
}

func testingPersonalAccount() {
	var franklinAccount account.PersonalAccount = account.PersonalAccount{
		PersonalClient: client.ClientData{
			Name:             "Franklin",
			SocialDataNumber: "3161531",
			Profession:       "Engineer",
		},
		AgencyNumber:  1456,
		AccountNumber: 112425,
	}

	fmt.Println(franklinAccount)

	marinaAccount := account.PersonalAccount{
		PersonalClient: client.ClientData{
			Name:             "Marina",
			SocialDataNumber: "0651515",
			Profession:       "Aquaculture Engineer",
		},
		AgencyNumber:  1457,
		AccountNumber: 121218,
	}

	fmt.Println(marinaAccount)

	//using pointers
	var lucieneAccount *account.PersonalAccount
	lucieneAccount = new(account.PersonalAccount)

	lucieneAccount.AccountNumber = 115557
	lucieneAccount.PersonalClient = client.ClientData{Name: "Luciene"}

	fmt.Println(*lucieneAccount)

	// testing withdraw
	fmt.Println(franklinAccount.WithdrawFromAccount(1300))
	fmt.Println(franklinAccount)

	// testing deposit
	fmt.Println(marinaAccount.DepositToAccount(2265.54))
	fmt.Println(marinaAccount)

	// testing transfer
	fmt.Println("Transfering from Franklin to Marina:", franklinAccount.TransferFromAccounts(4000.00, &marinaAccount))
	fmt.Println(franklinAccount)
	fmt.Println(marinaAccount.GetBalance())

}
