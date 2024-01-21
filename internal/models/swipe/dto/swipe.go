package dto

import "datink/internal/models/swipe/entity"

type SwipingRequest struct {
	SwipeDirection entity.SwipeDirection `json:"swipeDirection" validate:"required"`
	FriendID       int64                 `json:"friendID" validate:"required"`
}

type SwipingResponse struct {
	ID int64 `json:"id"`
}
