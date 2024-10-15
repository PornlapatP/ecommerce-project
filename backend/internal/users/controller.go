package users

import (
	"ecommerce-backend/internal/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB, secret string, hashsecret string) Controller {
	return Controller{
		Service: NewService(db, secret, hashsecret),
	}
}

func (ctrl Controller) Createuser(ctx *gin.Context) {
	var (
		request model.RequestCreateUser
	)
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Bind body").Error(),
		})
		return
	}

	result, err := ctrl.Service.CreateUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Create User").Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.ResponseUser[any]{
		Message: "User created successfully",
		Userid:  result.ID,
	})
}

func (ctrl Controller) Login(ctx *gin.Context) {
	var (
		request model.RequestLogin
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := ctrl.Service.Login(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.SetCookie(
		"token",
		fmt.Sprintf("Bearer %v", token), int(100*time.Minute),
		"/",
		"localhost",
		false,
		true,
	)
	fmt.Printf("Generated Token: %s\n", token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login succeeded",
		"token":   token,
	})

}
func (ctrl Controller) GetUserById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	var (
		request model.RequestGetUserByID
	)

	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Validata").Error(),
		})
		return
	}
	result, err := ctrl.Service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Find user by Id").Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, model.ResponseUser[any]{
		Userid:    result.ID,
		Username:  result.Username,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	})

}
func (ctrl Controller) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: "Invalid ID format",
		})
		return
	}

	var (
		request model.RequestCreateUser
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Bind body").Error(),
		})
		return
	}

	// log.Printf("req:%v", request)
	_, err = ctrl.Service.UpdateUser(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Update User").Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseUser[any]{
		Message: "User updated successfully",
	})

}

func (ctrl Controller) UpdateRole(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	var (
		request model.RequestUpdateRole
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Bind body").Error(),
		})
		return
	}
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}
	log.Printf("id:%v", id)
	_, err := ctrl.Service.UpdateRole(uint(id), request.Role)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Update User Role").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[any]{
			Message: "User updated Role successfully",
		},
	)
}
