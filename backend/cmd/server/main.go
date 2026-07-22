package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/arthsalgia/messages-api/controllers"
	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	dbPathPtr := flag.String("db", "./chat.db", "path to the iMessage SQLite chat file")
	portPtr := flag.String("port", ":8000", "The port for the API server to run on")
	flag.Parse()

	fmt.Println("Starting server...")
	fmt.Println("Reading through", *dbPathPtr, "...")

	services.ConnectDb(*dbPathPtr)
	start := controllers.LoadAllMessages()
	if !start {
		fmt.Println("Empty or invalid database")
		fmt.Println("Stopping server...")
		return
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/get-all", controllers.GetAll)
	router.GET("/best-friend", controllers.BestFriend)
	router.GET("/biggest-fan", controllers.BiggestFan)
	router.GET("/celebrity", controllers.Celebrity)
	router.GET("/nth-common", controllers.NthCommon)
	router.GET("/most-common-word", controllers.MostCommonWord)
	router.GET("/most-texted-date", controllers.MostTextedDate)
	router.GET("/search", controllers.Search)

	router.Run(*portPtr)
}
