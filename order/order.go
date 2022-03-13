package order

import (
	"api_hackathon/api"
	"api_hackathon/app"
	"time"
)

type Order struct {
	ID              int       `json:"id,omitempty"`
	Status          int       `json:"status"`
	FirstName       string    `json:"first_name" db:"first_name"`
	LastName        string    `json:"last_name" db:"last_name"`
	Number          string    `json:"number"`
	Email           string    `json:"email"`
	Info            string    `json:"info"`
	Product         string    `json:"product"`
	Count           int       `json:"count"`
	Price           int       `json:"price"`
	Created         time.Time `json:"created"`
	DeliveryAddress string    `json:"delivery_address" db:"delivery_address"`
}

func (i Order) New() Order {
	err := app.DB.Get(&i.ID, "INSERT INTO orders (status,first_name,last_name,number,email,info,product,count,price,created,delivery_address)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id",
		i.Status, i.FirstName, i.LastName, i.Number, i.Email, i.Info, i.Product, i.Count, i.Price, i.Created, i.DeliveryAddress)
	api.CheckErrInfo(err, "order new")
	return i
}

func Get(id int) (Order, error) {
	i := Order{}
	err := app.DB.Get(&i, "select * from orders where id=$1", id)
	api.CheckErrInfo(err, "order get")
	return i, err
}

func Orders(status int) ([]Order, error) {
	var i Order
	var orders = []Order{}
	rows, err := app.DB.Queryx("SELECT * FROM orders where status=$1 ORDER BY created DESC;", status)
	api.CheckErrInfo(err, "notice1")
	for rows.Next() {
		err = rows.StructScan(&i)
		api.CheckErrInfo(err, "notice22")
		orders = append(orders, i)
		//	text := fmt.Sprintf("%s, в поиске сражения! Начать атаку?\n\n⛔ Управление уведомлениями: Прочее -> Настройки", user.GetNick(id))
		//	bb := k.ButtonText("🔥 В атаку!", "attack", "positive", vk.H{"id": api.ToString(idd)})
		//	params := k.New(true, []k.Button{bb})
		//	vkapi.SendMessagee(text, i, params)
	}
	rows.Close()
	api.CheckErrInfo(err, "close notice attack")
	return orders, err
}

func SetOrderStatus(id, v int) (bool, error) {
	_, err := app.DB.Exec("update orders set status=$1 where id=$2", v, id)
	api.CheckErrInfo(err, "SetOrderStatus")
	return true, err
}

func DeleteOrder(id int) (bool, error) {
	_, err := app.DB.Exec("delete from orders  where id=$1", id)
	api.CheckErrInfo(err, "DeleteOrder")
	return true, err
}
