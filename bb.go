package main
import(
	"bank/structs"
	"bank/cards"
)

func bb()(bank structs.Bank){
	id := 0
	bank = structs.Bank{
		Name:     "GOBANK",
		Accaunts: make(map[int]structs.Accaunt),
		History:  make(map[int]structs.Tranzaction),
		Cards:    make(map[int]any),
		NextUserId:   1,
	}
		client := structs.Client{
			Name: "Kirill",
			Age:  19,
		}

			account := structs.Accaunt{
				Id:      id,
				Owner:   client,
				Balance: 1000,
				Blocked: false,
			}
			bank.Accaunts[id] = account
			cards.NewCard(&bank, id)
			cards.NewCreditCard(&bank, id)
			id += 1




	clien := structs.Client{
		Name: "Dmitry",
		Age:  27,
	}

	accoun := structs.Accaunt{
		Id:      id,
		Owner:   clien,
		Balance: 2000,
		Blocked: false,
	}
	bank.Accaunts[id] = accoun

	cards.NewCard(&bank, id)
	cards.NewCreditCard(&bank, id)
	
	return
}