package main

import "fmt"

func (card *Card) Dep(sum float64) {
	card.Balanse += sum
}
func (creditcard *CreditCard) Dep(sum float64) {
	creditcard.Debt -= sum
}
type Deps interface {
	Dep(float64)
}
func Deposit(i Deps, sum float64) (err error) {
	i.Dep(sum)
	return
}


func NewCard(bank *Bank, accountid int){
	newcard := Card{
		Number: bank.NextCardId,
		Balanse: 0,
		Accauntid: bank.Accaunts[accountid].Id,
		IsBlocked: false,
	}
	account := bank.Accaunts[accountid]
	account.UserCards = append(account.UserCards, newcard.Number)
	bank.Accaunts[accountid] = account
	bank.Cards[bank.NextCardId] = newcard
	bank.NextCardId += 1
	fmt.Println("Карта создана!")
}

func NewCreditCard(bank *Bank, accountid int)(){ 
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

func(creditcard CreditCard) GetMoney()(float64, float64){
	return 0, creditcard.Debt
}
func(creditcard Card) GetMoney()(float64, float64){
	return creditcard.Balanse, 0
}
type FinansCard interface{
	GetMoney()(float64, float64)
}
func FinansInfo(i FinansCard)(float64, float64){
	return i.GetMoney()
}


func CreditLimit(bank *Bank, accountid int)(lim float64){
	bal := bank.Accaunts[accountid].Balanse
	var debts float64 
	var bals float64
	for _, cardid := range bank.Accaunts[accountid].UserCards{
		cardid := bank.Cards[cardid]
		cid, ok := cardid.(FinansCard)
			if ok!=true{
				return
			}
		bals, debts = FinansInfo(cid)
	}
	lim = (bal + bals)*1.4 - debts
	return 
	}
