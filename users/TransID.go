package users
import (
	"fmt"
	"bank/structs"
	"errors"
)


func TranzactionForId(bank *structs.Bank, accountid int)error{

	var ids int
	var sum float64

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
		fmt.Printf("Введите Сумму: ")
		fmt.Scan(&sum)

			if sum <= 0{
				fmt.Println("Сумма меньше или 0...")
				continue

			}
			break
	}

		if sum > account1.Balance {
		return errors.New("Недостаточно средств")
		}
		
	account1.Balance -= sum
	bank.Accaunts[accountid] = account1

	account2.Balance += sum
	bank.Accaunts[ids] = account2

	fmt.Println("Перевод прошел успешно!")

	return nil
}