package repository

import (
	"net/http"

	"github.com/RianIhsan/fullstack-goreact/database/migrations"
	"github.com/RianIhsan/fullstack-goreact/database/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(user models.User) []*ErrorResponse {
	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) GetUsers(c *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Users{})
	pg := paginate.New(&paginate.Config{
		DefaultSize:        20,
		CustomParamEnabled: true,
	})
	page := pg.With(model).Request(c.Request()).Response(&[]migrations.Users{})
	c.Status(http.StatusOK).JSON(fiber.Map{
		"date": page,
	})
	return nil
}

func (r *Repository) CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request Failed",
		})
		return err
	}
	errors := ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed create new user",
			"data":    err,
		})
	}
	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "succes",
		"message": "User has been added",
		"data":    user,
	})
	return nil

}

func (r *Repository) UpdateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Request Failed",
		})
		return err
	}
	errors := ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil
	}
	if db.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not get user profile",
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "succes",
		"message": "User successfully update",
		"data":    user,
	})
	return nil
}
func (r *Repository) DeleteUser(c *fiber.Ctx) error {
	userModel := migrations.Users{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(userModel, id)
	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"messgae": "Could not delete",
		})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "succes",
		"messgae": "This Id successfully delete",
		"data":    userModel,
	})
}

func (r *Repository) GetUserByID(c *fiber.Ctx) error {
	userModel := migrations.Users{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID cannot be empty",
		})
		return nil
	}
	if err := r.DB.Where("id = ?", id).First(&userModel).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get User",
		})
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "succes",
		"messgae": "Successfully Get User",
		"data":    userModel,
	})
}
