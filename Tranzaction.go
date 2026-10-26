package main

import"fmt"
import"errors"

func Sum() (sum float64) {
		for{
	fmt.Printf("Введите Сумму: ")
	fmt.Scan(&sum)
        if sum <= 0{
			fmt.Println("Сумма меньше или 0...")
			continue
		}
		break
	}
	return
}


func TranzactionForId(bank *Bank, accountid int)error{
	var ids int
	fmt.Printf("Введите ID аккаунта для получения средств: ")
	fmt.Scan(&ids)
	account1 := bank.Accaunts[accountid]
	account2, ok:= bank.Accaunts[ids]
		if ok != true{
			return errors.New("Аккаунт не найден")
		}else if ids == accountid{
			return errors.New("Самоме себе нельзя перевести")
		}
	for{
	sum := Sum() 
	if sum <= 0{
			fmt.Println("Сумма меньше или 0...")
			continue
		}
	account1.Balanse -= sum
	bank.Accaunts[accountid] = account1
	account2.Balanse += sum
	bank.Accaunts[ids] = account2
	fmt.Println("Перевод прошел успешно!")
	break
	}
return nil
}

func (c Card)UpDep(bank *Bank, sum float64){
id := c.Accauntid
account := bank.Accaunts[id]
account.Balanse +=sum
bank.Accaunts[id] = account
}

func (ccr CreditCard)UpDep(bank *Bank, sum float64){
id := ccr.Accauntid
account := bank.Accaunts[id]
    if ccr.Debt == sum{
		fmt.Println("Задолжность погашена!")
		ccr.Debt = 0
	} else if ccr.Debt > sum {
		ccr.Debt -= sum
		fmt.Printf("Осталась задолжность %v\n ", ccr.Debt)
	} else {
		PlusBalans := (ccr.Debt - sum)*-1
		ccr.Debt = 0
		account.Balanse += PlusBalans
		fmt.Printf("Задолжность погашена и остаток %v переведён на ваш счёт.\n", PlusBalans )}
	bank.Accaunts[id] = account
}

type Deposit interface{
	UpDep(*Bank, float64)
}

func UpDepCard(bank *Bank)error{
	var cid int
	fmt.Printf("Введите номер карты пополнения: ")
	fmt.Scanln(&cid)
	_, ok := bank.Cards[cid]
		if !ok {
			return errors.New("Карта не найдена")
		}
	sum := Sum()
	card, ok := bank.Cards[cid].(Deposit)
	if !ok{
		return errors.New("Не верный тип")
	}
	card.UpDep(bank, sum)
	bank.Cards[cid] = card
	return nil
}
