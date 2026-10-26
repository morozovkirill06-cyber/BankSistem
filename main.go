package main

import (
	"fmt"
)


func main() {
	id := 0
	var d int
	var a int
	var ids int
	var err error
	bank := Bank{
		Name:     "GOBANK",
		Accaunts: make(map[int]Accaunt),
		History:  make(map[int]Tranzaction),
		Cards:    make(map[int]any),
		NextUserId:   1,
	}
	client := Client{
		Name: "Kirill",
		Age:  19,
	}
	account := Accaunt{
		Id:      id,
		Owner:   client,
		Balanse: 1000,
		Blocked: false,
	}
	bank.Accaunts[id] = account
	for{
		fmt.Println("========== Главное меню ==========")
		fmt.Println("Создать пользователя - 1")
		fmt.Println("Все пользователи - 2")
		fmt.Println("Войти в аккаунт - 3")
		fmt.Printf("Номер действия:")
		fmt.Scan(&d)
		switch d {
		case 1:
			NewAccaunt(&bank)
			
		case 2:
			AllUsers(bank.Accaunts)

		case 3:
			fmt.Println("_____________________________")
			fmt.Printf("Введите ID аккаунта: ")
			fmt.Scan(&ids)
			for{
				account, ok:= bank.Accaunts[ids]
				if ok != true{
					fmt.Println("Аккаунт не найден")
					continue
				}
			fmt.Println("========== Аккаунт ==========")
			fmt.Printf("ID: %v, Баланс: %v\n", account.Id, account.Balanse)
			fmt.Printf("Имя: %v,  Возраст: %v\n", account.Owner.Name, account.Owner.Age)
			fmt.Println("Создать карту - 1")
			fmt.Println("Создать кредитку - 2")
			fmt.Println("Посмотреть карты - 3")
			fmt.Println("Пополнить карты - 4")
			fmt.Println("В главное меню - 5")
			fmt.Println("Перевести по ID - 7")
			fmt.Printf("Введите действие:")
			fmt.Scan(&a)
			switch a {
			case 1:
				NewCard(&bank, ids)
				continue
			case 2:
				NewCreditCard(&bank, ids)
				continue
			case 3:
				ViewUserCards(&bank, ids)
				continue
			case 4:
				err = UpDepCard(&bank)
				if err!=nil{fmt.Println(err)}
				continue
			case 6: 

			case 7:
				err = TranzactionForId(&bank, ids)
				if err!=nil{fmt.Println(err)}
				continue
			default:
			}
			break
		}	
		continue
		case 4:
		}

	}
}

