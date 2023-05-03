package controllers

import (
	"money-tracker/models"
	S "money-tracker/services"
	"time"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTransaction(c *fiber.Ctx) error {

	var transaction models.Transaction
	userId := c.Params("userId")

	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	if validationErr := validate.Struct(transaction); validationErr != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validationErr.Error(),
			},
		)
	}

	transaction.Id = primitive.NewObjectID()
	trxUserId, err := primitive.ObjectIDFromHex(userId)
	transaction.CreatedAt = time.Now()

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

	transaction.UserId = trxUserId

	newTransaction, e := S.AddTransaction(userId, &transaction)

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
			Data:    &fiber.Map{"Id": newTransaction.Id},
		},
	)

}
