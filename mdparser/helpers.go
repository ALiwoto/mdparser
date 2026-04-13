// mdparser library Project
// Copyright (C) 2021-2026 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strconv"
	"strings"
)

func GetEmpty() WMarkDown {
	return &wotoMarkDown{}
}

func GetNormal(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newTextSegment(segmentNormal, text))
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

	return newWotoMD(newTextSegment(segmentBold, text))
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

	return newWotoMD(newTextSegment(segmentItalic, text))
}

func toItalic(value string) string {
	return "_" + repairValue(value) + "_"
}

func GetMono(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newTextSegment(segmentMono, text))
}

func toMono(value string) string {
	if value == "" {
		return ""
	}

	return "`" + repairValue(value) + "`"
}

func GetStyled(text string, styles ...TextStyle) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newStyledSegment(text, styles...))
}

func GetCodeBlock(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newTextSegment(segmentCodeBlock, text))
}

func toCodeBlock(value string) string {
	if value == "" {
		return ""
	}

	return markdownCodeFence + "\n" + repairCodeValue(value) + "\n" + markdownCodeFence
}

func GetCodeBlockLang(lang, text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newCodeBlockLangSegment(lang, text))
}

func toCodeBlockLang(lang, value string) string {
	if value == "" {
		return ""
	}

	lang = normalizeCodeLanguage(lang)
	if lang == "" {
		return toCodeBlock(value)
	}

	return markdownCodeFence + lang + "\n" + repairCodeValue(value) + "\n" + markdownCodeFence
}

func GetSpoiler(text string) WMarkDown {
	if text == "" {
		return GetEmpty()
	}

	return newWotoMD(newTextSegment(segmentSpoiler, text))
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

	return newWotoMD(newTextSegment(segmentUnderline, text))
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

	return newWotoMD(newTextSegment(segmentStrike, text))
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

	return newWotoMD(newHyperLinkSegment(text, url))
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

	return newWotoMD(newMentionSegment(text, userID))
}

func IsSpecial(r rune) bool {
	return strings.ContainsRune(specialChars, r)
}

func toWotoMD(text string) WMarkDown {
	if text == "" {
		return nil
	}

	return newWotoMD(newRawSegment(text))
}

func repairValue(value string) string {
	if value == "" {
		return ""
	}

	return escapeValue(checkSecrets(value))
}

func repairCodeValue(value string) string {
	if value == "" {
		return ""
	}

	return escapeCodeValue(checkSecrets(value))
}

func renderStyledValue(value string, styles TextStyle) string {
	value = escapeValue(value)
	styles = normalizeTextStyles(styles)
	if styles == 0 {
		return value
	}

	if styles&StyleUnderline != 0 && styles&StyleItalic != 0 {
		value = "___" + value + "_\r__"
		styles &^= StyleUnderline | StyleItalic
	} else {
		if styles&StyleItalic != 0 {
			value = "_" + value + "_"
		}
		if styles&StyleUnderline != 0 {
			value = "__" + value + "__"
		}
	}

	if styles&StyleStrike != 0 {
		value = "~" + value + "~"
	}
	if styles&StyleBold != 0 {
		value = "*" + value + "*"
	}
	if styles&StyleSpoiler != 0 {
		value = "||" + value + "||"
	}

	return value
}

func normalizeCodeLanguage(lang string) string {
	lang = strings.TrimSpace(lang)
	if lang == "" {
		return ""
	}

	lang = strings.NewReplacer("\r", "", "\n", "", "\t", "").Replace(lang)

	var builder strings.Builder
	builder.Grow(len(lang))

	for _, current := range lang {
		if current == markdownEscapeChar || current == '`' {
			builder.WriteRune(markdownEscapeChar)
		}
		builder.WriteRune(current)
	}

	return builder.String()
}

