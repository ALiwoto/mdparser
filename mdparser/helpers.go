// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strconv"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
)

func GetEmpty() WMarkDown {
	return &wotoMarkDown{
		_value: strongStringGo.EMPTY,
	}
}

func GetNormal(value string) WMarkDown {
	if strongStringGo.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toNormal(value))
}

func toNormal(value string) string {
	return repairValue(value)
}

func GetBold(value string) WMarkDown {
	if strongStringGo.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toBold(value))
}

func toBold(value string) string {
	return string(CHAR_S3) + repairValue(value) +
		string(CHAR_S3)
}

func GetItalic(value string) WMarkDown {
	if strongStringGo.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toItalic(value))
}

func toItalic(value string) string {
	return string(CHAR_S4) + repairValue(value) + string(CHAR_S4)
}

func GetMono(value string) WMarkDown {
	if strongStringGo.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toMono(value))
}

func toMono(value string) string {
	return string(CHAR_S16) + repairValue(value) +
		string(CHAR_S16)
}

func GetHyperLink(text string, url string) WMarkDown {
	if strongStringGo.IsEmpty(&text) {
		return GetEmpty()
	}

	return toWotoMD(toHyperLink(text, url))
}

func toHyperLink(text, url string) string {
	fText := repairValue(text)
	fUrl := repairValue(url)
	return string(CHAR_S7) + fText + string(CHAR_S8) +
		string(CHAR_S9) + fUrl + string(CHAR_S10)
}

// GetUserMention will give you a mentioning style username with the
// specified text.
// WARNING: you don't need to repair text before sending it as first arg,
// this function will check it itself.
func GetUserMention(text string, userID int64) WMarkDown {
	if strongStringGo.IsEmpty(&text) {
		return GetEmpty()
	}

	if userID == strongStringGo.BaseIndex {
		return GetMono(text)
	}

	return toWotoMD(toUserMention(text, userID))
}

func toUserMention(text string, userID int64) string {
	return string(CHAR_S7) + repairValue(text) +
		string(CHAR_S8) +
		string(CHAR_S9) + TG_USER_ID +
		strconv.FormatInt(userID, BaseTen) +
		string(CHAR_S10)
}

func IsSpecial(r rune) bool {
	for _, current := range _sChars {
		if r == current {
			return true
		}
	}
	return false
}

func toWotoMD(value string) WMarkDown {
	if strongStringGo.IsEmpty(&value) {
		return nil
	}

	return &wotoMarkDown{
		_value: value,
	}
}

func repairValue(value string) string {
	finally := strongStringGo.EMPTY
	escape := false
	lasEscape := false
	escapeCount := strongStringGo.BaseIndex
	for i, current := range value {
		escape = (current == CHAR_S1)
		if IsSpecial(current) {
			if escape {
				escapeCount++
			} else {
				escapeCount = strongStringGo.BaseIndex
			}
			if i != strongStringGo.BaseIndex {
				if !lasEscape {
					finally += string(CHAR_S1) + string(current)
				} else {
					finally += string(current)
				}
			} else {
				finally += string(CHAR_S1) + string(current)
			}
		} else {
			if escapeCount != strongStringGo.BaseIndex {
				tmpR := escapeCount % BaseTwoIndex
				if tmpR != strongStringGo.BaseIndex {
					finally += string(CHAR_S1) + string(current)
				} else {
					finally += string(current)
				}
			} else {
				finally += string(current)
			}
		}
		lasEscape = escape
	}
	return finally
}
