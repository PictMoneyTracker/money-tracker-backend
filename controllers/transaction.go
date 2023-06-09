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

func GetTransactions(c *fiber.Ctx) error {

	userId := c.Params("userId")

	transactions, err := S.GetTransactions(userId)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    transactions,
		},
	)
}

func GetTransaction(c *fiber.Ctx) error {

	userId := c.Params("userId")

	transactionId := c.Params("transactionId")

	transaction, err := S.GetTransaction(userId, transactionId)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    transaction,
		},
	)
}

func DeleteTransaction(c *fiber.Ctx) error {

	userId := c.Params("userId")
	transactionId := c.Params("transactionId")

	delCnt, err := S.DeleteTransaction(userId, transactionId)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"deletedCount": delCnt},
		},
	)
}

func UpdateTransaction(c *fiber.Ctx) error {

	userId := c.Params("userId")
	transactionId := c.Params("transactionId")

	var transaction models.Transaction

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

	updCnt, err := S.UpdateTransaction(userId, transactionId, &transaction)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"Transaction": updCnt},
		},
	)
}

func CalculateTotalCategory(c *fiber.Ctx) error {

	userId := c.Params("userId")
	category := c.Params("category")

	total, err := S.CalculateTotalCategory(userId, category)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"Total": total},
		},
	)
}

func CalculateTotalSpendFrom(c *fiber.Ctx) error {

	userId := c.Params("userId")
	spendFrom := c.Params("spendFrom")

	total, err := S.CalculateTotalSpendFrom(userId, spendFrom)

	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			models.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	return c.Status(
		http.StatusOK).JSON(
		models.Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"Total": total},
		},
	)
}
