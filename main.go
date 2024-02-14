package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Fancy11111/gotmx/persistence"
	"github.com/Fancy11111/gotmx/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	store, err := persistence.SetupStore()
	if err != nil {
		log.Fatal("Could not setup persistence store")
	}

	app := fiber.New()

	routes.SetupRoutes(app, store)

	go func() {
		log.Print("Starting listening on port 3000\n")
		app.Listen(":3000")
	}()

	c := make(chan os.Signal, 2)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c
	println()
	log.Printf("Got signal %v\n", sig)
	log.Print("Shutting down")
	app.Shutdown()
}
