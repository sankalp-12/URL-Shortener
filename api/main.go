package main


import(
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/Daksh-Axel/shorten-url-fiber-redis/routes"
	"github.com/ansrivas/fiberprometheus/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/:url",routes.ResolveURL)
	app.Post("/api/v1",routes.ShortenURL)
}

func main(){
	err := godotenv.Load()
	if err!=nil{
		fmt.Println(err)
	}
	app := fiber.New()
	
	// adding middleware: Exporter for prometheus
	prometheus := fiberprometheus.New("my-service-name")
  	prometheus.RegisterAt(app, "/metrics")
 	app.Use(prometheus.Middleware)
 	
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
