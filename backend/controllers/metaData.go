package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func MetaData(c *gin.Context) {

	colMe := MessagesDF.Filter(dataframe.F{
		Colname:    "IsFromMe",
		Comparator: series.Eq,
		Comparando: 1}).
		Col("Text")

	var totalLenMe, maxLenMe int

	for i := 0; i < colMe.Len(); i++ {
		s := colMe.Elem(i).String()
		l := len(s)

		totalLenMe += l
		if l > maxLenMe {
			maxLenMe = l
		}
	}

	avgLenMe := float64(totalLenMe) / float64(colMe.Len())

	colOther := MessagesDF.Filter(dataframe.F{
		Colname:    "IsFromMe",
		Comparator: series.Eq,
		Comparando: 0}).
		Col("Text")

	var totalLenOther, maxLenOther int

	for i := 0; i < colOther.Len(); i++ {
		s := colOther.Elem(i).String()
		l := len(s)

		totalLenOther += l
		if l > maxLenOther {
			maxLenOther = l
		}
	}

	avgLenOther := float64(totalLenOther) / float64(colOther.Len())

	totalMessagesMe := MessagesDF.Filter(dataframe.F{
		Colname:    "IsFromMe",
		Comparator: series.Eq,
		Comparando: 1}).Nrow()
	totalMessagesOther := MessagesDF.Filter(dataframe.F{
		Colname:    "IsFromMe",
		Comparator: series.Eq,
		Comparando: 0}).Nrow()

	c.JSON(200, gin.H{
		"total_messages":        totalMessagesMe + totalMessagesOther,
		"total_messages_me":     totalMessagesMe,
		"total_messages_others": totalMessagesOther,
		"avg_len_me":            fmt.Sprintf("%.1f", avgLenMe),
		"max_len_me":            maxLenMe,
		"avg_len_other":         fmt.Sprintf("%.1f", avgLenOther),
		"max_len_other":         maxLenOther,
	})
}
