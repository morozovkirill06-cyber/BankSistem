package main

import"fmt"

func TranzactionForId(bank *Bank, accountid int){
	var sum float64
	var ids int
	fmt.Printf("Введите ID аккаунта для получения средств: ")
	fmt.Scan(&ids)
	account1 := bank.Accaunts[accountid]
	account2, ok:= bank.Accaunts[ids]
		if ok != true{
			fmt.Println("Аккаунт не найден")
			return
		}else if ids == accountid{
			fmt.Println("Самоме себе нельзя перевести")
			return
		}
	for{
	fmt.Printf("Введите Сумму: ")
	fmt.Scan(&sum)
		if sum > account1.Balanse{
			fmt.Println("Недостаточно средств...")
			continue
		}else if sum <= 0{
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
}

func (c *Card)UpDep(bank *Bank, sum float64){
id := c.Accauntid
account := bank.Accaunts[id]
account.Balanse +=sum
bank.Accaunts[id] = account
}

func (ccr *CreditCard)UpDep(bank *Bank, sum float64){
id := ccr.Accauntid
account := bank.Accaunts[id]
    if ccr.Debt == sum{
		fmt.Println("Задолжность погашена!")
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

func TranzactionForCard(i Deposit, bank *Bank,  sum float64){
	i.UpDep(bank, sum)
}
