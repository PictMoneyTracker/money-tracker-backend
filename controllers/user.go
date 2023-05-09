package controllers

import (
	"money-tracker/models"
	S "money-tracker/services"
	U "money-tracker/utils"

	"net/http"

	"github.com/shareed2k/goth_fiber"

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

	signedToken, err := U.GetJwtToken(&user)

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

	c.Set("Authorization", signedToken)

	// check if user already exists
	if existingUser, err := S.GetUserByEmail(user.Email); existingUser != nil {
		if err == nil {
			return c.Status(
				http.StatusOK).JSON(
				models.Response{
					Status:  http.StatusOK,
					Message: "success",
					Data:    &fiber.Map{"id": existingUser.Id},
				},
			)
		}
	}

	newUser, err := S.CreateUser(&user)

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
		http.StatusCreated).JSON(
		models.Response{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"id": newUser.Id},
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

// google auth callback for web platform
func HandleAuth(c *fiber.Ctx) error {

	user, err := goth_fiber.CompleteUserAuth(c)

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

	// Create a new user with the required properties
	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Email:    user.Email,
		Name:     user.Name,
		PhotoUrl: user.AvatarURL,
	}

	signedToken, err := U.GetJwtToken(&newUser)

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

	c.Set("Authorization", signedToken)

	// check if user already exists
	if existingUser, err := S.GetUserByEmail(user.Email); existingUser != nil {

		if err == nil {

			return c.Status(
				http.StatusOK).JSON(
				models.Response{
					Status:  http.StatusOK,
					Message: "success",
					Data:    &fiber.Map{"user": existingUser},
				},
			)
		}
	}

	// Insert the user into the database
	newUserDb, err := S.CreateUser(&newUser)

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
			Data:    &fiber.Map{"newUser": newUserDb},
		},
	)

}

func HandleLogin(ctx *fiber.Ctx) error {
	err := goth_fiber.BeginAuthHandler(ctx)

	if err != nil {
		return err
	}
	return nil
}
