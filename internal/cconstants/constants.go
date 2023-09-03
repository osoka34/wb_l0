package cconstants

import "time"

const (
	PaymentDB  = " public.payment "
	ItemDB     = " public.item "
	OrdersDB   = " public.orders "
	DeliveryDB = " public.delivery "
)

const (
	Ok = iota + 1000
	CantInsertOrder
	CantInsertPayment
	CantInsertDelivery
	CantInsertItem
	CantSelectOrders
	CantSelectPayment
	CantSelectDelivery
	CantSelectItems
	OrderNotFound
)

const (
	GoRoutineSleepSeconds time.Duration = 20
)
