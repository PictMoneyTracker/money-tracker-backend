package controllers

import (
	"money-tracker/models"
	S "money-tracker/services"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddStock(c *fiber.Ctx) error {

	userId := c.Params("userId")

	var stock models.Stock

	if err := c.BodyParser(&stock); err != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	if validationErr := validate.Struct(stock); validationErr != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validationErr.Error(),
			},
		)
	}

	newStock, e := S.AddStock(userId, &stock)

	if e != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    e.Error(),
			},
		)
	}

	return c.Status(
		http.StatusCreated).JSON(
		models.Response{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"Id": newStock.Id},
		},
	)

}

func GetStocks(c *fiber.Ctx) error {
	
	userId := c.Params("userId")

	stocks, e := S.GetStocks(userId)

	if e != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    e.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    stocks,
		},
	)

}

func DeleteStock(c *fiber.Ctx) error {
	
	userId := c.Params("userId")
	stockIdStr := c.Params("stockId")
	stockId, err := strconv.ParseInt(stockIdStr, 10, 32)

	if err != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	deletedStockId, e := S.DeleteStock(userId, int32(stockId))

	if e != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    e.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"Id": deletedStockId},
		},
	)

}
