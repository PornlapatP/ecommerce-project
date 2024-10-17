package model

import "ecommerce-backend/internal/constant"

type RequestLogin struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
	// Surname  string `binding:"required"`
	// Email string `binding:"required"`
}

type RequestCreateUser struct {
	Username string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type RequestGetUserByID struct {
	ID uint `json:"id"`
}

type RequestUpdateRole struct {
	Role constant.UserRole
}

type RequestCreateProduct struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	// ImageURL    string  `json:"imageurl"`
}
