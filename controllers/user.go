package controllers

import (
	"money-tracker/configs"
	"money-tracker/models"
	S "money-tracker/services"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			models.Response{Status: http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			models.Response{Status: http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	user.Id = primitive.NewObjectID()

	newUser, e := S.CreateUser(&user, userCollection)

	if e != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			models.Response{Status: http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": e.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(
		models.Response{Status: http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"Id": newUser.Id}})
}
