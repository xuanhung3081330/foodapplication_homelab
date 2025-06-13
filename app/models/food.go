package models

import (
	"time"
)

type Food struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Category Category `json:"category"`
	Price float32 `json:"price"`
	Image_Url string `json:"image_url"`
	Video_Url string `json:"video_url"`
	Deleted bool `json:"delete"`
	Created_At time.Time `json:"created_at"`
	Created_By string `json:"created_by"`
	Updated_At time.Time `json:"updated_at"`
	Updated_By string `json:"updated_by"`
}
