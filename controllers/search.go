package controllers

import (
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func Search(c *gin.Context) {
	partial := c.Query("partial")
	word := c.Query("word")
	fromDate := c.DefaultQuery("from", "2000-01-01")
	toDate := c.DefaultQuery("to", time.Now().Format("2006-01-02"))

	var containsWord func(series.Element) bool

	if partial == "true" {
		containsWord = func(el series.Element) bool {
			if val, ok := el.Val().(string); ok {
				return strings.Contains(val, word)
			}
			return false
		}
	} else {
		containsWord = func(el series.Element) bool {
			if val, ok := el.Val().(string); ok {
				pattern := `\b` + regexp.QuoteMeta(word) + `\b`
				re := regexp.MustCompile(pattern)
				return re.MatchString(val)
			}
			return false
		}
	}

	countOther := MessagesDF.
		Filter(dataframe.F{Colname: "Date", Comparator: series.GreaterEq, Comparando: fromDate}).
		Filter(dataframe.F{Colname: "Date", Comparator: series.LessEq, Comparando: toDate}).
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
