package types





//! попытка в 14:41


//Money - представялет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

type Category string


// новый Status представляет собой статус платежа
type Status string


const (
	StatusOk Status = "OK"
	StatusFail 	Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// new Payment - представляет информацию о платеже.
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}