func escapeValue(value string) string {
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

func escapeCodeValue(value string) string {
	var builder strings.Builder
	builder.Grow(len(value) * 2)

	for _, current := range value {
		if current == markdownEscapeChar || current == '`' {
			builder.WriteRune(markdownEscapeChar)
		}
		builder.WriteRune(current)
	}

	return builder.String()
}

func newWotoMD(segments ...markdownSegment) WMarkDown {
	return &wotoMarkDown{
		_segments: cloneSegments(segments),
	}
}

func cloneSegments(segments []markdownSegment) []markdownSegment {
	if len(segments) == 0 {
		return nil
	}

	cloned := make([]markdownSegment, len(segments))
	copy(cloned, segments)
	return cloned
}

func newRawSegment(value string) markdownSegment {
	return markdownSegment{
		kind: segmentRaw,
		text: value,
	}
}

func newTextSegment(kind markdownSegmentKind, text string) markdownSegment {
	return markdownSegment{
		kind: kind,
		text: sanitizeValue(text),
	}
}

func newStyledSegment(text string, styles ...TextStyle) markdownSegment {
	normalized := normalizeTextStyles(styles...)
	if normalized == 0 {
		return newTextSegment(segmentNormal, text)
	}

	return markdownSegment{
		kind:   segmentStyled,
		text:   sanitizeValue(text),
		styles: normalized,
	}
}

func newCodeBlockLangSegment(lang, text string) markdownSegment {
	normalized := normalizeCodeLanguage(lang)
	if normalized == "" {
		return newTextSegment(segmentCodeBlock, text)
	}

	return markdownSegment{
		kind: segmentCodeBlockLang,
		text: sanitizeValue(text),
		meta: normalized,
	}
}

func newHyperLinkSegment(text, url string) markdownSegment {
	return markdownSegment{
		kind: segmentHyperLink,
		text: sanitizeValue(text),
		meta: sanitizeValue(url),
	}
}

func newMentionSegment(text string, id int64) markdownSegment {
	return markdownSegment{
		kind: segmentMention,
		text: sanitizeValue(text),
		meta: strconv.FormatInt(id, 10),
	}
}

func sanitizeValue(value string) string {
	if value == "" {
		return ""
	}

	return checkSecrets(value)
}

func normalizeTextStyles(styles ...TextStyle) TextStyle {
	var normalized TextStyle
	for _, current := range styles {
		normalized |= current & supportedTextStyles
	}

	return normalized
}

func renderSegments(segments []markdownSegment) string {
	if len(segments) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, current := range segments {
		appendRenderedSegment(&builder, current)
	}

	return builder.String()
}

func appendRenderedSegment(builder *strings.Builder, segment markdownSegment) {
	switch segment.kind {
	case segmentRaw:
		builder.WriteString(segment.text)
	case segmentNormal:
		builder.WriteString(escapeValue(segment.text))
	case segmentBold:
		builder.WriteByte('*')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteByte('*')
	case segmentItalic:
		builder.WriteByte('_')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteByte('_')
	case segmentMono:
		builder.WriteByte('`')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteByte('`')
	case segmentStyled:
		builder.WriteString(renderStyledValue(segment.text, segment.styles))
	case segmentCodeBlock:
		builder.WriteString(markdownCodeFence)
		builder.WriteByte('\n')
		builder.WriteString(escapeCodeValue(segment.text))
		builder.WriteByte('\n')
		builder.WriteString(markdownCodeFence)
	case segmentCodeBlockLang:
		builder.WriteString(markdownCodeFence)
		builder.WriteString(segment.meta)
		builder.WriteByte('\n')
		builder.WriteString(escapeCodeValue(segment.text))
		builder.WriteByte('\n')
		builder.WriteString(markdownCodeFence)
	case segmentSpoiler:
		builder.WriteString("||")
		builder.WriteString(escapeValue(segment.text))
		builder.WriteString("||")
	case segmentUnderline:
		builder.WriteString("__")
		builder.WriteString(escapeValue(segment.text))
		builder.WriteString("__")
	case segmentStrike:
		builder.WriteByte('~')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteByte('~')
	case segmentHyperLink:
		builder.WriteByte('[')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteString("](")
		builder.WriteString(escapeValue(segment.meta))
		builder.WriteByte(')')
	case segmentMention:
		builder.WriteByte('[')
		builder.WriteString(escapeValue(segment.text))
		builder.WriteString("](")
		builder.WriteString(telegramUserIDPrefix)
		builder.WriteString(segment.meta)
		builder.WriteByte(')')
	}
}

func checkSecrets(value string) string {
	secretMu.RLock()
	defer secretMu.RUnlock()

	for _, current := range secrets {
		value = strings.ReplaceAll(value, current.value, current.name)
	}

	return value
}
