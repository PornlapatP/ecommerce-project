package middleware

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint, username string, role string, secret string) (string, error) {
	// สร้าง claims สำหรับ JWT
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(100 * time.Second).Unix(), // ระยะเวลาหมดอายุ
	}

	// สร้าง JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// เซ็นโทเค็น
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("error signing token:", err)
		return "", err
	}

	return signedToken, nil
}
