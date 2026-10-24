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



func NewCard(bank *Bank, accid int){
	newcard := Card{
		Number: *idcard,
		Balanse: 0,
		Accaunt: bank.Accaunts[accid].Id,
		IsBlocked: false,
	}
	bank.NaxtCardId += 1
	account.Cards = append(account.Cards, newcard)
	fmt.Println("Карта создана")
}
func NewCreditCard(account *Accaunt, idcard *int)(){ 
	newcreditcard := CreditCard{
		Number: *idcard,
		Owner: account.Owner,
		CreditLimit: CreditLimit(account),
		Debt: 0,
	}
	*idcard += 1
	account.CreditCards = append(account.CreditCards, newcreditcard)
	fmt.Println("-------Карта офрмлена!-------")
}

func CreditLimit(account *Accaunt)(lim float64){
	var debt float64
	for _, creditcard := range account.CreditCards{
		debt += creditcard.Debt
	}
	lim = account.Balanse / 2 - debt
	
	return
}