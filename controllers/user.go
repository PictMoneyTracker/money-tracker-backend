package controllers

import (
	"money-tracker/models"
	S "money-tracker/services"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	if validationErr := validate.Struct(user); validationErr != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validationErr.Error(),
			},
		)
	}

	user.Id = primitive.NewObjectID()

	newUser, e := S.CreateUser(&user)

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
			Data:    &fiber.Map{"Id": newUser.Id},
		},
	)
}

func GetUser(c *fiber.Ctx) error {

	userId := c.Params("userId")

	user, e := S.GetUser(userId)

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
			Data:    &fiber.Map{"user": user},
		},
	)
}

func UpdateUser(c *fiber.Ctx) error {

	userId := c.Params("userId")

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    err.Error(),
			},
		)
	}

	if validationErr := validate.Struct(user); validationErr != nil {
		return c.Status(
			http.StatusBadRequest).JSON(
			models.Response{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    validationErr.Error(),
			},
		)
	}

	modifiedCount, e := S.UpdateUser(userId, &user)

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
			Data:    &fiber.Map{"count": modifiedCount},
		},
	)
}
