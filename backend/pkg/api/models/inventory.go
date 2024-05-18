package models


type Inventory struct{
	ID int
	Item string 
	Quantinty int
}

type Invoice struct{
	ID int
	SID int
	GST float64
	TotalAmount float64
	PaymentMode string 
}