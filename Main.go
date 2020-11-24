package main

import (
	Controllers "api/Controllers"
	env "api/Env"
	Migration "api/Migration/Migrate"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test gorm")
	switch os.Getenv("ENV") {
	case "prod":
		fmt.Println("Prod")
		env.SetProdEnv()
	default:
		fmt.Println("dev")
		env.SetDevEnv()
	}
	Migration.MigrationProject()
	app := gin.Default()
	app.GET("test", Controllers.Welcome)
	log.Fatal(app.Run(":8080"))
}
