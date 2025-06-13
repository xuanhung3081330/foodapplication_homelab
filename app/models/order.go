package models

import (
	"time"
)

type Order struct {
	Id int `json:"id"`
	Status OrderStatus `json:"status"`
	Foods []OrderFood `json:"foods"`
	Ordered_At time.Time `json:"ordered_at"`
	Ordered_By string `json:"ordered_by"`
	Updated_At time.Time `json:"updated_at"`
	Updated_By string `json:"updated_by"`
}
