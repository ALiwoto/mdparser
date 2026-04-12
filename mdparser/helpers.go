// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strconv"
	"strings"
)

// AddSecret adds a new secret variable to the mdparser *globally*.
// from now on, the library itself will automatically censor all of the
// "value"s with their name.
// an example of usage would be:
//  mdparser.AddSecret(bot.Token, "$TOKEN")
func AddSecret(value, name string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	index := getSecretIndexByValue(value)
	if index != -1 {
		secrets[index].name = name
		return
	}

	secrets = append(secrets, secretContainer{
		value: value,
		name:  name,
	})
}

func RemoveSecretByValue(value string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	index := getSecretIndexByValue(value)
	if index != -1 {
		secrets = append(secrets[:index], secrets[index+1:]...)
	}
}

func RemoveSecretByName(name string) {
	secretMu.Lock()
	defer secretMu.Unlock()

	var newSecrets []secretContainer
	for _, current := range secrets {
		if current.name != name {
			newSecrets = append(newSecrets, current)
		}
	}

	secrets = newSecrets
}

func GetSecretIndexByValue(value string) int {
	secretMu.RLock()
	defer secretMu.RUnlock()

	return getSecretIndexByValue(value)
}

func getSecretIndexByValue(value string) int {
	for index, current := range secrets {
		if current.value == value {
			return index
		}
	}

	return -1
}

func SecretValueExists(value string) bool {
	secretMu.RLock()
	defer secretMu.RUnlock()

	return getSecretIndexByValue(value) != -1
}

func GetEmpty() WMarkDown {
	return &wotoMarkDown{}
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

	if userID == baseIndex {
		return GetMono(text)
	}

	return toWotoMD(toUserMention(text, userID))
}

func toUserMention(text string, id int64) string {
	return "[" + repairValue(text) + "]" + "(" + telegramUserIDPrefix + strconv.FormatInt(id, 10) + ")"
}

func IsSpecial(r rune) bool {
	return strings.ContainsRune(specialChars, r)
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
	if value == "" {
		return ""
	}

	value = checkSecrets(value)

	var builder strings.Builder
	builder.Grow(len(value) * 2)

	for _, current := range value {
		if IsSpecial(current) {
			builder.WriteRune(markdownEscapeChar)
		}
		builder.WriteRune(current)
	}

	return builder.String()
}

func checkSecrets(value string) string {
	secretMu.RLock()
	defer secretMu.RUnlock()

	for _, current := range secrets {
		value = strings.ReplaceAll(value, current.value, current.name)
	}

	return value
}
