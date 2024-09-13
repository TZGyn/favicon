package main

import (
	"time"
	"tzgyn/favicon/providers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/patrickmn/go-cache"
)

func main() {
	cache_client := cache.New(1*time.Hour, 5*time.Minute)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:domain", func(c *fiber.Ctx) error {
		domain := c.Params("domain")

		cached_result, found := cache_client.Get(domain)

		if found {
			c.Response().Header.Set("Content-Type", cached_result.(providers.Avatar).Data_type)
			return c.SendString(cached_result.(providers.Avatar).Data)
		}

		avatar, err := providers.Domain(domain)
		if err != nil {
			return c.SendString("Invalid")
		}

		cache_client.Set(domain, avatar, cache.DefaultExpiration)

		c.Response().Header.Set("Content-Type", avatar.Data_type)

		return c.SendString(avatar.Data)
	})

	app.Get("/youtube/@:channel", func(c *fiber.Ctx) error {
		channel := c.Params("channel")

		cached_result, found := cache_client.Get("yt-" + channel)

		if found {
			c.Response().Header.Set("Content-Type", cached_result.(providers.Avatar).Data_type)
			return c.SendString(cached_result.(providers.Avatar).Data)
		}

		avatar, err := providers.YoutubeChannel(channel)

		if err != nil {
			return c.SendString("Invalid")
		}

		c.Response().Header.Set("Content-Type", avatar.Data_type)

		cache_client.Set("yt-"+channel, avatar, cache.DefaultExpiration)

		return c.SendString(avatar.Data)
	})

	app.Get("/twitch/:channel", func(c *fiber.Ctx) error {
		channel := c.Params("channel")

		cached_result, found := cache_client.Get("twitch-" + channel)

		if found {
			c.Response().Header.Set("Content-Type", cached_result.(providers.Avatar).Data_type)
			return c.SendString(cached_result.(providers.Avatar).Data)
		}

		avatar, err := providers.TwitchChannel(channel)

		if err != nil {
			return c.SendString("Invalid")
		}

		c.Response().Header.Set("Content-Type", avatar.Data_type)

		cache_client.Set("twitch-"+channel, avatar, cache.DefaultExpiration)

		return c.SendString(avatar.Data)
	})

	app.Get("/github/:account", func(c *fiber.Ctx) error {
		account := c.Params("account")
		cached_result, found := cache_client.Get("gh-" + account)

		if found {
			c.Response().Header.Set("Content-Type", cached_result.(providers.Avatar).Data_type)
			return c.SendString(cached_result.(providers.Avatar).Data)
		}

		avatar, err := providers.GithubAccount(account)

		if err != nil {
			return c.SendString("Invalid")
		}

		c.Response().Header.Set("Content-Type", avatar.Data_type)

		cache_client.Set("gh-"+account, avatar, cache.DefaultExpiration)

		return c.SendString(avatar.Data)
	})

	app.Listen(":3000")
}
