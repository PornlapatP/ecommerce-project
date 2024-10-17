package products

import (
	"ecommerce-backend/internal/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(dbconn *gorm.DB) Controller {
	return Controller{
		Service: NewService(dbconn),
	}
}

func (ctrl Controller) CreateProduct(ctx *gin.Context) {
	var (
		request model.RequestCreateProduct
	)

	// Bind the request body
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Bind body").Error(),
		})
		return
	}

	// Handle file upload
	file, err := ctx.FormFile("ImageURL")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// Save the file to the 'uploads' folder
	savePath := fmt.Sprintf("./uploads/%s", filename) // Adjusted to save in the correct folder

	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Unable to save file").Error(),
		})
		return
	}

	// Create the image URL
	imageURL := fmt.Sprintf("http://localhost:2025/uploads/%s", filename)

	// Save product with image URL
	product, err := ctrl.Service.Create(request, imageURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Error saving product").Error(),
		})
		return
	}

	// Return the created product
	ctx.JSON(http.StatusCreated, model.BaseResponse[any]{
		Data: product,
	})
}

func (ctrl Controller) GetAllProduct(ctx *gin.Context) {
	var (
		request model.RequestGetProduct
	)
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Bind Query").Error(),
		})
		return
	}
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Validate").Error(),
		})
		return
	}
	results, err := ctrl.Service.GetAllProduct(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Find Product").Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,
		model.BaseResponse[model.BaseResponseList[[]model.Product]]{
			Data: model.BaseResponseList[[]model.Product]{
				Count:   len(results),
				Results: results,
			},
		})

}

func (ctrl Controller) GetProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var (
		request model.RequestGetProductById
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Invalid ID").Error(),
		})
		return
	}

	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Validate").Error(),
		})
		return
	}
	result, err := ctrl.Service.GetProductById(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.BaseResponse[any]{
			Message: errors.Wrap(err, "Find item by ID").Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.BaseResponse[any]{
		Data: result,
	})

}

func (ctrl Controller) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// Replace
	if err := ctrl.Service.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Update item status").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[any]{
			Message: "success",
		},
	)
}
