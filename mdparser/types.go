// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

type WMarkDown interface {
	// Append method appends the w-markdown value to a new copy of the current
	// w-markdown and returns the copy.
	Append(md WMarkDown) WMarkDown
	AppendThis(md WMarkDown) WMarkDown
	ReplaceMd(md1, md2 WMarkDown) WMarkDown
	ReplaceMdN(md1, md2 WMarkDown, n int) WMarkDown
	ReplaceMdThis(md1, md2 WMarkDown) WMarkDown
	ReplaceMdThisN(md1, md2 WMarkDown, n int) WMarkDown
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
	ReplaceToNew(text1, text2 string) WMarkDown
	ReplaceToNewN(text1, text2 string, n int) WMarkDown

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
	Replace(text1, text2 string) WMarkDown
}

type wotoMarkDown struct {
	_value string
}

// secretContainer contains information about a secret value which needs
// to be replaced by its name when it's being appended to a markdown value.
type secretContainer struct {
	name  string
	value string
}
