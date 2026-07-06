package services

import (
	"bytes"
	"regexp"
	"strings"
	"time"

	"github.com/arthsalgia/messages-api/models"
	"github.com/go-gota/gota/dataframe"
)

var MessagesDF dataframe.DataFrame
var blobRegex = regexp.MustCompile(`[\x20-\x7E\x{2010}-\x{201F}\x{2026}]{2,}`)

func LoadAllMessages() {
	var messages []models.MessagesAll
	DB.Order("ROWID desc").Find(&messages)

	if len(messages) == 0 {
		return
	}

	type FlatMessage struct {
		ID       int    `json:"id"`
		ROWID    int    `json:"row_id"`
		Text     string `json:"text"`
		ChatID   string `json:"chat_id"`
		IsFromMe int    `json:"is_from_me"`
		Date     string `json:"date"`
	}

	flat := make([]FlatMessage, 0, len(messages))

	macEpoch := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

	for i, m := range messages {
		var decoded string
		var text string
		var c_id string
		var rawTimestamp int64
		rawTimestamp = int64(m.Date)
		if len(m.AttributedBody) > 0 {
			matches := blobRegex.FindAll(m.AttributedBody, -1)
			decoded = string(bytes.Join(matches, []byte(" ")))
			NSString := strings.Index(decoded, "NSString")
			remaining := decoded[NSString+8:]
			iINSDictionary := strings.Index(remaining, " iI NSDictionary")
			text = remaining[:iINSDictionary]
			text = strings.TrimSpace(text)
		}

		if len(m.Ck_chat_id) >= 11 {
			c_id = m.Ck_chat_id[11:]
		} else {
			c_id = m.Ck_chat_id
		}

		text = ParseMessage(text)

		if text == "" || text == "an image" {
			continue
		}

		var messageTime time.Time
		if rawTimestamp > 9999999999 {
			messageTime = macEpoch.Add(time.Duration(rawTimestamp) * time.Nanosecond)
		} else {
			messageTime = macEpoch.Add(time.Duration(rawTimestamp) * time.Second)
		}

		flat = append(flat, FlatMessage{
			ID:       i,
			ROWID:    m.ROWID,
			Text:     text,
			ChatID:   c_id,
			IsFromMe: m.Is_from_me,
			Date:     messageTime.Local().Format("2006-01-02"),
		})
	}

	MessagesDF = dataframe.LoadStructs(flat)

}
