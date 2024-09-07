package main

import (
	"log"

	database "github.com/Scorpios7/To-do-list/internal"
	"github.com/Scorpios7/To-do-list/internal/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	// 设置配置文件名和类型
	viper.SetConfigName("config")             // 配置文件名 (不要包含扩展名)
	viper.SetConfigType("yaml")               // 配置文件类型
	viper.AddConfigPath("../internal/config") // 配置文件路径

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fail to read config: %v", err)
	}
	app := fiber.New()
	if err := database.InitDatabase(); err != nil {
		log.Fatalf("fail to init database: %v", err)
	}
	routes.InitTodoRoute(app)

	port := viper.GetString("PORT")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(port)
	})
	log.Fatal(app.Listen(":" + port))
}
