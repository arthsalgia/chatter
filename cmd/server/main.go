package main

import (
	"flag"
	"fmt"

	"github.com/arthsalgia/messages-api/controllers"
	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	dbPathPtr := flag.String("db", "chat.db", "path to the iMessage SQLite chat file")
	portPtr := flag.String("port", ":8000", "The port for the API server to run on")
	flag.Parse()

	fmt.Println("Starting server...")

	services.ConnectDb(*dbPathPtr)
	controllers.LoadAllMessages()

	router := gin.Default()

	router.GET("/get-all", controllers.GetAll)
	router.GET("/best-friend", controllers.BestFriend)
	router.GET("/biggest-fan", controllers.BiggestFan)
	router.GET("/celebrity", controllers.Celebrity)
	router.GET("/nth-common", controllers.NthCommon)
	router.GET("/most-common-word", controllers.MostCommonWord)
	router.GET("/most-texted-date", controllers.MostTextedDate)

	router.Run(*portPtr)
}
