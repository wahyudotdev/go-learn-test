package handlers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"learn-mock/models"
	"learn-mock/repository"
	"time"
)

var (
	v = validator.New()
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) ProductHandler {
	return ProductHandler{
		repo: repo,
	}
}

func (r ProductHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody models.CreateProductRequest
		err := c.BodyParser(&reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		err = v.Struct(reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		result, err := r.repo.CreateProduct(ctx, &reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(models.GeneralResponse{
			Message: "success",
			Data:    result,
		})
	}
}

func (r ProductHandler) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqQuery models.GetProductRequest
		err := c.QueryParser(&reqQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error:   err.Error(),
				Message: "error",
			})
		}
		err = v.Struct(&reqQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error:   err.Error(),
				Message: "error",
			})
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		result, err := r.repo.GetAll(ctx, &reqQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error:   err.Error(),
				Message: "error",
			})
		}
		return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
			Data: result,
		})
	}
}
