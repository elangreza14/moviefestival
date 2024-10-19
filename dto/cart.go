package dto

import "github.com/google/uuid"

type (
	AddCartPayload struct {
		movieID int `json:"movie_id" binding:"required"`
		UserID  uuid.UUID
	}

	CartListResponseElement struct {
		ID          int    `json:"id"`
		Quantity    int    `json:"quantity"`
		Message     string `json:"message"`
		movieID     int    `json:"movie_id"`
		movieName   string `json:"movie_name"`
		movieImage  string `json:"movie_image"`
		moviePrice  int    `json:"movie_price"`
		ActualStock int    `json:"actual_stock"`
	}

	CartListResponse []CartListResponseElement

	CheckoutCart struct {
		CartIDs []int `json:"cart_ids" binding:"required,gt=0"`
		UserID  uuid.UUID
	}
)
