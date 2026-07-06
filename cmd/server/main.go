package main

import (
	"github.com/arthsalgia/messages-api/controllers"
	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-gonic/gin"
)

func init() {
	services.LoadEnv()
	services.ConnectDb("chat.db")
	controllers.LoadAllMessages()
}

func main() {
	router := gin.Default()

	router.GET("/get-all", controllers.GetAll)
	router.GET("/best-friend", controllers.BestFriend)
	router.GET("/biggest-fan", controllers.BiggestFan)
	router.GET("/celebrity", controllers.Celebrity)
	router.GET("/nth-common", controllers.NthCommon)
	router.GET("/most-common-word", controllers.MostCommonWord)
	router.GET("/most-texted-date", controllers.MostTextedDate)

	router.Run()
}
