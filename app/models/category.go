package models

type Category int

const (
	SavoryDish Category = iota + 1 // EnumIndex = 1
	Dessert                        // EnumIndex = 2
)

func (c Category) String() string {
	switch c {
	case SavoryDish:
		return "SavoryDish"
	case Dessert:
		return "Dessert"
	default:
		return "Unknown"
	}
}
