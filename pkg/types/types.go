package types



//Money - представялет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

// Currency - представляет собой код валюты
type Currency string


// Коды валют

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"

)

// PAN - представялет номер карты
type PAN string


type Card struct {
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
}






type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
	
   }



// Category представляет собой катенорию в которой был совершён платёж
type Category string

// Payment представляет информацию о платежах
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}

//Статус предстваляет собой статус платежа
type Status string

//Предопределённые статусы платежей
const(
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS" 
)