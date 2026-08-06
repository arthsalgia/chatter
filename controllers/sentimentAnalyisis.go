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

	getAverageCompound := func(texts []string) (int, int, int) {

		var pos, neg, neu int

		for _, text := range texts {
			if text != "" {
				parsedText := sentitext.Parse(text, lexicon.DefaultLexicon)
				scores := sentitext.PolarityScore(parsedText)
				if scores.Compound >= 0.05 {
					pos++
				} else if scores.Compound <= -0.05 {
					neg++
				} else {
					neu++
				}
			}
		}

		return pos, neg, neu
	}

	posMe, negMe, neuMe := getAverageCompound(textsMe)
	posOther, negOther, neuOther := getAverageCompound(textsOther)

	sentimentMe := float64(posMe-negMe) / float64(posMe+negMe+neuMe)
	sentimentOther := float64(posOther-negOther) / float64(posOther+negOther+neuOther)

	totalPos := posMe + posOther
	totalNeg := negMe + negOther
	totalNeu := neuMe + neuOther

	totalSentiment := 0.0
	totalCount := totalPos + totalNeg + totalNeu
	if totalCount > 0 {
		totalSentiment = float64(totalPos-totalNeg) / float64(totalCount)
	}

	c.JSON(200, gin.H{
		"totalSentiment": fmt.Sprintf("%.2f", totalSentiment),
		"sentimentMe":    fmt.Sprintf("%.2f", sentimentMe),
		"sentimentOther": fmt.Sprintf("%.2f", sentimentOther),
		"posMe":          posMe,
		"negMe":          negMe,
		"neuMe":          neuMe,
		"posOther":       posOther,
		"negOther":       negOther,
		"neuOther":       neuOther,
	})
}
