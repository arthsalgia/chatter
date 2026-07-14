package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func GetAll(c *gin.Context) {
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))

	df := MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.GreaterEq, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.LessEq, Comparando: toDate})

	c.JSON(200, df.Maps())
}
