// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
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
	return _sChars[r]
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
