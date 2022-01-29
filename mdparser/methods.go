// mdparser library Project
// Copyright (C) 2021 ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package mdparser

import (
	"reflect"

	"github.com/ALiwoto/StrongStringGo/strongStringGo"
)

func (m *wotoMarkDown) setValue(text string) {
	m._value = text
}

func (m *wotoMarkDown) Append(v WMarkDown) WMarkDown {
	if reflect.TypeOf(v) == reflect.TypeOf(m) {
		md := v.(*wotoMarkDown)
		str := m._value + md._value
		wmd := &wotoMarkDown{
			_value: str,
		}

		return wmd
	}

	return nil
}

func (m *wotoMarkDown) AppendThis(v WMarkDown) WMarkDown {
	m.setValue(m.getValue() + v.getValue())

	return m
}

func (m *wotoMarkDown) ToString() string {
	return m._value
}

func (m *wotoMarkDown) AppendNormal(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toNormal(text))
}

func (m *wotoMarkDown) AppendNormalThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toNormal(text))

	return m
}

func (m *wotoMarkDown) AppendBold(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toBold(text))
}

func (m *wotoMarkDown) AppendBoldThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toBold(text))

	return m
}

func (m *wotoMarkDown) AppendItalic(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toItalic(text))
}

func (m *wotoMarkDown) AppendItalicThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toItalic(text))

	return m
}

func (m *wotoMarkDown) AppendMono(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toMono(text))
}

func (m *wotoMarkDown) AppendMonoThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toMono(text))

	return m
}

func (m *wotoMarkDown) AppendUnderline(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toUnderline(text))
}

func (m *wotoMarkDown) AppendUnderlineThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toUnderline(text))

	return m
}

func (m *wotoMarkDown) AppendStrike(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toStrike(text))
}

func (m *wotoMarkDown) AppendStrikeThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toStrike(text))

	return m
}

func (m *wotoMarkDown) AppendHyperLink(text, url string) WMarkDown {
	if text == "" || strongStringGo.IsEmpty(&url) {
		return m
	}

	return toWotoMD(m.getValue() + toHyperLink(text, url))
}

func (m *wotoMarkDown) AppendHyperLinkThis(text, url string) WMarkDown {
	if text == "" || strongStringGo.IsEmpty(&url) {
		return m
	}

	m.setValue(m.getValue() + toHyperLink(text, url))

	return m
}

func (m *wotoMarkDown) AppendMention(text string, id int64) WMarkDown {
	if text == "" || id == strongStringGo.BaseIndex {
		return m
	}

	return toWotoMD(m.getValue() + toUserMention(text, id))
}

func (m *wotoMarkDown) AppendMentionThis(text string, id int64) WMarkDown {
	if text == "" || id == strongStringGo.BaseIndex {
		return m
	}

	m.setValue(m.getValue() + toUserMention(text, id))

	return m
}

func (m *wotoMarkDown) AppendSpoiler(text string) WMarkDown {
	if text == "" {
		return m
	}

	return toWotoMD(m.getValue() + toSpoiler(text))
}

func (m *wotoMarkDown) AppendSpoilerThis(text string) WMarkDown {
	if text == "" {
		return m
	}

	m.setValue(m.getValue() + toSpoiler(text))

	return m
}

func (m *wotoMarkDown) Normal(text string) WMarkDown {
	return m.AppendNormalThis(text)
}

func (m *wotoMarkDown) Bold(text string) WMarkDown {
	return m.AppendBoldThis(text)
}

func (m *wotoMarkDown) Italic(text string) WMarkDown {
	return m.AppendItalicThis(text)
}

func (m *wotoMarkDown) Mono(text string) WMarkDown {
	return m.AppendMonoThis(text)
}

func (m *wotoMarkDown) Strike(text string) WMarkDown {
	return m.AppendStrikeThis(text)
}

func (m *wotoMarkDown) Underline(text string) WMarkDown {
	return m.AppendUnderlineThis(text)
}

func (m *wotoMarkDown) HyperLink(text, url string) WMarkDown {
	return m.AppendHyperLinkThis(text, url)
}

func (m *wotoMarkDown) Link(text, url string) WMarkDown {
	return m.AppendHyperLinkThis(text, url)
}

// Mention, mentions a user.
func (m *wotoMarkDown) Mention(text string, id int64) WMarkDown {
	return m.AppendMentionThis(text, id)
} //UserMention

// UserMention, mentions a user.
func (m *wotoMarkDown) UserMention(text string, id int64) WMarkDown {
	return m.AppendMentionThis(text, id)
}

func (m *wotoMarkDown) Spoiler(text string) WMarkDown {
	return nil
}

// El method appends a new line (Endline) to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) El() WMarkDown {
	return m.AppendNormal("\n")
}

// ElThis method appends a new line (Endline) to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) ElThis() WMarkDown {
	return m.AppendNormalThis("\n")
}

// Space method appends a space to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) Space() WMarkDown {
	return m.AppendNormal(" ")
}

// SpaceThis method appends a new line to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) SpaceThis() WMarkDown {
	return m.AppendNormalThis(" ")
}

// Tab method appends a tab ("\t") to the markdown value and returns
// a new instance of WMarkDown.
func (m *wotoMarkDown) Tab() WMarkDown {
	return m.AppendNormal("\t")
}

// TabThis method appends a tab ("\t") to the current markdown value and
// returns the current instance of WMarkDown.
func (m *wotoMarkDown) TabThis() WMarkDown {
	return m.AppendNormalThis("\t")
}

func (m *wotoMarkDown) getValue() string {
	return m._value
}
