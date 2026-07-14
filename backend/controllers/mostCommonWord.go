package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func MostCommonWord(c *gin.Context) {
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))
	n := c.DefaultQuery("n", "3")

	numberOfRecords, err := strconv.Atoi(n)

	if err != nil {
		log.Fatal(err)
	}

	data := services.NthCommonString(MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.GreaterEq, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.LessEq, Comparando: toDate}).
		Col("Text").Records(), numberOfRecords)

	c.JSON(200, data)
}
