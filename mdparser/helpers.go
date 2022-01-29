// mdparser library Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func GetEmpty() WMarkDown {
	return &wotoMarkDown{
		_value: ws.EMPTY,
	}
}

func GetNormal(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toNormal(text))
}

func toNormal(value string) string {
	if value == "" {
		return ""
	}

	return repairValue(value)
}

func GetBold(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toBold(text))
}

func toBold(value string) string {
	if value == "" {
		return ""
	}

	return "*" + repairValue(value) + "*"
}

func GetItalic(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toItalic(text))
}

func toItalic(value string) string {
	return "_" + repairValue(value) + "_"
}

func GetMono(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toMono(text))
}

func toMono(value string) string {
	if value == "" {
		return ""
	}

	return "`" + repairValue(value) + "`"
}

func GetSpoiler(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toSpoiler(text))
}

func toSpoiler(value string) string {
	if value == "" {
		return ""
	}

	return "||" + repairValue(value) + "||"
}

func GetUnderline(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toUnderline(text))
}

func toUnderline(value string) string {
	if value == "" {
		return ""
	}

	return "__" + repairValue(value) + "__"
}

func GetStrike(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toStrike(text))
}

func toStrike(value string) string {
	if value == "" {
		return ""
	}

	return "~" + repairValue(value) + "~"
}

func GetHyperLink(text string, url string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return toWotoMD(toHyperLink(text, url))
}

func toHyperLink(text, url string) string {
	fText := repairValue(text)
	fUrl := repairValue(url)
	return "[" + fText + "]" + "(" + fUrl + ")"
}

// GetUserMention will give you a mentioning style username with the
// specified text.
// WARNING: you don't need to repair text before sending it as first arg,
// this function will check it itself.
func GetUserMention(text string, userID int64) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	if userID == ws.BaseIndex {
		return GetMono(text)
	}

	return toWotoMD(toUserMention(text, userID))
}

func toUserMention(text string, id int64) string {
	return "[" + repairValue(text) + "]" + "(" + _TG_USER_ID + ws.ToBase10(id) + ")"
}

func IsSpecial(r rune) bool {
	for _, current := range _sChars {
		if r == current {
			return true
		}
	}
	return false
}

func ParseFromMessage(message *gotgbot.Message) WMarkDown {
	w := GetEmpty()
	if message == nil || len(message.Entities) == 0 {
		return w
	}

	for _, current := range message.Entities {
		switch current.Type {
		// Type of the entity, can be
		// "mention" (@username),
		// "hashtag" (#hashtag),
		// "cashtag" ($USD),
		// "bot_command" (/start@jobs_bot),
		// "url" (https://telegram.org),
		// "email" (do-not-reply@telegram.org),
		// "phone_number" (+1-212-555-0123),
		// "bold" (bold text),
		// "italic" (italic text),
		// "underline" (underlined text),
		// "strikethrough" (strikethrough text),
		// "spoiler" (spoiler message),
		// "code" (monowidth string),
		// "pre" (monowidth block),
		// "text_link" (for clickable text URLs),
		// "text_mention" (for users without usernames)
		case "mention", "hashtag", "cashtag", "bot_command", "url", "email", "phone_number":
			w.Normal(message.Text[current.Offset : current.Offset+current.Length])
		case "bold":
			w.Bold(message.Text[current.Offset : current.Offset+current.Length])
		case "italic":
			w.Italic(message.Text[current.Offset : current.Offset+current.Length])
		case "code", "pre":
			w.Mono(message.Text[current.Offset : current.Offset+current.Length])
		case "text_link":
			w.HyperLink(message.Text[current.Offset:current.Offset+current.Length], current.Url)
		case "text_mention":
			w.Mention(message.Text[current.Offset:current.Offset+current.Length], current.User.Id)
		case "spoiler":
			w.Spoiler(message.Text[current.Offset : current.Offset+current.Length])
		case "strikethrough":
			w.Strike(message.Text[current.Offset : current.Offset+current.Length])
		case "underline":
			w.Underline(message.Text[current.Offset : current.Offset+current.Length])
		}
	}

	return w
}

func toWotoMD(text string) WMarkDown {
	if text == "" {
		return nil
	}

	return &wotoMarkDown{
		_value: text,
	}
}

func repairValue(value string) string {
	finally := ws.EMPTY
	for _, current := range value {
		if IsSpecial(current) {
			finally += string(_CHAR_S1)
		}
		finally += string(current)
	}

	return finally
}
