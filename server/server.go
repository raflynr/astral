package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/raflynr/astral/config"
	"github.com/raflynr/astral/db"
)

type Server struct {
	app  *fiber.App
	conf config.AppConfig
}

func NewServer() *Server {
	app := fiber.New()
	conf := config.NewConfig()

	return &Server{
		app:  app,
		conf: conf,
	}
}

func (s *Server) Run() {
	db, err := db.NewDB(s.conf)
	if err != nil {
		log.Fatalf("error when connecting to database, %s", err)
	}

	s.app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type",
		AllowCredentials: false,
	}))

	NewRoute(db, validator.New(), s.app)

	s.GracefulShutdown(s.conf.Fiber.Port)
}

func (s *Server) GracefulShutdown(port string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.app.Listen(":" + port); err != nil {
			log.Fatalf("error when listening to :%s, %s", port, err)
		}
	}()

	log.Printf("server is running on :%s", port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}
