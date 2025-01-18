package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/samar2k17/url-shortner/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	db := database.CreateClient(0)
	defer db.Close()
	val, err := db.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found on database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}
	idb := database.CreateClient(1)
	defer idb.Close()
	_ = idb.Incr(database.Ctx, "counter")
	return c.Redirect(val, 301)
}
