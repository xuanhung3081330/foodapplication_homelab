package models

type OrderStatus int

const (
	OrderPlaced OrderStatus = iota + 1 // EnumIndex = 1
	Confirmed
	Preparing
	ReadyForPickup
	PickedUp
	OutForDelivery
	Delivered
	Cancelled
	FailedDelivery
	Refunded
	Delayed
)

func (s OrderStatus) String() string {
	switch s {
	case OrderPlaced:
		return "Order placed"
	case Confirmed:
		return "Confirmed order"
	case Preparing:
		return "Preparing the food"
	case ReadyForPickup:
		return "Foods finished cooking. Prepare to be picked up"
	case PickedUp:
		return "Shipper has picked up the order"
	case OutForDelivery:
		return "Shipping"
	case Delivered:
		return "Foods received"
	case Cancelled:
		return "Order cancelled"
	case FailedDelivery:
		return "Delivery failed"
	case Refunded:
		return "Money refunded"
	case Delayed:
		return "Order delayed"
	default:
		return "Unknown"
	}
}
