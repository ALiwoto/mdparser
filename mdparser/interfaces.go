// mdparser library Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	AppendThis(md WMarkDown) WMarkDown
	ToString() string
	AppendNormal(v string) WMarkDown
	AppendNormalThis(v string) WMarkDown
	AppendBold(v string) WMarkDown
	AppendBoldThis(v string) WMarkDown
	AppendItalic(v string) WMarkDown
	AppendItalicThis(v string) WMarkDown
	AppendMono(v string) WMarkDown
	AppendMonoThis(v string) WMarkDown
	AppendHyperLink(text, url string) WMarkDown
	AppendHyperLinkThis(text, url string) WMarkDown
	AppendMention(text string, id int64) WMarkDown
	AppendMentionThis(text string, id int64) WMarkDown
	El() WMarkDown
	ElThis() WMarkDown
	Space() WMarkDown
	SpaceThis() WMarkDown
	Tab() WMarkDown
	TabThis() WMarkDown
	getValue() string
	setValue(v string)
}
