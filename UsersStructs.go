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

	Cards []Card
	CreditCards []CreditCard
}
type Bank struct{
	Name string
	Accaunts map[int]Accaunt
	History map[int]Tranzaction
	NextId int
}