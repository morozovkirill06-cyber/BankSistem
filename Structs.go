package main

type Client struct{
	Name string		  
	Age int			  
	}
type Tranzaction struct{
	Id 		int     
	Type	string 	
	FromId  int	 	 
	ToId    int		
	Amount  float64  
}
type Accaunt struct{
	Id int
	Owner Client
	Balanse float64
	Blocked bool
	UserCards []int
}
type Bank struct{
	Name string
	Accaunts map[int]Accaunt
	History map[int]Tranzaction
	Cards map[int]any
	NextUserId int
	NextCardId int
}

type Card struct {
	Number     int
	Balanse    float64 // Номер карты
	Accauntid  int // Владелец
	IsBlocked  bool    // Заблокирована ли карта
}
type CreditCard struct {
	Number      int  // Номер карты
	Accauntid   int  // Владелец
	CreditLimit float64 // Кредитный лимит
	Debt        float64 // Задолженность
}
