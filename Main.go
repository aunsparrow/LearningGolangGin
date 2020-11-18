package main

import (
	Controllers "api/Controllers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("test gorm")
	//Migration.MigrationProject()
	app := gin.Default()
	app.GET("test", Controllers.Welcome)
	log.Fatal(app.Run(":8080"))
}
