package structs

type Client struct{
	Name		     string		  
	Age				 int			  
	}
type Tranzaction struct{
	Id 				 int     
	Type			 string 	
	FromId			 int	 	 
	ToId   			 int		
	Amount 			 float64  
}
type Accaunt struct{
	Id 			     int
	Owner 			 Client
	Balance 		 float64
	Blocked			 bool
	UserCards		 []int
}
type Bank struct{
	Name			 string
	Accaunts		 map[int]Accaunt
	History			 map[int]Tranzaction
	Cards 			 map[int]any
	DebitCardid  	 []int
	CreditCardCardid []int
	NextUserId 		 int
	NextCardId 		 int
}

