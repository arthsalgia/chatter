package controllers

import (
	"time"

	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func BestFriend(c *gin.Context) {
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))

	mostCommon, frequency := services.MostCommonString(
		MessagesDF.
			Filter(dataframe.F{Colname: "Date", Comparator: series.GreaterEq, Comparando: fromDate}).
			Filter(dataframe.F{Colname: "Date", Comparator: series.LessEq, Comparando: toDate}).
			Col("ChatID").Records())

	c.JSON(200, gin.H{
		"best_friend":        mostCommon,
		"number_of_messages": frequency,
	})
}
