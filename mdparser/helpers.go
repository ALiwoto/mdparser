// mdparser library Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strconv"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
)

func GetEmpty() WMarkDown {
	return &wotoMarkDown{
		_value: ws.EMPTY,
	}
}

func GetNormal(value string) WMarkDown {
	if ws.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toNormal(value))
}

func toNormal(value string) string {
	return repairValue(value)
}

func GetBold(value string) WMarkDown {
	if ws.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toBold(value))
}

func toBold(value string) string {
	return string(_CHAR_S3) + repairValue(value) +
		string(_CHAR_S3)
}

func GetItalic(value string) WMarkDown {
	if ws.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toItalic(value))
}

func toItalic(value string) string {
	return string(_CHAR_S4) + repairValue(value) + string(_CHAR_S4)
}

func GetMono(value string) WMarkDown {
	if ws.IsEmpty(&value) {
		return GetEmpty()
	}

	return toWotoMD(toMono(value))
}

func toMono(value string) string {
	return string(_CHAR_S16) + repairValue(value) +
		string(_CHAR_S16)
}

func GetHyperLink(text string, url string) WMarkDown {
	if ws.IsEmpty(&text) {
		return GetEmpty()
	}

	return toWotoMD(toHyperLink(text, url))
}

func toHyperLink(text, url string) string {
	fText := repairValue(text)
	fUrl := repairValue(url)
	return string(_CHAR_S7) + fText + string(_CHAR_S8) +
		string(_CHAR_S9) + fUrl + string(_CHAR_S10)
}

// GetUserMention will give you a mentioning style username with the
// specified text.
// WARNING: you don't need to repair text before sending it as first arg,
// this function will check it itself.
func GetUserMention(text string, userID int64) WMarkDown {
	if ws.IsEmpty(&text) {
		return GetEmpty()
	}

	if userID == ws.BaseIndex {
		return GetMono(text)
	}

	return toWotoMD(toUserMention(text, userID))
}

func toUserMention(text string, userID int64) string {
	return string(_CHAR_S7) + repairValue(text) +
		string(_CHAR_S8) +
		string(_CHAR_S9) + tG_USER_ID +
		strconv.FormatInt(userID, baseTen) +
		string(_CHAR_S10)
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
	if ws.IsEmpty(&value) {
		return nil
	}

	return &wotoMarkDown{
		_value: value,
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
