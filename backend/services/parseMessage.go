package services

func ParseMessage(text string) string {
	end := len(text)
	if end == 0 {
		return ""
	}

	start := 0

	if text[0] == '+' && end >= 2 {
		c := text[1]
		if c >= 32 && c <= 126 {
			start = 2
		}
	}

	for start < end && (text[start] == ' ' || text[start] == '\t') {
		start++
	}

	remainingLen := end - start
	if remainingLen >= 12 && text[start:start+10] == "Laughed at" {
		start += 10
	} else if remainingLen >= 11 && text[start:start+10] == "Emphasized" {
		start += 10
	} else if remainingLen >= 11 && text[start:start+10] == "Questioned" {
		start += 10
	} else if remainingLen >= 9 && text[start:start+7] == "Disliked" {
		start += 7
	} else if remainingLen >= 7 && (text[start:start+5] == "Loved" || text[start:start+5] == "Liked") {
		start += 5
	} else if remainingLen >= 14 && text[start:start+12] == "Reacted   to" {
		start += 12
	}

	for {
		progress := false

		for start < end && (text[start] == ' ' || text[start] == '\t') {
			start++
			progress = true
		}

		if start < end && text[start] == '"' {
			start++
			progress = true
		}

		if end-start >= 3 && text[start] == 0xE2 && text[start+1] == 0x80 && text[start+2] == 0x9C {
			start += 3
			progress = true
		}

		if end-start >= 3 && text[start] == 0xE2 && text[start+1] == 0x80 && text[start+2] == 0x9D {
			start += 3
			progress = true
		}

		for end > start {
			c := text[end-1]
			if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
				end--
				progress = true
			} else {
				break
			}
		}

		if end > start && text[end-1] == '"' {
			end--
			progress = true
		}

		if end-start >= 3 && text[end-3] == 0xE2 && text[end-2] == 0x80 && text[end-1] == 0x9D {
			end -= 3
			progress = true
		}

		if end-start >= 3 && text[end-3] == 0xE2 && text[end-2] == 0x80 && text[end-1] == 0x9C {
			end -= 3
			progress = true
		}

		if !progress {
			break
		}
	}

	if start >= end {
		return ""
	}

	return text[start:end]
}
