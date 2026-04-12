// mdparser library Project
// Copyright (C) 2021-2022 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"strings"
)

func (m *wotoMarkDown) Append(v WMarkDown) WMarkDown {
	if v == nil {
		return nil
	}

	return m.appendRaw(v.ToString())
}

func (m *wotoMarkDown) ReplaceMd(md1, md2 WMarkDown) WMarkDown {
	return m.replaceRaw(md1.ToString(), md2.ToString(), -1)
}

func (m *wotoMarkDown) ReplaceMdN(md1, md2 WMarkDown, n int) WMarkDown {
	return m.replaceRaw(md1.ToString(), md2.ToString(), n)
}

func (m *wotoMarkDown) ReplaceToNew(text1, text2 string) WMarkDown {
	return m.replaceRaw(toNormal(text1), toNormal(text2), -1)
}

func (m *wotoMarkDown) ReplaceToNewN(text1, text2 string, n int) WMarkDown {
	return m.replaceRaw(toNormal(text1), toNormal(text2), n)
}

func (m *wotoMarkDown) Clone() WMarkDown {
	return &wotoMarkDown{_value: m._value}
}

func (m *wotoMarkDown) ToString() string {
	return m._value
}

func (m *wotoMarkDown) Normal(text string) WMarkDown {
	return m.appendText(text, toNormal)
}

func (m *wotoMarkDown) Bold(text string) WMarkDown {
	return m.appendText(text, toBold)
}

func (m *wotoMarkDown) Italic(text string) WMarkDown {
	return m.appendText(text, toItalic)
}

func (m *wotoMarkDown) Mono(text string) WMarkDown {
	return m.appendText(text, toMono)
}

func (m *wotoMarkDown) CodeBlock(text string) WMarkDown {
	return m.appendText(text, toCodeBlock)
}

func (m *wotoMarkDown) CodeBlockLang(lang, text string) WMarkDown {
	return m.appendPair(lang, text, toCodeBlockLang)
}

func (m *wotoMarkDown) Strike(text string) WMarkDown {
	return m.appendText(text, toStrike)
}

func (m *wotoMarkDown) Underline(text string) WMarkDown {
	return m.appendText(text, toUnderline)
}

func (m *wotoMarkDown) HyperLink(text, url string) WMarkDown {
	return m.appendPair(text, url, toHyperLink)
}

func (m *wotoMarkDown) Link(text, url string) WMarkDown {
	return m.appendPair(text, url, toHyperLink)
}

// Mention, mentions a user.
func (m *wotoMarkDown) Mention(text string, id int64) WMarkDown {
	return m.appendMention(text, id)
}

// UserMention, mentions a user.
func (m *wotoMarkDown) UserMention(text string, id int64) WMarkDown {
	return m.appendMention(text, id)
}

func (m *wotoMarkDown) Spoiler(text string) WMarkDown {
	return m.appendText(text, toSpoiler)
}

// El method appends a new line (End-line) to the markdown value.
func (m *wotoMarkDown) El() WMarkDown {
	return m.Normal("\n")
}

// Space method appends a space to the markdown value.
func (m *wotoMarkDown) Space() WMarkDown {
	return m.Normal(" ")
}

// Tab method appends a tab ("\t") to the markdown value.
func (m *wotoMarkDown) Tab() WMarkDown {
	return m.Normal("\t")
}

func (m *wotoMarkDown) Replace(text1, text2 string) WMarkDown {
	return m.replaceRaw(toNormal(text1), toNormal(text2), -1)
}

func (m *wotoMarkDown) ReplaceN(text1, text2 string, n int) WMarkDown {
	return m.replaceRaw(toNormal(text1), toNormal(text2), n)
}

func (m *wotoMarkDown) appendRaw(value string) WMarkDown {
	m._value += value
	return m
}

func (m *wotoMarkDown) appendText(text string, formatter func(string) string) WMarkDown {
	if text == "" {
		return m
	}

	return m.appendRaw(formatter(text))
}

func (m *wotoMarkDown) appendPair(text, extra string, formatter func(string, string) string) WMarkDown {
	if text == "" || extra == "" {
		return m
	}

	return m.appendRaw(formatter(text, extra))
}

func (m *wotoMarkDown) appendMention(text string, id int64) WMarkDown {
	if text == "" || id == baseIndex {
		return m
	}

	return m.appendRaw(toUserMention(text, id))
}

func (m *wotoMarkDown) replaceRaw(oldValue, newValue string, n int) WMarkDown {
	m._value = strings.Replace(m._value, oldValue, newValue, n)
	return m
}
