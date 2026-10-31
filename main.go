package main

import (
	"bank/data"
	"bank/cards"
	"bank/users"
	"bank/structs"
	"fmt"
)

func main() {
	var d int
	var a int
	var ids int
	var err error

		bank := structs.Bank{
			Name:     "GOBANK",
			Accaunts:	 make(map[int]structs.Accaunt),
			History: 	 make(map[int]structs.Tranzaction),
			Cards:       make(map[int]any),
			NextUserId:  1,
			NextCardId:  1,
		}

		err = data.ReadFileJson(&bank)
		if err != nil {
			fmt.Println(err)
		}

	for {
		

		fmt.Println("========== Главное меню ==========")
		fmt.Println("Создать пользователя - 1")
		fmt.Println("Все пользователи - 2")
		fmt.Println("Войти в аккаунт - 3")
		fmt.Println("Перевод между картами - 4")
		fmt.Printf("Номер действия:")
		fmt.Scan(&d)

		switch d {
		case 1:
			users.NewAccaunt(&bank)

			err = data.SaveFileJson(&bank)
			if err != nil {
				fmt.Println(err)
				break
			}
			continue

		case 2:
			users.AllUsers(bank.Accaunts)

		case 3:
			fmt.Println("_____________________________")
			fmt.Printf("Введите ID аккаунта: ")
			fmt.Scan(&ids)

			for {
				account, ok := bank.Accaunts[ids]
				if ok != true {
					fmt.Println("Аккаунт не найден")
					break
				}

				fmt.Println("========== Аккаунт ==========")
				fmt.Printf("ID: %v, Баланс: %v\n", account.Id, account.Balance)
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
					cards.NewCard(&bank, ids)

					err = data.SaveFileJson(&bank)
					if err != nil {
						fmt.Println(err)
						break
					}

					continue

				case 2:
					cards.NewCreditCard(&bank, ids)
					
					fmt.Println("-------Карта офрмлена!-------")

					err = data.SaveFileJson(&bank)
					if err != nil {
						fmt.Println(err)
						break
					}

					continue

				case 3:
					cards.ViewUserCards(&bank, ids)
					continue

				case 4:
					err = cards.UpDepCard(&bank)

					if err != nil {
						fmt.Println(err)
						continue
					}

					err = data.SaveFileJson(&bank)
						if err != nil {
							fmt.Println(err)
							break
						}
					continue

				case 6:

				case 7:
					err = users.TranzactionForId(&bank, ids)
					if err != nil {
						fmt.Println(err)
					}

					err = data.SaveFileJson(&bank)
					if err != nil {
						fmt.Println(err)
						break
					}

					continue

				default:
				}
				break
			}
			continue

		case 4:
			err = cards.TranzactionForCard(&bank)
			if err != nil {
				fmt.Println(err)
			}

			err = data.SaveFileJson(&bank)
			if err != nil {
				fmt.Println(err)
				break
			}

			continue
		}

	}
}
