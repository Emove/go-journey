package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()
	service := StockService{}
	app.Post("/stock", func(ctx *fiber.Ctx) error {
		stock, err := service.Stock("product01", 10*time.Second, 3, nil)
		if err != nil {
			return err
		}
		data := map[string]int64{
			"remain": stock,
		}

		marshal, _ := json.Marshal(Response{
			Code:    1,
			Message: "success",
			Data:    data,
		})
		_, _ = ctx.Write(marshal)
		return nil
	})
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
