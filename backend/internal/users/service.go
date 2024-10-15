package users

import (
	"ecommerce-backend/internal/constant"
	"ecommerce-backend/internal/middleware"
	"ecommerce-backend/internal/model"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	secret     string
	hashsecret string
}

func NewService(dbconn *gorm.DB, secret string, hashsecret string) Service {
	return Service{
		Repository: NewRepository(dbconn),
		secret:     secret,
		hashsecret: hashsecret,
	}
}

func (service Service) CreateUser(req model.RequestCreateUser) (model.User, error) {
	secret := service.hashsecret
	now := time.Now()

	hashedPassword, err := HashPassword(req.Password, secret)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password")
	}

	user := model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      constant.UserRoleStatus,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := service.Repository.CreateUser(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (service Service) Login(req model.RequestLogin) (string, error) {
	user, err := service.Repository.FindOneByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid Email")
	}

	if ok := service.CheckPasswordHash(req.Password, user.Password); !ok {
		return "", errors.New("Invalid password")
	}

	token, err := middleware.CreateToken(user.ID, user.Username, string(user.Role), service.secret)
	if err != nil {
		log.Println("Fail to create token")
		return "", errors.New("Something went wrong")
	}
	return token, nil

}

func (service Service) CheckPasswordHash(password, hash string) bool {
	secret := service.hashsecret

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+secret))
	return err == nil
}

func (service Service) GetByID(id uint) (model.User, error) {
	user, err := service.Repository.GetByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (service Service) UpdateUser(id uint, req model.RequestCreateUser) (model.User, error) {
	result, err := service.Repository.GetByID(id)

	secret := service.hashsecret
	hashedPassword, err := HashPassword(req.Password, secret)

	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password")
	}

	// log.Printf("req:%v", req.Email)
	// log.Printf("req:%v", req.Username)

	user := model.User{
		ID:        result.ID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      result.Role,
		CreatedAt: result.CreatedAt,
		UpdatedAt: time.Now(),
	}
	// log.Printf("user:%v", user)
	if err := service.Repository.Update(user); err != nil {
		return model.User{}, err
	}
	return user, err
}
func (service Service) UpdateRole(id uint, role constant.UserRole) (model.User, error) {
	user, err := service.Repository.GetByID(id)
	if err != nil {
		return model.User{}, err
	}

	user.Role = role
	user.UpdatedAt = time.Now()

	// Replace
	if err := service.Repository.Update(user); err != nil {
		return model.User{}, err
	}

	return user, nil
}
