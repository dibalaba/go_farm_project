package model



//Client structure
type Client struct{

	ID int `json:"id"`
   
	Surname string `json:"surname"`
   
	Firstname string `json:"firstname"`
   
	Phonenumber int `json:"phonenumber"`
   
	Address string `json:"address"`
   
	Credit int `json:"credit"`
   
	Debit int `json:"debit"`
   
   }

