package users

import (
	"fmt"
	"bank/structs"
)


//Новый аккаунт
func NewAccaunt(bank *structs.Bank)(err error){
	var name string
	var age int
	var balance float64 
	fmt.Println("Создание нового пользователя.")
	fmt.Printf("Имя: ")
	fmt.Scan(&name)
	for{
		fmt.Printf("Возраст: ")
		if _, err := fmt.Scan(&age); err!=nil{
			fmt.Println("Возраст должен быть числом!")
			continue
		}else if age <= 0{
			fmt.Println("Возраст не может быть меньше 0")
			continue
		}else if age < 14 {
			fmt.Println("Мелковат ещё((")
			continue
		}
		break
	}
	for{
		fmt.Printf("Баланс: ")
		if _, err := fmt.Scan(&balance); err!=nil{
		fmt.Println("Баланс должен быть числом!")
		continue
		}else if balance == 0{
			fmt.Println("Ты студент? Соболезную...")
			break
		}else if balance < 0{
			fmt.Println("Баланс не может быть отрицательным")
			continue
		}
		break
	}
	accaunt := structs.Accaunt{
		Id: bank.NextUserId,
		Owner: structs.Client{
			Name: name,
			Age: age,
			},
		Balance: balance,
		Blocked: false,		
	}
	bank.NextUserId += 1
	bank.Accaunts[accaunt.Id] = accaunt

	fmt.Printf("Здравствуйте, %v!\nВаш аккаунт успешно создан!\n\n", bank.Accaunts[accaunt.Id].Owner.Name)
return 
}


//Все пользователи
func AllUsers(accaunts map[int]structs.Accaunt){
	fmt.Println("========= ПОЛЬЗОВАТЕЛИ =========")
		for _, accaunt := range accaunts{
			fmt.Printf("ID: %v, Баланс: %v\n", accaunt.Id, accaunt.Balance)
			fmt.Printf("Имя: %v,  Возраст: %v\n", accaunt.Owner.Name, accaunt.Owner.Age)
			fmt.Println("--------------------------------")
	}
}