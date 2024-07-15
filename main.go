package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"log"
	"github.com/go-playground/validator/v10"
)

func main() {
	app := fiber.New()
	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "admin",
		},
	}))

	
	//Grouping api
	api := app.Group("/api") // /api
    v1 := api.Group("/v1") // /api/v1

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,ffff World!") //SendString is Ctx method
	})

	//Body Parser method => Client send json to server
	type Person struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	v1.Post("/", func(c *fiber.Ctx) error {
		p := new(Person)
		

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name) // john : log can show timestamp
		log.Println(p.Pass) // doe
		str := p.Name + p.Pass
		return c.JSON(str)
	})

	v1.Get("/user/:name", func(c *fiber.Ctx) error {

		str := "hello ==> " + c.Params("name")
		return c.JSON(str)
	})

	v1.Post("/inet", func(c *fiber.Ctx) error {
		a := c.Query("search")
		str := "my search is  " + a
		return c.JSON(str)
	})

	
	//Validation
	v1.Post("/valid", func(c *fiber.Ctx) error {
		//Connect to database
		type User struct {
			Name     string `json:"name" validate:"required,min=3,max=32"`
			IsActive *bool  `json:"isactive" validate:"required"`
			Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
		}
		user := new(User)
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		validate := validator.New()
		errors := validate.Struct(user)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
		}
		return c.JSON(user)
	})
 

	app.Listen(":3000")
}
