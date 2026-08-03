package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func longestMessage() (string, string, []string, []string) {
	TextMe := MessagesDF.Filter(dataframe.F{Colname: "IsFromMe", Comparator: series.Eq, Comparando: 1}).Col("Text")

	var longestTextMe string
	maxLenMe := 0

	for i := 0; i < TextMe.Len(); i++ {
		val := TextMe.Elem(i).Val().(string)
		if len(val) > maxLenMe {
			maxLenMe = len(val)
			longestTextMe = val
		}
	}

	sentTo := MessagesDF.Filter(dataframe.F{Colname: "Text", Comparator: series.Eq, Comparando: longestTextMe}).Col("ChatID").Records()

	TextOther := MessagesDF.Filter(dataframe.F{Colname: "IsFromMe", Comparator: series.Eq, Comparando: 0}).Col("Text")

	var longestTextOther string
	maxLenOther := 0

	for i := 0; i < TextOther.Len(); i++ {
		val := TextOther.Elem(i).Val().(string)
		if len(val) > maxLenOther {
			maxLenOther = len(val)
			longestTextOther = val
		}
	}
	sentBy := MessagesDF.Filter(dataframe.F{Colname: "Text", Comparator: series.Eq, Comparando: longestTextOther}).Col("ChatID").Records()

	return longestTextMe, longestTextOther, sentTo, sentBy
}

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

	longestTextMe, longestTextOther, sentTo, sentBy := longestMessage()

	c.JSON(200, gin.H{
		"total_messages":        totalMessagesMe + totalMessagesOther,
		"total_messages_me":     totalMessagesMe,
		"total_messages_others": totalMessagesOther,
		"avg_len_me":            fmt.Sprintf("%.1f", avgLenMe),
		"max_len_me":            maxLenMe,
		"avg_len_other":         fmt.Sprintf("%.1f", avgLenOther),
		"max_len_other":         maxLenOther,
		"longestTextMe":         longestTextMe,
		"longestTextOther":      longestTextOther,
		"sentTo":                sentTo[0],
		"sentBy":                sentBy[0],
	})
}
