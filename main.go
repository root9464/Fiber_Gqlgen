package main

import (
	"log"
	"root/database"
	gql "root/graph/out"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {
	_, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
	app := fiber.New()
	
	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))
	app.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("Graphql Playground", "/query"))) //прикольная хуйня
	app.All("/query", adaptor.HTTPHandler(srv))

	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})	
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(app.Listen(":9090"))
}