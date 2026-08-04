package controllers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func SentimentAnalysis(c *gin.Context) {
	Chat := c.Query("chat")
	Chat = strings.TrimSpace(Chat)

	textsMe := MessagesDF.
		Filter(dataframe.F{Colname: "ChatID", Comparator: series.Eq, Comparando: Chat}).
		Filter(dataframe.F{Colname: "IsFromMe", Comparator: series.Eq, Comparando: 1}).Col("Text").Records()

	textsOther := MessagesDF.
		Filter(dataframe.F{Colname: "ChatID", Comparator: series.Eq, Comparando: Chat}).
		Filter(dataframe.F{Colname: "IsFromMe", Comparator: series.Eq, Comparando: 0}).Col("Text").Records()

	getAverageCompound := func(texts []string) float64 {

		totalCompound := 0.0
		validCount := 0

		for _, text := range texts {
			if text != "" {
				parsedText := sentitext.Parse(text, lexicon.DefaultLexicon)
				scores := sentitext.PolarityScore(parsedText)
				totalCompound += scores.Compound
				validCount++
			}
		}

		if validCount == 0 {
			return 0.0
		}
		return totalCompound / float64(validCount)
	}

	avgMe := getAverageCompound(textsMe)
	avgOther := getAverageCompound(textsOther)

	c.JSON(200, gin.H{
		"totalSentiment": fmt.Sprintf("%.2f", (avgMe*100*float64(len(textsMe))+(avgOther*100*float64(len(textsOther))))/(float64(len(textsMe))+float64(len(textsOther)))),
		"sentimentMe":    fmt.Sprintf("%.2f", avgMe*100),
		"sentimentOther": fmt.Sprintf("%.2f", avgOther*100),
	})
}
