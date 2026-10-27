package cards

import (
	"bank/structs"
	"fmt"
)

//Новая дебетовая карта
func NewCard(bank *structs.Bank, accountid int){

	newcard := Card{
		Number: bank.NextCardId,
		Accauntid: bank.Accaunts[accountid].Id,
		IsBlocked: false,
	}

	account := bank.Accaunts[accountid]
	account.UserCards = append(account.UserCards, newcard.Number)
	bank.Accaunts[accountid] = account

	bank.Cards[bank.NextCardId] = newcard
	bank.NextCardId += 1

	fmt.Println("-------Карта офрмлена!-------")
}


//Создание новой кредитки 
func NewCreditCard(bank *structs.Bank, accountid int)(){ 

	newcreditcard := CreditCard{
		Number: bank.NextCardId,
		Accauntid: accountid,
		CreditLimit: CreditLimit(bank, accountid),
		Debt: 0,
	}

	account := bank.Accaunts[accountid]
	account.UserCards = append(account.UserCards, newcreditcard.Number)
	bank.Accaunts[accountid] = account

	bank.Cards[bank.NextCardId] = newcreditcard
	bank.NextCardId += 1

	fmt.Println("-------Карта офрмлена!-------")
}


//Расчёт кредитного лимита
func CreditLimit(bank *structs.Bank, accountid int)(lim float64){

	var debts float64  

	bal := bank.Accaunts[accountid].Balance

	for _, cardid := range bank.Accaunts[accountid].UserCards{
		card, ok := bank.Cards[cardid].(CreditCard)
			if ok!=true{
				continue
			}

		debts += card.Debt
	}

	lim = (bal)*1.4 - debts

	return 
}
