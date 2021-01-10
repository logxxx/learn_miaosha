package datamodels

type Order struct {
	Id int64 `sql:"id"`
	UserId int64 `sql:"user_id"`
	ProductId int64 `sql:"product_id"`
	OrderStatus int64 `sql:"order_status"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)