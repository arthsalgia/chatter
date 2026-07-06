package controllers

import (
	"time"

	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func BiggestFan(c *gin.Context) {
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))

	mostCommon, frequency := services.MostCommonString(MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.Greater, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.Less, Comparando: toDate}).
		Filter(dataframe.F{
			Colname:    "IsFromMe",
			Comparator: series.Eq,
			Comparando: 0}).
		Col("ChatID").Records())

	if mostCommon == "" && frequency == 0 {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"biggest_fan":        mostCommon,
		"number_of_messages": frequency,
	})
}
