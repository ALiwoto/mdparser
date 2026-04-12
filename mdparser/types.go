// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

type WMarkDown interface {
	// Append method appends the given markdown value to the current markdown.
	Append(md WMarkDown) WMarkDown
	Clone() WMarkDown
	ReplaceMd(md1, md2 WMarkDown) WMarkDown
	ReplaceMdN(md1, md2 WMarkDown, n int) WMarkDown
	ToString() string

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
	Space() WMarkDown
	Tab() WMarkDown
	Replace(text1, text2 string) WMarkDown
	ReplaceN(text1, text2 string, n int) WMarkDown
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
