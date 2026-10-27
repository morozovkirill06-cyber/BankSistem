package cards

import(
	"fmt"
	"errors"
	"bank/structs"
)

type Card struct {
	Number     int
	Accauntid  int // Владелец
	IsBlocked  bool    // Заблокирована ли карта
}
type CreditCard struct {
	Number      int  // Номер карты
	Accauntid   int  // Владелец
	CreditLimit float64 // Кредитный лимит
	Debt        float64 // Задолженность
}

func (c Card)UpDep(bank *structs.Bank, sum float64){

	id := c.Accauntid

	account := bank.Accaunts[id]

	account.Balance +=sum

	bank.Accaunts[id] = account
	}

func (ccr CreditCard)UpDep(bank *structs.Bank, sum float64){

	id := ccr.Accauntid
	account := bank.Accaunts[id]

    if ccr.Debt == sum{
		fmt.Println("Задолжность погашена!")
		ccr.Debt = 0

	} else if ccr.Debt > sum {
		ccr.Debt -= sum
		fmt.Printf("Осталась задолжность %v\n ", ccr.Debt)

	} else {
		PlusBalanc := (ccr.Debt - sum)*-1
		account.Balance += PlusBalanc

		ccr.Debt = 0

		fmt.Printf("Задолжность погашена и остаток %v переведён на ваш счёт.\n", PlusBalanc )
	}

	bank.Accaunts[id] = account
}

type Deposit interface{
	UpDep(*structs.Bank, float64)
}

func UpDepCard(bank *structs.Bank)error{

	var sum float64
	var cid int

	fmt.Printf("Введите номер карты пополнения: ")
	fmt.Scan(&cid)

	_, ok := bank.Cards[cid]
		if !ok {
			return errors.New("Карта не найдена")
		}

	for{
		fmt.Printf("Введите Сумму: ")
		fmt.Scan(&sum)

			if sum <= 0{
				fmt.Println("Сумма меньше или 0...")
				continue

			}
			break
	}
	card, ok := bank.Cards[cid].(Deposit)
	if !ok{
		return errors.New("Не верный тип")
	}

	card.UpDep(bank, sum)

	bank.Cards[cid] = card

	return nil
}



func (c Card)DownDep(bank *structs.Bank, sum float64)error{

	id := c.Accauntid
	account := bank.Accaunts[id]

		if sum > account.Balance{
			return errors.New("Недостаточно средств")
		}

	account.Balance -= sum

	bank.Accaunts[id] = account
	return nil
	}


func (ccr CreditCard)DownDep(bank *structs.Bank, sum float64)error{

	id := ccr.Accauntid
	account := bank.Accaunts[id]
	creditlimit := CreditLimit(bank, id)

    if creditlimit < sum{
		return errors.New("Лимит превышен!")

	} else if ccr.Debt > sum {
		ccr.Debt -= sum
		fmt.Printf("Осталась задолжность %v\n ", ccr.Debt)

	} else {
		PlusBalanc := (ccr.Debt - sum)*-1
		account.Balance += PlusBalanc

		ccr.Debt = 0

		fmt.Printf("Задолжность погашена и остаток %v переведён на ваш счёт.\n", PlusBalanc )
	}

	bank.Accaunts[id] = account
	return nil
}

type Send interface{
	DownDep(*structs.Bank, float64)error
}

func TranzactionForCard(bank *structs.Bank)(err error){

	var getCardId int
	var sendCardId int

	fmt.Printf("Введите номер карты отправления: ")
	fmt.Scan(&sendCardId)
	fmt.Printf("Введите номер карты пополнения: ")
	fmt.Scan(&getCardId)


	return
}