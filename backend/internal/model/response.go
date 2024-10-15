package model

import "time"

type BaseResponse[DataType any] struct {
	Message string   `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}
type ResponseUser[DataType any] struct {
	Message   string    `json:"message,omitempty"`
	Data      DataType  `json:"data,omitempty"`
	Userid    uint      `json:"userid,omitempty"`    // ควรเป็น uint
	Username  string    `json:"username,omitempty"`  // ควรเป็น string
	Email     string    `json:"email,omitempty"`     // ควรเป็น string
	CreatedAt time.Time `json:"createdat,omitempty"` // ควรเป็น time.Time
}
