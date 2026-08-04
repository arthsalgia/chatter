package controllers

import "github.com/gin-gonic/gin"

func GetAllChats(c *gin.Context) {
	records := MessagesDF.Col("ChatID").Records()

	seen := make(map[string]bool)
	var uniqueRecords []string

	for _, val := range records {
		if !seen[val] {
			seen[val] = true
			uniqueRecords = append(uniqueRecords, val)
		}
	}

	c.JSON(200, gin.H{
		"chats": uniqueRecords,
	})
}
