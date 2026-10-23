package main
import "fmt"


func (card Card) ViewCard(){
	fmt.Println("_____________________________")
	fmt.Printf("Номер: %v \n", card.Number)
	fmt.Printf("Баланс: %v \n", card.Balanse)
	fmt.Printf("Пользователь: %v \n", card.Owner.Name)
		if card.IsBlocked == false{
			fmt.Println("Активна!")
		}else {
			fmt.Println("Заблокирована!")
		}
	}

func (card CreditCard) ViewCard(){
	fmt.Println("_____________________________")
	fmt.Printf("Номер: %v \n", card.Number)
	fmt.Printf("Пользователь: %v \n", card.Owner.Name)
	fmt.Printf("Кредитный лимит: %v \n", card.CreditLimit)
	fmt.Printf("Задолжность: %v \n", card.Debt)
	fmt.Println("_____________________________")
	}

type View interface{
	ViewCard()
}

func PrintCards(card View){
	card.ViewCard()
}

func ViewCards(account *Accaunt){

	for _, card := range account.Cards{
		PrintCards(card)
		fmt.Println("----------------------------")
	}
	for _, creditcard := range account.CreditCards{
		PrintCards(creditcard)
		fmt.Println("----------------------------")
	}
}



