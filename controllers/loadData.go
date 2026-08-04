package controllers

import (
	"bytes"
	"regexp"
	"strings"
	"time"

	"github.com/arthsalgia/messages-api/models"
	"github.com/arthsalgia/messages-api/services"
	"github.com/go-gota/gota/dataframe"
)

var MessagesDF dataframe.DataFrame
var blobRegex = regexp.MustCompile(`[\x20-\x7E\x{2010}-\x{201F}\x{2026}]{2,}`)

func LoadAllMessages() bool {
	var messages []models.MessagesAll
	var chat []models.Chat
	services.DB.Order("ROWID desc").Find(&messages)
	services.DB.Order("ROWID desc").Find(&chat)

	var count int64
	err := services.DB.Table("message").Count(&count).Error
	if err != nil {
		return false
	}
	chatErr := services.DB.Table("chat").Count(&count).Error
	if chatErr != nil {
		return false
	}

	type Chat struct {
		ID          int    `json:"id"`
		GroupID     string `json:"group_id"`
		DisplayName string `json:"display_name"`
	}

	flatChat := make([]Chat, 0, len(chat))

	for i, c := range chat {
		if len(c.DisplayName) > 0 {
			flatChat = append(flatChat, Chat{
				ID:          i,
				GroupID:     c.GroupID,
				DisplayName: c.DisplayName,
			})
		}
	}

	groupToDisplay := make(map[string]string, len(flatChat))
	for _, c := range flatChat {
		groupToDisplay[c.GroupID] = c.DisplayName
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

	const imessagePrefix = "iMessage;-;"
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
			if iINSDictionary != -1 {
				text = remaining[:iINSDictionary]
				text = strings.TrimSpace(text)
			} else {
				continue
			}
		}

		if strings.HasPrefix(m.Ck_chat_id, imessagePrefix) {
			c_id = m.Ck_chat_id[len(imessagePrefix):]
		} else {
			c_id = m.Ck_chat_id
		}

		if display, ok := groupToDisplay[c_id]; ok {
			c_id = display
		}

		text = services.ParseMessage(text)

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
			Text:     strings.TrimSpace(text),
			ChatID:   strings.TrimSpace(c_id),
			IsFromMe: m.Is_from_me,
			Date:     messageTime.Local().Format("2006-01-02"),
		})
	}

	MessagesDF = dataframe.LoadStructs(flat)
	return true
}
