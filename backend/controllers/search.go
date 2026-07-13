package controllers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func Search(c *gin.Context) {
	word := c.Query("word")
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))

	containsWord := func(el series.Element) bool {
		if val, ok := el.Val().(string); ok {
			return strings.Contains(val, word)
		}
		return false
	}

	countOther := MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.Greater, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.Less, Comparando: toDate}).
		Filter(dataframe.F{
			Colname:    "IsFromMe",
			Comparator: series.Eq,
			Comparando: 0}).
		Filter(dataframe.F{
			Colname:    "Text",
			Comparator: series.CompFunc,
			Comparando: containsWord}).
		Nrow()

	countMe := MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.Greater, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.Less, Comparando: toDate}).
		Filter(dataframe.F{
			Colname:    "IsFromMe",
			Comparator: series.Eq,
			Comparando: 1}).
		Filter(dataframe.F{
			Colname:    "Text",
			Comparator: series.CompFunc,
			Comparando: containsWord}).
		Nrow()

	c.JSON(200, gin.H{
		"total":       countMe + countOther,
		"from_me":     countMe,
		"from_others": countOther,
	})
}
