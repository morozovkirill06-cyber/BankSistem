package cards

import ( 
	"fmt"
	"bank/structs"
)

//Информация по картам
func (card Card) ViewCard(){
	fmt.Println("_____________________________")
	fmt.Printf("Номер: %v \n", card.Number)
	fmt.Printf("Пользовательский ID: %v \n", card.Accauntid)
		if card.IsBlocked == false{
			fmt.Println("Активна!")
		}else {
			fmt.Println("Заблокирована!")
		}
	}
//Информация по кредитной карте
func (card CreditCard) ViewCard(){
	fmt.Println("_____________________________")
	fmt.Printf("Номер: %v \n", card.Number)
	fmt.Printf("Пользовательский ID: %v \n", card.Accauntid)
	fmt.Printf("Кредитный лимит: %v \n", card.CreditLimit)
	fmt.Printf("Задолжность: %v \n", card.Debt)
	}

type View interface{
	ViewCard()
}

func PrintCards(card View){
	card.ViewCard()
}
//Отображение информации по картам конкретного пользователя
func ViewUserCards(bank *structs.Bank, accountid int){
	if len(bank.Accaunts[accountid].UserCards) == 0{
		fmt.Println("Нет карт...")
	}
	for _, cid := range bank.Accaunts[accountid].UserCards{
		card, ok := bank.Cards[cid].(View)
			if ok == false{
				return
			}
		PrintCards(card)
		fmt.Println("______________________________________")
	}
}