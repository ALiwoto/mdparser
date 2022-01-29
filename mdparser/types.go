// mdparser library Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	AppendThis(md WMarkDown) WMarkDown
	ToString() string

	AppendNormal(text string) WMarkDown
	AppendNormalThis(text string) WMarkDown
	AppendBold(text string) WMarkDown
	AppendBoldThis(text string) WMarkDown
	AppendItalic(text string) WMarkDown
	AppendItalicThis(text string) WMarkDown
	AppendStrike(text string) WMarkDown
	AppendStrikeThis(text string) WMarkDown
	AppendMono(text string) WMarkDown
	AppendMonoThis(text string) WMarkDown
	AppendUnderline(text string) WMarkDown
	AppendUnderlineThis(text string) WMarkDown
	AppendHyperLink(text, url string) WMarkDown
	AppendHyperLinkThis(text, url string) WMarkDown
	AppendMention(text string, id int64) WMarkDown
	AppendMentionThis(text string, id int64) WMarkDown

	Normal(text string) WMarkDown
	Bold(text string) WMarkDown
	Italic(text string) WMarkDown
	Mono(text string) WMarkDown
	HyperLink(text, url string) WMarkDown
	Link(text, url string) WMarkDown
	Mention(text string, id int64) WMarkDown
	UserMention(text string, id int64) WMarkDown
	Spoiler(text string) WMarkDown
	Strike(text string) WMarkDown
	Underline(text string) WMarkDown

	El() WMarkDown
	ElThis() WMarkDown
	Space() WMarkDown
	SpaceThis() WMarkDown
	Tab() WMarkDown
	TabThis() WMarkDown
	getValue() string
	setValue(text string)
}

type wotoMarkDown struct {
	_value string
}
