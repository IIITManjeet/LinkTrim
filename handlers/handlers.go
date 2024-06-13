package handlers

import (
	"github.com/IIITManjeet/shortify/database"
	"github.com/IIITManjeet/shortify/models"
	"github.com/IIITManjeet/shortify/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

func InitRouter(app *fiber.App) {
	app.Use(cors.New())
	short := app.Group("/api/v1")

	short.Get("/shorts", GetAllURLs)
	short.Get("/shorts/:id", GetURL)
	short.Post("/shorts", CreateURL)
	short.Delete("/shorts/:id", DeleteURL)

	app.Get("/r/:redirect", Redirect)
}

func GetAllURLs(c *fiber.Ctx) error {
	URLs, err := database.GetAllURLs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cannot get all urls: " + err.Error(),
		})
	}
	return c.JSON(URLs)
}

func GetURL(c *fiber.Ctx) error {
	id := c.Params("id")
	URL, err := database.GetURL(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting url from database: " + err.Error(),
		})
	}
	return c.JSON(URL)
}

func CreateURL(c *fiber.Ctx) error {
	var URL models.URL
	if err := c.BodyParser(&URL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body: " + err.Error(),
		})
	}
	URL.ID = uuid.NewString()
	URL.Random = URL.ShortURL == ""
	if URL.Random {
		URL.ShortURL = utils.GenerateWord()
	}
	if err := utils.ValidateURL(URL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error validating URL: " + err.Error(),
		})
	}
	if err := database.CreateURL(URL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error creating record for URL: " + err.Error(),
		})
	}
	return c.JSON(URL)
}

func DeleteURL(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DeleteURL(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting url by id: " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func Redirect(c *fiber.Ctx) error {
	reirectURL := c.Params("redirect")
	URL, err := database.FindURLbyShortURL(reirectURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error finding url by short url: " + err.Error(),
		})
	}
	URL.Clicked++
	if err := database.UpdateURL(URL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error updating url clicks: " + err.Error(),
		})
	}
	return c.Redirect(URL.TargetURL, fiber.StatusTemporaryRedirect)
}
