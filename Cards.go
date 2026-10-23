package main

import "fmt"

type Card struct {
	Number    int
	Balanse   float64 // Номер карты
	Owner     Client  // Владелец
	IsBlocked bool    // Заблокирована ли карта
}
type CreditCard struct {
	Number      int  // Номер карты
	Owner       Client  // Владелец
	CreditLimit float64 // Кредитный лимит
	Debt        float64 // Задолженность
}
func (card *Card) Dep(sum float64) {
	card.Balanse += sum
}
func (creditcard *CreditCard) Dep(sum float64) {
	creditcard.Debt += sum
}
type Deps interface {
	Dep(float64)
}
func Deposit(i Deps, sum float64) (err error) {
	i.Dep(sum)
	return
}

func NewCard(account *Accaunt, idcard *int){
	var balanse float64
	fmt.Println("Введите баланс карты:")
	fmt.Scan(&balanse)
	newcard := Card{
		Number: *idcard,
		Balanse: balanse,
		Owner: account.Owner,
		IsBlocked: false,
	}
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
	account.CreditCards = append(account.CreditCards, newcreditcard)
	fmt.Println("-------Карта офрмлена!-------")
}

func CreditLimit(account *Accaunt)(lim float64){
	var bal float64
	var debt float64
	for _, card := range account.Cards{
		bal += card.Balanse
	}
	for _, creditcard := range account.CreditCards{
		debt += creditcard.Debt
	}
	k := bal/(0-debt)/2
	lim = k*account.Balanse
	
	return
}