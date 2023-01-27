package handlers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"learn-mock/config"
	"learn-mock/models"
	"learn-mock/repository"
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
		ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut())
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
		ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut())
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

func (r ProductHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut())
		defer cancel()
		err := r.repo.Delete(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		return c.JSON(models.GeneralResponse{
			Message: fmt.Sprintf("product with id : %s has been deleted", id),
		})
	}
}

func (r ProductHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var reqBody models.UpdateProductRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		ctx, cancel := context.WithTimeout(context.Background(), config.TimeOut())
		defer cancel()
		result, err := r.repo.Update(ctx, id, &reqBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
				Error: err.Error(),
			})
		}
		return c.JSON(models.GeneralResponse{
			Data: result,
		})
	}
}
