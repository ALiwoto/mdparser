// mdparser library Project
// Copyright (C) 2021-2026 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

const (
	baseIndex            = int64(0)
	markdownCodeFence    = "```"
	markdownEscapeChar   = '\\'
	telegramUserIDPrefix = "tg://user?id="
	specialChars         = "\\'_*{}[]()#+-.!`=><~|"
)

const (
	segmentRaw markdownSegmentKind = iota
	segmentNormal
	segmentBold
	segmentItalic
	segmentStyled
	segmentMono
	segmentCodeBlock
	segmentCodeBlockLang
	segmentSpoiler
	segmentUnderline
	segmentStrike
	segmentHyperLink
	segmentMention
)

const (
	// StyleBold applies bold formatting.
	StyleBold TextStyle = 1 << iota
	// StyleItalic applies italic formatting.
	StyleItalic
	// StyleUnderline applies underline formatting.
	StyleUnderline
	// StyleStrike applies strike-through formatting.
	StyleStrike
	// StyleSpoiler applies spoiler formatting.
	StyleSpoiler
)

const supportedTextStyles = StyleBold | StyleItalic | StyleUnderline | StyleStrike | StyleSpoiler
